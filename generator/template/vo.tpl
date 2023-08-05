{{define "vo"}}package vo

import (
    "time"
)


// 查询{{.Table.TableName}} {{.Table.TableComment}}
type {{.Table.TableName}} struct {
 {{range .Fields}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}
     {{end}}
}

// 查询{{.Table.TableName}} {{.Table.TableComment}}
type Pager{{.Table.TableName}} struct {
	Total    int64                  `json:"total"`
	List     []{{.Table.TableName}}          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询{{.Table.TableName}} {{.Table.TableComment}}
type List{{.Table.TableName}}Request struct {
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}
    {{end}}
     {{end}}
     PageNum  *uint   `json:"pageNum" form:"pageNum"`
     PageSize *uint   `json:"pageSize" form:"pageSize"`
}


// 创建{{.Table.TableName}} {{.Table.TableComment}}
type Create{{.Table.TableName}}Request struct {
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}
    {{end}}
     {{end}}
}


// 更新{{.Table.TableName}} {{.Table.TableComment}}
type Update{{.Table.TableName}}Request struct {
    ID      *int `json:""`
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}
    {{end}}
     {{end}}
}

// 删除{{.Table.TableName}} {{.Table.TableComment}}
type Delete{{.Table.TableName}}Request struct {
    ID      int `json:"id" uri:"id" form:"id"`
}

{{end}}
