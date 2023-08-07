package controller

import (
    "github.com/gin-gonic/gin"
    "go-web-mini/vo"
    "go-web-mini/service"

)

//@description branch管理
//@middleware auth
//@router /api
type BranchController struct {
   branchService service.BranchService
}

//@description 查询单个branch
//@middleware auth,transition
// @router /branch [get]
func (c *BranchController) List(ctx *gin.Context, req *vo.ListBranchRequest) (interface{}, error) {
	return c.branchService.List(ctx, req)
}


//@description 查询branch列表
//@middleware auth,transition
// @router /branch/:id [get]
func (c *BranchController) GetById(ctx *gin.Context, req *vo.GetBranchRequest) (interface{}, error) {
	return c.branchService.GetById(ctx, req)
}


//@description 创建branch
//@middleware auth,transition
// @router /branch [post]
func (c *BranchController) Create(ctx *gin.Context, req *vo.CreateBranchRequest) (interface{}, error) {
	return c.branchService.Create(ctx, req)
}

//@description 删除单个branch
//@middleware auth,transition
// @router /branch/:id [delete]
func (c *BranchController) Delete(ctx *gin.Context, req *vo.DeleteBranchRequest) (interface{}, error) {
	return c.branchService.Delete(ctx, req)
}

//@description 批量删除branch
//@middleware auth,transition
// @router /branch [delete]
func (c *BranchController) DeleteBatch(ctx *gin.Context, req *vo.DeleteBranchRequest) (interface{}, error) {
	return c.branchService.Delete(ctx, req)
}


//@description 更新branch
//@middleware auth,transition
// @router /branch/:id [put]
func (c *BranchController) Update(ctx *gin.Context, req *vo.UpdateBranchRequest) (interface{}, error) {
	return c.branchService.Update(ctx, req)
}


