{{define "controller"}}package controller

import (
    "github.com/gin-gonic/gin"
    "{{.ModuleName}}/vo"
    "{{.ModuleName}}/service"

)

//@middleware auth
//@router /api
type {{.Table.TableName}}Controller struct {
   {{.Table.TableName}}Service service.{{.Table.TableName}}Service
}
/*
// {{.Table.TableName}}QueryPage {{.Table.TableComment}}分页查询
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}QueryPage(ctx *gin.Context) {

    var param dto.{{.Table.TableName}}PageDTO

    // 绑定参数
    err := ctx.ShouldBindJSON(&param)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
        return
    }

    // 查询新增
    ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}QueryPage(param)})
}
*/

// 查询接口
//@middleware auth,transition
// @router /list [get]
func (c *{{.Table.TableName}}Controller) List(ctx *gin.Context, req *vo.List{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.TableName}}Service.List(ctx, req)
}

// 创建接口
//@middleware auth,transition
// @router /create [post]
func (c *{{.Table.TableName}}Controller) Create(ctx *gin.Context, req *vo.Create{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.TableName}}Service.Create(ctx, req)
}

// 删除接口
//@middleware auth,transition
// @router /delete/:id [delete]
func (c *{{.Table.TableName}}Controller) Delete(ctx *gin.Context, req *vo.Delete{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.TableName}}Service.Delete(ctx, req)
}

// 更新接口
//@middleware auth,transition
// @router /update [post]
func (c *{{.Table.TableName}}Controller) Update(ctx *gin.Context, req *vo.Update{{.Table.TableName}}Request) (interface{}, error) {
	return c.{{.Table.TableName}}Service.Update(ctx, req)
}


{{end}}
