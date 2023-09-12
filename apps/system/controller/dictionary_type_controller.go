package controller

import (
	"context"
	"go-web-mini/apps/system/service"
	"go-web-mini/apps/system/vo"
)

//@description 字典类型管理

//@router /api/system
type DictionaryTypeController struct {
	dictionaryTypeService service.DictionaryTypeService
}

//@description 查询dictionary_type
//@tags DictionaryType
// @Accept       json
// @Produce      json
//@param req    query  vo.ListDictionaryTypeRequest  false  "入参req"
//@success 200 {object} vo.ListDictionaryTypeResponse
//@middleware
// @router /dictionaryType/pager [get]
func (c *DictionaryTypeController) ListForPager(ctx context.Context, req *vo.ListForPagerDictionaryTypeRequest) (*vo.ListForPagerDictionaryTypeResponse, error) {
	return c.dictionaryTypeService.ListForPager(ctx, req)
}

//@description 查询dictionary_type
//@tags DictionaryType
// @Accept       json
// @Produce      json
//@param req    query  vo.ListDictionaryTypeRequest  false  "入参req"
//@success 200 {object} vo.ListDictionaryTypeResponse
//@middleware
// @router /dictionaryType [get]
func (c *DictionaryTypeController) List(ctx context.Context, req *vo.ListDictionaryTypeRequest) (*vo.ListDictionaryTypeResponse, error) {
	return c.dictionaryTypeService.List(ctx, req)
}

//@description 查询单个dictionary_type
//@tags DictionaryType
// @Accept       json
// @Produce      json
//@param req    query  vo.GetDictionaryTypeRequest  false  "入参req"
//@success 200 {object} vo.GetDictionaryTypeResponse
//@middleware
// @router /dictionaryType/:id [get]
func (c *DictionaryTypeController) GetById(ctx context.Context, req *vo.GetDictionaryTypeRequest) (*vo.DictionaryTypeResponse, error) {
	return c.dictionaryTypeService.GetById(ctx, req)
}

//@description 创建dictionary_type
//@tags DictionaryType
// @Accept       json
// @Produce      json
//@param req    body  vo.CreateDictionaryTypeRequest  false "入参req"
//@success 200 {object} vo.CreateDictionaryTypeResponse
//@middleware
// @router /dictionaryType [post]
func (c *DictionaryTypeController) Create(ctx context.Context, req *vo.CreateDictionaryTypeRequest) (*vo.CreateDictionaryTypeResponse, error) {
	return c.dictionaryTypeService.Create(ctx, req)
}

//@description 批量删除dictionary_type
//@tags DictionaryType
// @Accept       json
// @Produce      json
//@param req    body  vo.DeleteDictionaryTypeRequest  false  "入参req"
//@success 200  {object} vo.DeleteDictionaryTypeResponse
//@middleware
// @router /dictionaryType [delete]
func (c *DictionaryTypeController) DeleteBatch(ctx context.Context, req *vo.DeleteDictionaryTypeRequest) (int64, error) {
	return c.dictionaryTypeService.Delete(ctx, req)
}

//@description 更新dictionary_type
// @tags DictionaryType
// @Accept       json
// @Produce      json
//@param req    body  vo.UpdateDictionaryTypeRequest  false  "入参req"
//@success 200 {object} vo.UpdateDictionaryTypeResponse
//@middleware
// @router /dictionaryType/:id [put]
func (c *DictionaryTypeController) Update(ctx context.Context, req *vo.UpdateDictionaryTypeRequest) (*vo.UpdateDictionaryTypeResponse, error) {
	return c.dictionaryTypeService.Update(ctx, req)
}
