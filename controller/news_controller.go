package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/service"
	"go-web-mini/vo"
)

//@middleware auth
//@router /api/news
type NewsController struct {
   NewsService service.NewsService
}
/*
// NewsQueryPage 分页查询
func (NewsController NewsController) NewsQueryPage(ctx *gin.Context) {

    var param dto.NewsPageDTO

    // 绑定参数
    err := ctx.ShouldBindJSON(&param)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
        return
    }

    // 查询新增
    ctx.JSON(http.StatusOK, gin.H{"data": service.NewsQueryPage(param)})
}
*/

// 查询接口
//@middleware auth,transition
// @router /list [get]
func (c *NewsController) List(ctx *gin.Context, req *vo.ListNewsRequest) (interface{}, error) {
	return c.NewsService.List(ctx, req)
}

// 创建接口
//@middleware auth,transition
// @router /create [post]
func (c *NewsController) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (interface{}, error) {
	return c.NewsService.Create(ctx, req)
}

// 删除接口
//@middleware auth,transition
// @router /delete/:id [delete]
func (c *NewsController) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (interface{}, error) {
	return c.NewsService.Delete(ctx, req)
}

// 更新接口
//@middleware auth,transition
// @router /update [post]
func (c *NewsController) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (interface{}, error) {
	return c.NewsService.Update(ctx, req)
}


