package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/service"
	"go-web-mini/vo"
)

//@description branch管理
//@middleware auth
//@router /api
type BranchController struct {
	branchService service.BranchService
}

//@description 查询单个branch
//@tags branch管理
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.ListBranchRequest  false  "name search by q"  Format(email)
//@success 200 {object} vo.PagerBranch
//@middleware auth,transition
// @router /branch [get]
func (c *BranchController) List(ctx *gin.Context, req *vo.ListBranchRequest) (interface{}, error) {
	return c.branchService.List(ctx, req)
}

//@description 查询branch列表
//@tags branch管理
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.GetBranchRequest  false  "name search by q"  Format(email)
//@success 200 {object} vo.Branch
//@middleware auth,transition
// @router /branch/:id [get]
func (c *BranchController) GetById(ctx *gin.Context, req *vo.GetBranchRequest) (interface{}, error) {
	return c.branchService.GetById(ctx, req)
}

//@description 创建branch
//@tags branch管理
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.CreateBranchRequest  false  "name search by q"  Format(email)
//@success 200 {object} vo.Branch
//@middleware auth,transition
// @router /branch [post]
func (c *BranchController) Create(ctx *gin.Context, req *vo.CreateBranchRequest) (interface{}, error) {
	return c.branchService.Create(ctx, req)
}

//@description 删除单个branch
//@tags branch管理
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.DeleteBranchRequest  false  "name search by q"  Format(email)
//@success 200 int int64
//@middleware auth,transition
// @router /branch/:id [delete]
func (c *BranchController) Delete(ctx *gin.Context, req *vo.DeleteBranchRequest) (interface{}, error) {
	return c.branchService.Delete(ctx, req)
}

//@description 批量删除branch
//@tags branch管理
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.DeleteBranchRequest  false  "name search by q"  Format(email)
//@success 200  int int64
//@middleware auth,transition
// @router /branch [delete]
func (c *BranchController) DeleteBatch(ctx *gin.Context, req *vo.DeleteBranchRequest) (interface{}, error) {
	return c.branchService.Delete(ctx, req)
}

//@description 更新branch
//@tags branch管理
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.UpdateBranchRequest  false  "name search by q"  Format(email)
//@success 200 {object} vo.Branch
//@middleware auth,transition
// @router /branch/:id [put]
func (c *BranchController) Update(ctx *gin.Context, req *vo.UpdateBranchRequest) (interface{}, error) {
	return c.branchService.Update(ctx, req)
}
