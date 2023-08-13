package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/service"
	"go-web-mini/vo"
)

//@description 模块配置表管理

//@router /api
type ModuleController struct {
	moduleService service.ModuleService
}

//@description 查询模块配置表
//@tags Module}
// @Accept       json
// @Produce      json
//@param req    query  vo.ListModuleRequest  false  "入参req"
//@success 200 {object} vo.ListModuleResponse
//@middleware auth,transition
// @router /module [get]
func (c *ModuleController) List(ctx *gin.Context, req *vo.ListModuleRequest) (interface{}, error) {
	return c.moduleService.List(ctx, req)
}

//@description 查询单个模块配置表列表
//@tags Module
// @Accept       json
// @Produce      json
//@param req    query  vo.GetModuleRequest  false  "入参req"
//@success 200 {object} vo.GetModuleResponse
//@middleware auth,transition
// @router /module/:id [get]
func (c *ModuleController) GetById(ctx *gin.Context, req *vo.GetModuleRequest) (interface{}, error) {
	return c.moduleService.GetById(ctx, req)
}

//@description 创建模块配置表
//@tags Module
// @Accept       json
// @Produce      json
//@param req    body  vo.CreateModuleRequest  false "入参req"
//@success 200 {object} vo.CreateModuleResponse
//@middleware auth,transition
// @router /module [post]
func (c *ModuleController) Create(ctx *gin.Context, req *vo.CreateModuleRequest) (interface{}, error) {
	return c.moduleService.Create(ctx, req)
}

//@description 批量删除模块配置表
//@tags Module
// @Accept       json
// @Produce      json
//@param req    body  vo.DeleteModuleRequest  false  "入参req"
//@success 200  {object} vo.DeleteModuleResponse
//@middleware auth,transition
// @router /module [delete]
func (c *ModuleController) DeleteBatch(ctx *gin.Context, req *vo.DeleteModuleRequest) (interface{}, error) {
	return c.moduleService.Delete(ctx, req)
}

//@description 更新模块配置表
// @tags Module
// @Accept       json
// @Produce      json
//@param req    body  vo.UpdateModuleRequest  false  "入参req"
//@success 200 {object} vo.UpdateModuleResponse
//@middleware auth,transition
// @router /module/:id [put]
func (c *ModuleController) Update(ctx *gin.Context, req *vo.UpdateModuleRequest) (interface{}, error) {
	return c.moduleService.Update(ctx, req)
}
