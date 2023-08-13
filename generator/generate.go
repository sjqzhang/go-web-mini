package generator

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"strings"
	"sync"
	"text/template"
)

type FieldResult struct {
	ColumnName           string
	ColumnType           string
	CamelField           string
	DataType             string
	RealType             string
	ColumnKey            string
	KeyStr               string
	Extra                string
	ColumnDefault        string
	ColumnComment        string
	ColumnCommentForView string
	Value                string
	VueTag               string
	VueType              string
	VueFunction          string
}

type TableResult struct {
	TableName       string
	Uri             string
	TableComment    string
	TableNameOrigin string
	TableNameTrim   string
	IsAuth          bool
	//ButtonTop       []string
	//ButtonRight     []string
}

type CommonObject struct {
	Fields     []FieldResult
	Table      TableResult
	ModuleName string
}

type Config struct {
	Tables      []string
	WebRoot     string
	ServerRoot  string
	DSN         string
	ModuleName  string
	TablePrefix string
	TableSuffix string
	IsAuth      bool
}

var cfg Config

func InitConfig(conf Config) {
	cfg = conf
}

func generate(url string, tableNames []string, moduleName string) {

	db, database := connect(url)

	// 获取数据库连接，生成
	generateCode(db, database, tableNames, moduleName)
}

func DoGenerate() {

	if cfg.DSN == "" {
		panic(fmt.Sprintf("please InitConfig First!"))
	}

	generate(cfg.DSN, cfg.Tables, cfg.ModuleName)
	cmds := []string{
		"cd", cfg.ServerRoot, "&&", "go", "list", "./...", "|", "xargs", "go", "fmt",
	}
	exec.Command("sh", "-c", strings.Join(cmds, " ")).Run()
}

// generate 生成代码
func generateCode(con *gorm.DB, database string, tableNames []string, moduleName string) {

	// 创建所需的文件夹
	createDirs("..")

	// 循环生成
	for _, tableName := range tableNames {
		wg.Add(1)
		tableName := tableName
		go doGenerate(con, database, tableName, moduleName)
	}
	wg.Wait()
}

var wg sync.WaitGroup

// 生成单个表的文件
func doGenerate(con *gorm.DB, database string, tableName string, moduleName string) {

	defer wg.Done()

	// 查询表信息
	tableQuery, _ := con.Raw("select "+
		"TABLE_NAME as TableName,"+
		"TABLE_COMMENT as TableComment ,"+
		"'' as TableNameOrigin "+
		"from "+
		"information_schema.TABLES "+
		"where "+
		"table_schema = ? and table_name = ?;", database, tableName).Rows()

	defer func(tableQuery *sql.Rows) {
		err := tableQuery.Close()
		if err != nil {
			fmt.Println(err)
			panic("failed to close")
		}
	}(tableQuery)

	// 查询属性信息
	fieldQuery, _ := con.Raw("select "+
		"COLUMN_NAME as ColumnName ,"+
		"COLUMN_TYPE AS ColumnType,"+
		"COLUMN_DEFAULT as ColumnDefault ,"+
		"DATA_TYPE as DataType,"+
		"COLUMN_KEY as ColumnKey,"+
		"EXTRA as Extra,"+
		"COLUMN_COMMENT as ColumnComment ,"+
		"'' as ColumnCommentForView "+
		"from "+
		"information_schema.columns "+
		"where "+
		"table_schema = ? and table_name = ?;", database, tableName).Rows()

	defer func(fieldQuery *sql.Rows) {
		err := fieldQuery.Close()
		if err != nil {
			fmt.Println(err)
			panic("failed to close")
		}
	}(fieldQuery)

	// 表信息转换到切片中
	tables := convertTable(con, tableQuery)
	// 表中的属性信息转换到切片中
	fields := convertField(con, fieldQuery)

	// 校验表是否存在
	if len(tables) == 0 {
		panic("cannot find the table: " + tableName)
	}

	// 处理属性
	handleFields(fields)

	// 设置表信息
	tableInfo := tables[0]
	tableInfo.TableNameOrigin = tableName // 原始表名，带前缀和后缀
	tableName = strings.Replace(tableName, cfg.TablePrefix, "", 1)
	tableName = strings.Replace(tableName, cfg.TableSuffix, "", 1)
	tableInfo.TableNameTrim = tableName // 去掉前缀和后缀的表名
	tableInfo.TableName = TransToCamel(tableName, false)
	tableInfo.Uri = TransToCamel(tableName, true)

	// 定义模板中需要访问的对象并赋值
	var object CommonObject
	object.Table = tableInfo
	object.Fields = fields
	object.ModuleName = moduleName

	// 创建文件
	createFiles(object, tableName)
}

