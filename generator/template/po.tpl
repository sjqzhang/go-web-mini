{{define "po"}}package model

import (

    "time"

)


// {{.Table.TableName}} {{.Table.TableComment}}
type {{.Table.TableName}} struct {
    Model
    {{range .Fields}}{{if  checkField .ColumnName}}
    {{.CamelField}} {{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}{{end}}
     {{end}}
}

// {{.Table.TableName}} {{.Table.TableComment}}
type {{.Table.TableName}}Query struct {
 {{range .Fields}}
    {{.CamelField}} *{{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}}{{end}}

          PageNum  int   `json:"pageNum" form:"pageNum"`
          PageSize int   `json:"pageSize" form:"pageSize"`
}

{{end}}
