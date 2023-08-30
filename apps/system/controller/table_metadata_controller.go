package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/apps/system/service"
	"go-web-mini/apps/system/vo"
)

//@description table_metadata管理

//@router /api
type TableMetadataController struct {
	tableMetadataService service.TableMetadataService
}

//@description 查询table_metadata
//@tags TableMetadata
// @Accept       json
// @Produce      json
//@param req    query  vo.ListTableMetadataRequest  false  "入参req"
//@success 200 {object} vo.ListTableMetadataResponse
//@middleware
// @router /tableMetadata [get]
func (c *TableMetadataController) List(ctx *gin.Context, req *vo.ListTableMetadataRequest) (interface{}, error) {
	return c.tableMetadataService.List(ctx, req)
}

//@description 查询单个table_metadata
//@tags TableMetadata
// @Accept       json
// @Produce      json
//@param req    query  vo.GetTableMetadataRequest  false  "入参req"
//@success 200 {object} vo.GetTableMetadataResponse
//@middleware
// @router /tableMetadata/:id [get]
func (c *TableMetadataController) GetById(ctx *gin.Context, req *vo.GetTableMetadataRequest) (interface{}, error) {
	return c.tableMetadataService.GetById(ctx, req)
}

//@description 创建table_metadata
//@tags TableMetadata
// @Accept       json
// @Produce      json
//@param req    body  vo.CreateTableMetadataRequest  false "入参req"
//@success 200 {object} vo.CreateTableMetadataResponse
//@middleware
// @router /tableMetadata [post]
func (c *TableMetadataController) Create(ctx *gin.Context, req *vo.CreateTableMetadataRequest) (interface{}, error) {
	return c.tableMetadataService.Create(ctx, req)
}

//@description 批量删除table_metadata
//@tags TableMetadata
// @Accept       json
// @Produce      json
//@param req    body  vo.DeleteTableMetadataRequest  false  "入参req"
//@success 200  {object} vo.DeleteTableMetadataResponse
//@middleware
// @router /tableMetadata [delete]
func (c *TableMetadataController) DeleteBatch(ctx *gin.Context, req *vo.DeleteTableMetadataRequest) (interface{}, error) {
	return c.tableMetadataService.Delete(ctx, req)
}

//@description 更新table_metadata
// @tags TableMetadata
// @Accept       json
// @Produce      json
//@param req    body  vo.UpdateTableMetadataRequest  false  "入参req"
//@success 200 {object} vo.UpdateTableMetadataResponse
//@middleware
// @router /tableMetadata/:id [put]
func (c *TableMetadataController) Update(ctx *gin.Context, req *vo.UpdateTableMetadataRequest) (interface{}, error) {
	return c.tableMetadataService.Update(ctx, req)
}