// 检查指定的filed 是否在指定的fields中
func checkField(field string) bool {
	fields := []string{"id", "created_at", "updated_at", "deleted_at"}
	for _, v := range fields {
		if v == field {
			return false
		}
	}
	return true
}

// 创建文件
func createFiles(obj CommonObject, tableName string) {

	// 创建po
	createGoFile(obj, tableName, fmt.Sprintf("%v.go", tableName), fmt.Sprintf("%v/model", cfg.ServerRoot), "./template/po.tpl", "po")

	// 创建vo
	createGoFile(obj, tableName, fmt.Sprintf("%v_request.go", tableName), fmt.Sprintf("%v/vo", cfg.ServerRoot), "./template/vo.tpl", "vo")

	// 创建add dto
	//createGoFile(obj, tableName, "AddDTO.go", "./dto", "./template/addDto.tpl", "addDto")

	// 创建page dto
	createGoFile(obj, tableName, fmt.Sprintf("%v_repository.go", tableName), fmt.Sprintf("%v/repository", cfg.ServerRoot), "./template/dao.tpl", "repository")

	// 创建controller
	createGoFile(obj, tableName, fmt.Sprintf("%v_controller.go", tableName), fmt.Sprintf("%v/controller", cfg.ServerRoot), "./template/controller.tpl", "controller")

	// 创建router
	//createGoFile(obj, tableName, "Router.go", "./router", "./template/router.tpl", "router")

	// 创建service
	createGoFile(obj, tableName, fmt.Sprintf("%v_service.go", tableName), "../service", "./template/service.tpl", "service")

	// 创建service
	createGoFile(obj, tableName, "index.vue", fmt.Sprintf("%v/src/views/business/%v", cfg.WebRoot, tableName), "./template/index.vue", "index.vue")

	createGoFile(obj, tableName, fmt.Sprintf("%v.js", tableName), fmt.Sprintf("%v/src/api/business", cfg.WebRoot), "./template/api.js", "api")

}

// 创建所需的文件夹
func createDirs(modulePath string) {

	// 创建 po 目录
	createDir(modulePath, "/model")

	// 创建 vo 目录
	createDir(modulePath, "/vo")

	// 创建 dto 目录
	createDir(modulePath, "/dto")

	// 创建 controller 目录
	createDir(modulePath, "/controller")
	//
	//// 创建 service 目录
	//createDir(modulePath, "/service")

	// 创建 router 目录
	//createDir(modulePath, "/router")
}

// 创建文件夹
func createDir(modulePath string, dirName string) {

	// 路径
	path := modulePath + dirName

	// 查询文件是否存在
	_, exist := os.Stat(path)

	// 如果不存在，创建文件夹
	if exist != nil {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic("cannot create the dictionary " + dirName)

		}
	}
}

func isNumber(fieldType string) bool {
	return fieldType == "int" || fieldType == "tinyint" || fieldType == "bigint" || fieldType == "decimal" || fieldType == "float" || fieldType == "double"
}

func isString(fieldType string) bool {
	return fieldType == "varchar" || fieldType == "char" || fieldType == "text" || fieldType == "longtext" || fieldType == "mediumtext"
}

func isInteger(fieldType string) bool {
	return fieldType == "int" || fieldType == "tinyint" || fieldType == "bigint"
}

func isDate(fieldType string) bool {
	return fieldType == "datetime" || fieldType == "date" || fieldType == "timestamp"
}

func isTime(fieldType string) bool {
	return fieldType == "datetime" || fieldType == "timestamp"
}

