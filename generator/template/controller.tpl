{{define "controller"}}package controller

import (
    "github.com/gin-gonic/gin"
    "{{.ModuleName}}/vo"
    "{{.ModuleName}}/service"

)

//@middleware auth
//@router /api
type {{.Table.TableName}}Controller struct {
   {{.Table.Uri}}Service service.{{.Table.TableName}}Service
}

// 查询接口
//@middleware auth,transition
// @router /{{.Table.Uri}} [get]
func (c *{{.Table.TableName}}Controller) List(ctx *gin.Context, req *vo.List{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.Uri}}Service.List(ctx, req)
}


// 查询接口
//@middleware auth,transition
// @router /{{.Table.Uri}}/:id [get]
func (c *{{.Table.TableName}}Controller) GetById(ctx *gin.Context, req *vo.Get{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.Uri}}Service.GetById(ctx, req)
}


// 创建接口
//@middleware auth,transition
// @router /{{.Table.Uri}} [post]
func (c *{{.Table.TableName}}Controller) Create(ctx *gin.Context, req *vo.Create{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.Uri}}Service.Create(ctx, req)
}

// 删除接口
//@middleware auth,transition
// @router /{{.Table.Uri}}/:id [delete]
func (c *{{.Table.TableName}}Controller) Delete(ctx *gin.Context, req *vo.Delete{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.Uri}}Service.Delete(ctx, req)
}

// 更新接口
//@middleware auth,transition
// @router /{{.Table.Uri}}/:id [put]
func (c *{{.Table.TableName}}Controller) Update(ctx *gin.Context, req *vo.Update{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.Uri}}Service.Update(ctx, req)
}


{{end}}
