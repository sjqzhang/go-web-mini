package routes

import (
	"fmt"
	"github.com/sjqzhang/gdi"
	"go-web-mini/apps/system/model"
	"go-web-mini/util"
	"gorm.io/gorm"
	"strings"
	"sync"

	//"fmt"
	"github.com/gin-gonic/gin"
	"go-web-mini/config"
	_ "go-web-mini/docs"
	"go-web-mini/global"
	"go-web-mini/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化
func InitRoutes() *gin.Engine {

	routerMap, _ := gdi.GetRouterInfoByPatten(".*/controller")

	restInfo, _ := gdi.GetRestInfoByPatten(".*/controller")

	var unit9 uint = 9
	for _, info := range restInfo {

		// 通过restInfo生成系统后台菜单

		if info.Description == "" {
			continue
		}
		name := strings.ReplaceAll(info.Controller, "Controller", "")
		var menu model.Menu
		component := fmt.Sprintf("/business/%v/index", util.ToUnderlineCase(name))
		err := global.DB.First(&menu, "component = ?", component).Error
		if err != nil && err == gorm.ErrRecordNotFound {
			menu := model.Menu{
				//Model:     gorm.Model{ID: 9},
				Name:  "Business",
				Title: info.Description,
				//Icon:     "table",
				Path:      strings.ToLower(name),
				Component: component,
				Sort:      23,
				ParentId:  &unit9,
				//Roles:     roles[:2],
				Creator: "系统",
			}
			global.DB.Create(&menu)
		}
	}

	ctrls, err := gdi.AutoRegisterByPackagePatten(`.*controller.*`)

	gdi.Init() //初始化对象池

	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件

	// 创建不带中间件的路由:
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DeepLinking(true), ginSwagger.PersistAuthorization(true),
		ginSwagger.DefaultModelsExpandDepth(5),
	))

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个
	//fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	//capacity := config.Conf.RateLimit.Capacity
	//r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	//r.Use(middleware.CacheMiddleware(time.Hour*5))

	// 启用操作日志中间件
	r.Use(middleware.OperationLogMiddleware())

	// 初始化JWT认证中间件
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		global.Log.Panicf("初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}

	// 路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 注册路由
	//InitBaseRoutes(apiGroup, authMiddleware) // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	//InitUserRoutes(apiGroup, authMiddleware)         // 注册用户路由, jwt认证中间件,casbin鉴权中间件
	//InitRoleRoutes(apiGroup, authMiddleware)         // 注册角色路由, jwt认证中间件,casbin鉴权中间件
	//InitMenuRoutes(apiGroup, authMiddleware)         // 注册菜单路由, jwt认证中间件,casbin鉴权中间件
	//InitApiRoutes(apiGroup, authMiddleware)          // 注册接口路由, jwt认证中间件,casbin鉴权中间件
	//InitOperationLogRoutes(apiGroup, authMiddleware) // 注册操作日志路由, jwt认证中间件,casbin鉴权中间件

	apiGroup.POST("/base/login", authMiddleware.LoginHandler)
	apiGroup.POST("/base/logout", authMiddleware.LogoutHandler)
	apiGroup.POST("/base/refreshToken", authMiddleware.RefreshHandler)

	global.Log.Info("初始化路由完成！")

	middlewaresMapFunc := make(map[string]func(sync.Map) gin.HandlerFunc)

	middlewaresMaping := make(map[string]gin.HandlerFunc)

	//middlewaresMaping["ratelimit"]=middleware.RateLimitMiddleware(time.Millisecond*time.Duration(config.Conf.RateLimit.FillInterval),config.Conf.RateLimit.Capacity)
	middlewaresMaping["cors"] = middleware.CORSMiddleware()
	middlewaresMaping["operationlog"] = middleware.OperationLogMiddleware()
	auth, err := middleware.InitAuth()
	if err != nil {
		panic(err)
	}
	middlewaresMaping["auth"] = auth.MiddlewareFunc()
	middlewaresMaping["casbin"] = middleware.CasbinMiddleware()
	middlewaresMaping["transition"] = middleware.TransitionMiddleware()

	middlewaresMapFunc["cache"] = middleware.CacheMiddleware

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	pkgPrefix:= gdi.GetAppModuleName()
	for _, o := range ctrls { //自动绑定路由
		ctrlName := o.Elem().Type().Name()
		packName := o.Elem().Type().PkgPath()
		packName = strings.TrimPrefix(packName,pkgPrefix+"/")

		for i := 0; i < o.NumMethod(); i++ {
			if o.NumMethod() == 0 {
				continue
			}
			if !o.Type().Method(i).IsExported() {
				continue
			}
			var v gdi.RouterInfo
			var ok bool
			methodName := o.Type().Method(i).Name
			key := fmt.Sprintf("%v.%v.%v", packName, ctrlName, methodName)
			if v, ok = routerMap[key]; !ok {
				v.Method = "POST"
				v.Uri = fmt.Sprintf("/%v/%v", ctrlName, methodName)
				v.Handler = methodName
			}

			x := o.Method(i).Interface()

			if o.Method(i).Type().NumIn() > 2 {
				continue
			}

			if o.Method(i).Type().NumIn() == 2 {
				if o.Method(i).Type().In(0).String() != "*gin.Context" && o.Method(i).Type().In(0).String() != "context.Context" {
					continue
				}
			}

			var mds []gin.HandlerFunc
			for _, middleWare := range v.Middlewares {
				count := 0
				middleWare.Params.Range(func(key, value interface{}) bool {
					count++
					return true
				})
				if count > 0 {
					if f, ok := middlewaresMapFunc[middleWare.Name]; ok {
						mds = append(mds, f(middleWare.Params))
					}
				} else {
					if f, ok := middlewaresMaping[middleWare.Name]; ok {
						mds = append(mds, f)
					}
				}
			}

			api := model.Api{
				Method:   v.Method,
				Path:     strings.TrimPrefix(v.Uri, "/"+config.Conf.System.UrlPathPrefix),
				Category: strings.ToLower(strings.TrimSuffix(v.Controller, "Controller")),
				Desc:     v.Description,
				Creator:  "系统",
			}

			global.DB.Model(model.Api{}).Create(&api)

			switch x.(type) {
			case func(*gin.Context):
				{
					mds = append(mds, x.(func(*gin.Context)))
					r.Handle(v.Method, v.Uri, mds...)

				}
			default:
				mds = append(mds, middleware.BinderMiddleware(o.Method(i)))
				r.Handle(v.Method, v.Uri, mds...)
			}
		}

	}

	return r
}
