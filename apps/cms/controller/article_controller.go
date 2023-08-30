package controller

import (
	"context"
	"go-web-mini/apps/cms/service"
	"go-web-mini/apps/cms/vo"
)

//@description article管理

//@router /api
type ArticleController struct {
	articleService service.ArticleService
}

//@description 查询article
//@tags Article
// @Accept       json
// @Produce      json
//@param req    query  vo.ListArticleRequest  false  "入参req"
//@success 200 {object} vo.ListArticleResponse
//@middleware
// @router /article [get]
func (c *ArticleController) List(ctx context.Context, req *vo.ListArticleRequest) (*vo.ListArticleResponse, error) {
	return c.articleService.List(ctx, req)
}

//@description 查询单个article
//@tags Article
// @Accept       json
// @Produce      json
//@param req    query  vo.GetArticleRequest  false  "入参req"
//@success 200 {object} vo.GetArticleResponse
//@middleware
// @router /article/:id [get]
func (c *ArticleController) GetById(ctx context.Context, req *vo.GetArticleRequest) (*vo.ArticleResponse, error) {
	return c.articleService.GetById(ctx, req)
}

//@description 创建article
//@tags Article
// @Accept       json
// @Produce      json
//@param req    body  vo.CreateArticleRequest  false "入参req"
//@success 200 {object} vo.CreateArticleResponse
//@middleware
// @router /article [post]
func (c *ArticleController) Create(ctx context.Context, req *vo.CreateArticleRequest) (*vo.CreateArticleResponse, error) {
	return c.articleService.Create(ctx, req)
}

//@description 批量删除article
//@tags Article
// @Accept       json
// @Produce      json
//@param req    body  vo.DeleteArticleRequest  false  "入参req"
//@success 200  {object} vo.DeleteArticleResponse
//@middleware
// @router /article [delete]
func (c *ArticleController) DeleteBatch(ctx context.Context, req *vo.DeleteArticleRequest) (int64, error) {
	return c.articleService.Delete(ctx, req)
}

//@description 更新article
// @tags Article
// @Accept       json
// @Produce      json
//@param req    body  vo.UpdateArticleRequest  false  "入参req"
//@success 200 {object} vo.UpdateArticleResponse
//@middleware
// @router /article/:id [put]
func (c *ArticleController) Update(ctx context.Context, req *vo.UpdateArticleRequest) (*vo.UpdateArticleResponse, error) {
	return c.articleService.Update(ctx, req)
}
