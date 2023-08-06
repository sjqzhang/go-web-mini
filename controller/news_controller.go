package controller

import (
    "github.com/gin-gonic/gin"
    "go-web-mini/vo"
    "go-web-mini/service"

)

//@description news管理
//@middleware auth
//@router /api
type NewsController struct {
   newsService service.NewsService
}

//@description 查询单个news
//@middleware auth,transition
// @router /news [get]
func (c *NewsController) List(ctx *gin.Context, req *vo.ListNewsRequest) (interface{}, error) {
	return c.newsService.List(ctx, req)
}


//@description 查询news列表
//@middleware auth,transition
// @router /news/:id [get]
func (c *NewsController) GetById(ctx *gin.Context, req *vo.GetNewsRequest) (interface{}, error) {
	return c.newsService.GetById(ctx, req)
}


//@description 创建news
//@middleware auth,transition
// @router /news [post]
func (c *NewsController) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (interface{}, error) {
	return c.newsService.Create(ctx, req)
}

//@description 删除单个news
//@middleware auth,transition
// @router /news/:id [delete]
func (c *NewsController) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (interface{}, error) {
	return c.newsService.Delete(ctx, req)
}

//@description 批量删除news
//@middleware auth,transition
// @router /news [delete]
func (c *NewsController) DeleteBatch(ctx *gin.Context, req *vo.DeleteNewsRequest) (interface{}, error) {
	return c.newsService.Delete(ctx, req)
}


//@description 更新news
//@middleware auth,transition
// @router /news/:id [put]
func (c *NewsController) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (interface{}, error) {
	return c.newsService.Update(ctx, req)
}