// 创建go文件
func createGoFile(obj CommonObject, tableName string, filename string, path string, templatePath string, templateName string) {

	os.MkdirAll(path, 07777)
	// 创建文件
	file := createFile(filename, path)

	// 如果为空，直接返回，无需创建
	if file == nil {
		return
	}

	t := template.New("template")

	t = t.Funcs(template.FuncMap{"checkField": checkField, "isNumber": isNumber,
		"isString": isString, "isInteger": isInteger, "isDate": isDate, "isTime": isTime})

	// 校验是否存在po模板
	t = template.Must(t.ParseGlob(templatePath))

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(file, templateName, obj)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic("cannot create files with the template")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("close file error")
		}
	}(file)
}

// 创建文件，如果存在不创建
func createFile(fileName string, path string) *os.File {

	_, err := os.Stat(path + "/" + fileName)

	// 文件存在直接跳过
	if err == nil {
		fi, err := os.OpenFile(path+"/"+fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return fi
	} else {
		fmt.Println(err)
	}

	// 创建文件
	create, _ := os.Create(path + "/" + fileName)
	return create
}

// 处理属性信息
func handleFields(fields []FieldResult) {

	// 类型map
	typeMap := GetTypeMap()

	// 处理属性
	for i, f := range fields {
		fields[i].RealType = typeMap[f.DataType]

		// 字段注释
		if len(fields[i].ColumnComment) != 0 {
			fields[i].ColumnComment = "// " + fields[i].ColumnComment
		}

		fields[i].CamelField = TransToCamel(f.ColumnName, false)
		// 如果是主键，设置
		if f.ColumnKey == "PRI" {
			fields[i].KeyStr = " gorm:\"primary_key\""
			if f.ColumnName == "id" {
				fields[i].KeyStr = " gorm:\"primary_key;AUTO_INCREMENT\""
				fields[i].CamelField = "ID"
			}
		}
	}
}

// 表属性信息转换成对象
func convertField(con *gorm.DB, query *sql.Rows) []FieldResult {

	var fields []FieldResult

	for query.Next() {
		var str FieldResult
		err := con.ScanRows(query, &str)
		if err != nil {
			fmt.Println(err)
			panic("failed to scan rows")
		}
		if strings.TrimSpace(str.ColumnComment) == "" {
			str.ColumnComment = str.ColumnName
		}
		str.ColumnCommentForView = str.ColumnComment
		if strings.Index(strings.ToLower(str.DataType), "int") != -1 {
			str.Value = "0"
			str.VueType = "number"
			str.VueTag = "el-input"
			str.VueFunction = "Number"
		} else if strings.Index(strings.ToLower(str.DataType), "text") != -1 {
			str.Value = "''"
			str.VueType = "textarea"
			str.VueTag = "el-input"
			str.VueFunction = "String"
		} else if isDate(strings.ToLower(str.DataType)) {
			str.Value = "''"
			str.VueType = "date"
			str.VueTag = "el-date-picker"
			str.VueFunction = "String"
		} else if isTime(strings.ToLower(str.DataType)) {
			str.Value = "''"
			str.VueType = "datetime"
			str.VueTag = "el-date-picker"
			str.VueFunction = "String"
		} else {
			str.Value = "''"
			str.VueType = "text"
			str.VueTag = "el-input"
			str.VueFunction = "String"
		}
		fields = append(fields, str)
		fmt.Println(fmt.Sprintf("%v", str))
	}
	return fields
}

// 表信息转换成对象
func convertTable(con *gorm.DB, query *sql.Rows) []TableResult {

	var tables []TableResult

	for query.Next() {
		var str TableResult
		err := con.ScanRows(query, &str)
		str.TableName = strings.TrimPrefix(str.TableName, cfg.TablePrefix)
		str.TableName = strings.TrimSuffix(str.TableName, cfg.TableSuffix)
		if err != nil {
			fmt.Println(err)
			panic("failed to scan rows")
		}
		if strings.TrimSpace(str.TableComment) == "" {
			str.TableComment = str.TableName
		}
		tables = append(tables, str)
	}
	return tables
}

// 获取连接
func connect(url string) (*gorm.DB, string) {

	if len(url) == 0 {
		panic("connection strings cannot be empty")
	}

	// 获取连接
	con, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database,check your connection strings")
	}
	con.Debug()
	var dbName string
	con.Raw(" SELECT DATABASE()").First(&dbName)
	return con, dbName
}
