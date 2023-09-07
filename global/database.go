package global

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjqzhang/gdi"
	"go-web-mini/config"

	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

// 全局mysql数据库变量
var DB *gorm.DB

var dbKey = "db"
var GinContextKey = "ginContext"

//func BindContextForGorm(ctx context.Context) context.Context {
//	db:=GetDB(ctx)
//	return context.WithValue(ctx, dbKey, db)
//}
func GetDB(ctx context.Context) *gorm.DB {
	db := ctx.Value("db")
	if db == nil {
		return DB.WithContext(ctx)
	}
	db.(*gorm.DB).WithContext(ctx)
	return db.(*gorm.DB)
}

type CustomLogger struct {
}

func (cl *CustomLogger) LogMode(logger.LogLevel) logger.Interface {
	return cl
}

func (cl *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	requestID := getRequestIDFromContext(ctx)
	dataLogger.Infof("[INFO] [Request ID: %v] %svn", requestID, fmt.Sprintf(msg, data...))
}

func (cl *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	requestID := getRequestIDFromContext(ctx)
	dataLogger.Infof("[WARN] [Request ID: %v] %v\n", requestID, fmt.Sprintf(msg, data...))
}

func (cl *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	requestID := getRequestIDFromContext(ctx)
	dataLogger.Infof("[ERROR] [Request ID: %v] %v\n", requestID, fmt.Sprintf(msg, data...))
}

func (cl *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	requestID := getRequestIDFromContext(ctx)
	elapsed := time.Since(begin)
	sql, rows := fc()
	if _, ok := ctx.(*gin.Context); ok {
		sqls := ctx.(*gin.Context).Value("SQL")
		if sqls != nil {
			sqls = append(sqls.([]string), sql)
			ctx.(*gin.Context).Set("SQL", sqls)
		} else {
			ctx.(*gin.Context).Set("SQL", []string{sql})
		}
	}
	//if strings.Contains(sql, "INSERT") || strings.Contains(sql, "UPDATE") || strings.Contains(sql, "DELETE") {
	//dataLogger.Infof("[TRACE] [Request ID: %v] [%.3fms] %v[SQL]: %v [%v]\n", requestID, float64(elapsed.Microseconds())/1000, sql, rows, err)
	//}
	dataLogger.Infof("[TRACE] [Request ID: %v] [%.3fms] %v[SQL]: %v [%v]\n", requestID, float64(elapsed.Microseconds())/1000, sql, rows, err)

}

func getRequestIDFromContext(ctx context.Context) string {
	if reqID := ctx.Value("X-Request-ID"); reqID != nil {
		if id, ok := reqID.(string); ok {
			return id
		}
	}
	return ""
}

// 初始化mysql数据库
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	// 隐藏密码
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	//Log.Info("数据库连接DSN: ", showDsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		//// 指定表前缀
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: config.Conf.Mysql.TablePrefix + "_",
		//},
	})

	db = db.Debug()
	if err != nil {
		Log.Panicf("初始化mysql数据库异常: %v", err)
		panic(fmt.Errorf("初始化mysql数据库异常: %v", err))
	}

	// 开启mysql日志
	if config.Conf.Mysql.LogMode {
		db.Debug()
	}
	// 全局DB赋值
	DB = db
	// 自动迁移表结构
	dbAutoMigrate()
	db.Logger = &CustomLogger{}
	Log.Infof("初始化mysql数据库完成! dsn: %s", showDsn)
}

// 自动迁移表结构
func dbAutoMigrate() {
	//DB.AutoMigrate(
	//	&model.User{},
	//	&model.Role{},
	//	&model.Menu{},
	//	&model.Api{},
	//	&model.OperationLog{},
	//	&model.News{},
	//	&model.TableMetadata{},
	//)
	models, err := gdi.GetAllTypesByPackName("model*")

	if err != nil {
		Log.Panicf("自动注册model失败: %v", err)
	}
	for _, mt := range models {
		// 从model类型，反射创建对象
		if mt.Kind() != reflect.Ptr {
			continue
		}
		if mt.Elem().Kind() != reflect.Struct {
			continue
		}
		m := reflect.New(mt.Elem()).Interface()
		//如果有TableName方法，就调用
		if _, ok := m.(interface{ TableName() string }); ok {
			t := fmt.Sprintf("%T", m)
			fmt.Println("自动注册model:", t)
			if strings.HasPrefix(t, "*model.") {
				err:=DB.AutoMigrate(m)
				if err!=nil{
					Log.Panicf("自动注册model失败: %v", err)
				}
			}
		}
	}
}
