package controller

import (
    "github.com/gin-gonic/gin"
    "go-web-mini/vo"
    "go-web-mini/service"

)

//@description branch_tab管理
//@middleware auth
//@router /api
type BranchTabController struct {
   branchTabService service.BranchTabService
}

//@description 查询单个branch_tab
//@middleware auth,transition
// @router /branchTab [get]
func (c *BranchTabController) List(ctx *gin.Context, req *vo.ListBranchTabRequest) (interface{}, error) {
	return c.branchTabService.List(ctx, req)
}


//@description 查询branch_tab列表
//@middleware auth,transition
// @router /branchTab/:id [get]
func (c *BranchTabController) GetById(ctx *gin.Context, req *vo.GetBranchTabRequest) (interface{}, error) {
	return c.branchTabService.GetById(ctx, req)
}


//@description 创建branch_tab
//@middleware auth,transition
// @router /branchTab [post]
func (c *BranchTabController) Create(ctx *gin.Context, req *vo.CreateBranchTabRequest) (interface{}, error) {
	return c.branchTabService.Create(ctx, req)
}

//@description 删除单个branch_tab
//@middleware auth,transition
// @router /branchTab/:id [delete]
func (c *BranchTabController) Delete(ctx *gin.Context, req *vo.DeleteBranchTabRequest) (interface{}, error) {
	return c.branchTabService.Delete(ctx, req)
}

//@description 批量删除branch_tab
//@middleware auth,transition
// @router /branchTab [delete]
func (c *BranchTabController) DeleteBatch(ctx *gin.Context, req *vo.DeleteBranchTabRequest) (interface{}, error) {
	return c.branchTabService.Delete(ctx, req)
}


//@description 更新branch_tab
//@middleware auth,transition
// @router /branchTab/:id [put]
func (c *BranchTabController) Update(ctx *gin.Context, req *vo.UpdateBranchTabRequest) (interface{}, error) {
	return c.branchTabService.Update(ctx, req)
}


