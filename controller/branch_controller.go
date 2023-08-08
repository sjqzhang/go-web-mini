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
//@tags Branch
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.ListBranchRequest  false  "入参req"
//@success 200 {object} vo.ListBranchResponse
//@middleware auth,transition
// @router /branch [get]
func (c *BranchController) List(ctx *gin.Context, req *vo.ListBranchRequest) (interface{}, error) {
	return c.branchService.List(ctx, req)
}


//@description 查询branch列表
//@tags Branch
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.GetBranchRequest  false  "入参req"
//@success 200 {object} vo.GetBranchResponse
//@middleware auth,transition
// @router /branch/:id [get]
func (c *BranchController) GetById(ctx *gin.Context, req *vo.GetBranchRequest) (interface{}, error) {
	return c.branchService.GetById(ctx, req)
}


//@description 创建branch
//@tags Branch
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    body  vo.CreateBranchRequest  false "入参req"
//@success 200 {object} vo.CreateBranchResponse
//@middleware auth,transition
// @router /branch [post]
func (c *BranchController) Create(ctx *gin.Context, req *vo.CreateBranchRequest) (interface{}, error) {
	return c.branchService.Create(ctx, req)
}


//@description 批量删除branch
//@tags Branch
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    body  vo.DeleteBranchRequest  false  "入参req"
//@success 200  {object} vo.DeleteBranchResponse
//@middleware auth,transition
// @router /branch [delete]
func (c *BranchController) DeleteBatch(ctx *gin.Context, req *vo.DeleteBranchRequest) (interface{}, error) {
	return c.branchService.Delete(ctx, req)
}


//@description 更新branch
// @tags Branch
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    body  vo.UpdateBranchRequest  false  "入参req"
//@success 200 {object} vo.UpdateBranchResponse
//@middleware auth,transition
// @router /branch/:id [put]
func (c *BranchController) Update(ctx *gin.Context, req *vo.UpdateBranchRequest) (interface{}, error) {
	return c.branchService.Update(ctx, req)
}


