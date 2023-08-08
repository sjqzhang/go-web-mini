package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/service"
	"go-web-mini/vo"
)

//@description 新闻管理
//@middleware auth
//@router /api
type NewsController struct {
	newsService service.NewsService
}

//@description 查询单个新闻
//@tags News
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.ListNewsRequest  false  "入参req"
//@success 200 {object} vo.ListNewsResponse
//@middleware auth,transition
// @router /news [get]
func (c *NewsController) List(ctx *gin.Context, req *vo.ListNewsRequest) (interface{}, error) {
	return c.newsService.List(ctx, req)
}

//@description 查询新闻列表
//@tags News
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.GetNewsRequest  false  "入参req"
//@success 200 {object} vo.GetNewsResponse
//@middleware auth,transition
// @router /news/:id [get]
func (c *NewsController) GetById(ctx *gin.Context, req *vo.GetNewsRequest) (interface{}, error) {
	return c.newsService.GetById(ctx, req)
}

//@description 创建新闻
//@tags News
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.CreateNewsRequest  false "入参req"
//@success 200 {object} vo.CreateNewsResponse
//@middleware auth,transition
// @router /news [post]
func (c *NewsController) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (interface{}, error) {
	return c.newsService.Create(ctx, req)
}

//@description 批量删除新闻
//@tags News
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.DeleteNewsRequest  false  "入参req"
//@success 200  {object} vo.DeleteNewsResponse
//@middleware auth,transition
// @router /news [delete]
func (c *NewsController) DeleteBatch(ctx *gin.Context, req *vo.DeleteNewsRequest) (interface{}, error) {
	return c.newsService.Delete(ctx, req)
}

//@description 更新新闻
// @tags News
// @Security JWT
// @Accept       json
// @Produce      json
//@param req    query  vo.UpdateNewsRequest  false  "入参req"
//@success 200 {object} vo.UpdateNewsResponse
//@middleware auth,transition
// @router /news/:id [put]
func (c *NewsController) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (interface{}, error) {
	return c.newsService.Update(ctx, req)
}
