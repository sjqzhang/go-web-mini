package controller

import (
    "github.com/gin-gonic/gin"
    "go-web-mini/vo"
    "go-web-mini/service"

)

//@middleware auth
//@router /api
type NewsController struct {
   newsService service.NewsService
}

// 查询接口
//@middleware auth,transition
// @router /news [get]
func (c *NewsController) List(ctx *gin.Context, req *vo.ListNewsRequest) (interface{}, error) {
	return c.newsService.List(ctx, req)
}


// 查询接口
//@middleware auth,transition
// @router /news/:id [get]
func (c *NewsController) GetById(ctx *gin.Context, req *vo.GetNewsRequest) (interface{}, error) {
	return c.newsService.GetById(ctx, req)
}


// 创建接口
//@middleware auth,transition
// @router /news [post]
func (c *NewsController) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (interface{}, error) {
	return c.newsService.Create(ctx, req)
}

// 删除接口
//@middleware auth,transition
// @router /news/:id [delete]
func (c *NewsController) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (interface{}, error) {
	return c.newsService.Delete(ctx, req)
}

// 更新接口
//@middleware auth,transition
// @router /news/:id [put]
func (c *NewsController) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (interface{}, error) {
	return c.newsService.Update(ctx, req)
}


