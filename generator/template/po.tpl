{{define "po"}}package model

import (

    "time"

)

var _=time.Now()


// {{.Table.TableName}} {{.Table.TableComment}}
type {{.Table.TableName}} struct {
    Model
    {{range .Fields}}{{if  checkField .ColumnName}}{{.CamelField}} {{.RealType}} `gorm:"{{if notEmpty .ColumnName}}{{.IndexName}}{{end}}{{.ColumnName}};type:{{.ColumnType}};comment:{{.ColumnCommentForView}}" validate:"{{.Validate}}" json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}
    {{end}}{{end}}
}

// {{.Table.TableName}} {{.Table.TableComment}}
type {{.Table.TableName}}Query struct {
 {{range .Fields}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}{{end}}
    PageNum  int   `json:"-" form:"pageNum"`
    PageSize int   `json:"-" form:"pageSize"`
}


func (t {{.Table.TableName}})TableName() string {
    return "{{.Table.TableNameOrigin}}"
}

{{end}}
