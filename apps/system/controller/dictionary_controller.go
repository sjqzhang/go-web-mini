package controller

import (
	"context"
	"go-web-mini/apps/system/service"
	"go-web-mini/apps/system/vo"
)

//@description 字典管理

//@router /api/system
type DictionaryController struct {
	dictionaryService service.DictionaryService
}

//@description 查询dictionary
//@tags Dictionary
// @Accept       json
// @Produce      json
//@param req    query  vo.ListDictionaryRequest  false  "入参req"
//@success 200 {object} vo.ListDictionaryResponse
//@middleware
// @router /dictionary/pager [get]
func (c *DictionaryController) ListForPager(ctx context.Context, req *vo.ListForPagerDictionaryRequest) (*vo.ListForPagerDictionaryResponse, error) {
	return c.dictionaryService.ListForPager(ctx, req)
}

//@description 查询dictionary
//@tags Dictionary
// @Accept       json
// @Produce      json
//@param req    query  vo.ListDictionaryRequest  false  "入参req"
//@success 200 {object} vo.ListDictionaryResponse
//@middleware
// @router /dictionary [get]
func (c *DictionaryController) List(ctx context.Context, req *vo.ListDictionaryRequest) (*vo.ListDictionaryResponse, error) {
	return c.dictionaryService.List(ctx, req)
}

//@description 查询单个dictionary
//@tags Dictionary
// @Accept       json
// @Produce      json
//@param req    query  vo.GetDictionaryRequest  false  "入参req"
//@success 200 {object} vo.GetDictionaryResponse
//@middleware
// @router /dictionary/:id [get]
func (c *DictionaryController) GetById(ctx context.Context, req *vo.GetDictionaryRequest) (*vo.DictionaryResponse, error) {
	return c.dictionaryService.GetById(ctx, req)
}

//@description 创建dictionary
//@tags Dictionary
// @Accept       json
// @Produce      json
//@param req    body  vo.CreateDictionaryRequest  false "入参req"
//@success 200 {object} vo.CreateDictionaryResponse
//@middleware
// @router /dictionary [post]
func (c *DictionaryController) Create(ctx context.Context, req *vo.CreateDictionaryRequest) (*vo.CreateDictionaryResponse, error) {
	return c.dictionaryService.Create(ctx, req)
}

//@description 批量删除dictionary
//@tags Dictionary
// @Accept       json
// @Produce      json
//@param req    body  vo.DeleteDictionaryRequest  false  "入参req"
//@success 200  {object} vo.DeleteDictionaryResponse
//@middleware
// @router /dictionary [delete]
func (c *DictionaryController) DeleteBatch(ctx context.Context, req *vo.DeleteDictionaryRequest) (int64, error) {
	return c.dictionaryService.Delete(ctx, req)
}

//@description 更新dictionary
// @tags Dictionary
// @Accept       json
// @Produce      json
//@param req    body  vo.UpdateDictionaryRequest  false  "入参req"
//@success 200 {object} vo.UpdateDictionaryResponse
//@middleware
// @router /dictionary/:id [put]
func (c *DictionaryController) Update(ctx context.Context, req *vo.UpdateDictionaryRequest) (*vo.UpdateDictionaryResponse, error) {
	return c.dictionaryService.Update(ctx, req)
}
