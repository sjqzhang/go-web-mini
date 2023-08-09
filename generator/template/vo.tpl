{{define "vo"}}package vo

import (
    "time"
)




type {{.Table.TableName}}Response struct {
{{range .Fields}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }} form:"{{.ColumnName}}"` {{.ColumnComment}}
{{end}}
}


// 查询{{.Table.TableName}} {{.Table.TableComment}}
type {{.Table.TableName}} struct {
 {{range .Fields}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }} form:"{{.ColumnName}}"` {{.ColumnComment}}
     {{end}}
}

// 查询{{.Table.TableName}} {{.Table.TableComment}}
type List{{.Table.TableName}}Response struct {
	Total    int64                  `json:"total"`
	List     []{{.Table.TableName}}          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询{{.Table.TableName}} {{.Table.TableComment}}
type List{{.Table.TableName}}Request struct {
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}  form:"{{.ColumnName}}"` {{.ColumnComment}}
    {{end}}
     {{end}}
     PageNum  *uint   `json:"pageNum" form:"pageNum"` //第几页
     PageSize *uint   `json:"pageSize" form:"pageSize"` //每页多少条
}

type Get{{.Table.TableName}}Response struct {
    {{.Table.TableName}}Response
}


// 创建{{.Table.TableName}} {{.Table.TableComment}}
type Create{{.Table.TableName}}Request struct {
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }} form:"{{.ColumnName}}"` {{.ColumnComment}}
    {{end}}
     {{end}}
}

type Create{{.Table.TableName}}Response struct {
    {{.Table.TableName}}Response
}



// 更新{{.Table.TableName}} {{.Table.TableComment}}
type Update{{.Table.TableName}}Request struct {
    ID      *int `json:""`
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }} form:"{{.ColumnName}}"` {{.ColumnComment}}
    {{end}}
     {{end}}
}

type Update{{.Table.TableName}}Response struct {
    {{.Table.TableName}}Response
}

// 删除{{.Table.TableName}} {{.Table.TableComment}}
type Delete{{.Table.TableName}}Request struct {
    Ids      []int64 `json:"ids" uri:"ids" form:"ids"` //待编号
}




// 删除{{.Table.TableName}} {{.Table.TableComment}}
type Get{{.Table.TableName}}Request struct {
    ID      int64 `json:"id" uri:"id" form:"id"` //待编号
}


type Delete{{.Table.TableName}}Response struct {
    Response
    Data int `json:"data"`
}


{{end}}
