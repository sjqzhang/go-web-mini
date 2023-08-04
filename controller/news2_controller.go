package controller

//@middleware auth
//@router /api/news
//type News2Controller struct {
//	newsService service.NewsService
//}
//
//// 查询接口
////@middleware auth,transition
//// @router /list [get]
//func (c *News2Controller) List(ctx *gin.Context, req *vo.ListNewsRequest) (interface{}, error) {
//	return c.newsService.List(ctx, req)
//}
//
//// 创建接口
////@middleware auth,transition
//// @router /create [post]
//func (c *News2Controller) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (interface{}, error) {
//	return c.newsService.Create(ctx, req)
//}
//
//// 删除接口
////@middleware auth,transition
//// @router /delete/:id [delete]
//func (c *News2Controller) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (interface{}, error) {
//	return c.newsService.Delete(ctx, req)
//}
//
//// 更新接口
////@middleware auth,transition
//// @router /update [post]
//func (c *News2Controller) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (interface{}, error) {
//	return c.newsService.Update(ctx, req)
//}
