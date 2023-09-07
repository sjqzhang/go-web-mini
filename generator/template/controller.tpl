{{define "controller"}}package controller

import (
	"context"
    "{{.ModuleName}}/apps/{{.AppName}}/vo"
    "{{.ModuleName}}/apps/{{.AppName}}/service"

)

//@description {{.Table.TableComment}}管理
{{if .Table.IsAuth}}
//@middleware auth{{end}}
//@router /api/{{.AppName}}
type {{.Table.TableName}}Controller struct {
   {{.Table.Uri}}Service service.{{.Table.TableName}}Service
}

//@description 查询{{.Table.TableComment}}
//@tags {{.Table.TableName}}{{if .Table.IsAuth}}
// @Security JWT{{end}}
// @Accept       json
// @Produce      json
//@param req    query  vo.List{{.Table.TableName}}Request  false  "入参req"
//@success 200 {object} vo.List{{.Table.TableName}}Response
//@middleware
// @router /{{.Table.Uri}} [get]
func (c *{{.Table.TableName}}Controller) List(ctx context.Context, req *vo.List{{.Table.TableName}}Request) (*vo.List{{.Table.TableName}}Response, error) {
	return c.{{.Table.Uri}}Service.List(ctx, req)
}


//@description 查询单个{{.Table.TableComment}}
//@tags {{.Table.TableName}}{{if .Table.IsAuth}}
// @Security JWT{{end}}
// @Accept       json
// @Produce      json
//@param req    query  vo.Get{{.Table.TableName}}Request  false  "入参req"
//@success 200 {object} vo.Get{{.Table.TableName}}Response
//@middleware
// @router /{{.Table.Uri}}/:id [get]
func (c *{{.Table.TableName}}Controller) GetById(ctx context.Context, req *vo.Get{{.Table.TableName}}Request) (*vo.{{.Table.TableName}}Response, error) {
	return c.{{.Table.Uri}}Service.GetById(ctx, req)
}


//@description 创建{{.Table.TableComment}}
//@tags {{.Table.TableName}}{{if .Table.IsAuth}}
// @Security JWT{{end}}
// @Accept       json
// @Produce      json
//@param req    body  vo.Create{{.Table.TableName}}Request  false "入参req"
//@success 200 {object} vo.Create{{.Table.TableName}}Response
//@middleware
// @router /{{.Table.Uri}} [post]
func (c *{{.Table.TableName}}Controller) Create(ctx context.Context, req *vo.Create{{.Table.TableName}}Request) (*vo.Create{{.Table.TableName}}Response, error) {
	return c.{{.Table.Uri}}Service.Create(ctx, req)
}


//@description 批量删除{{.Table.TableComment}}
//@tags {{.Table.TableName}}{{if .Table.IsAuth}}
// @Security JWT{{end}}
// @Accept       json
// @Produce      json
//@param req    body  vo.Delete{{.Table.TableName}}Request  false  "入参req"
//@success 200  {object} vo.Delete{{.Table.TableName}}Response
//@middleware
// @router /{{.Table.Uri}} [delete]
func (c *{{.Table.TableName}}Controller) DeleteBatch(ctx context.Context, req *vo.Delete{{.Table.TableName}}Request) (int64, error) {
	return c.{{.Table.Uri}}Service.Delete(ctx, req)
}


//@description 更新{{.Table.TableComment}}
// @tags {{.Table.TableName}}{{if .Table.IsAuth}}
// @Security JWT{{end}}
// @Accept       json
// @Produce      json
//@param req    body  vo.Update{{.Table.TableName}}Request  false  "入参req"
//@success 200 {object} vo.Update{{.Table.TableName}}Response
//@middleware
// @router /{{.Table.Uri}}/:id [put]
func (c *{{.Table.TableName}}Controller) Update(ctx context.Context, req *vo.Update{{.Table.TableName}}Request) (*vo.Update{{.Table.TableName}}Response, error) {
	return c.{{.Table.Uri}}Service.Update(ctx, req)
}


{{end}}
