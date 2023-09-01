package controller

import (
	"errors"
	"go-web-mini/apps/workflow/errcode"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"

	"github.com/gin-gonic/gin"
	"net/http"
)

//api.POST("/schema/infos", GetSchemaInfos)
//api.POST("/transitions/ids", GetInstanceTransitionsByIds)
//api.POST("/schema/upload", UploadSchemaFile)
//api.POST("/schema/reset", ResetInstanceSchemaCode)


//@Description schema信息
//@router /bpmn/api
type SchemaController struct{

}

//@description 获取schema信息
//@router /schema/infos [post]
func (cc *SchemaController)GetSchemaInfos(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetSchemaReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	serviceSchema := service.NewSchemeService()
	definitions, e := serviceSchema.GetSchemaTDefinitions(ctx, req.SchemaCodes)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	var schemaInfos []vo.SchemaInfo
	for _, definition := range definitions {
		serviceTasks := convertServiceTasks(definition.Process.ServiceTasks)
		commands := convertCommands(definition.Process.IntermediateCatchEvent)
		schemaInfos = append(schemaInfos, vo.SchemaInfo{
			ServiceTasks: serviceTasks,
			Commands:     commands,
		})
	}
	resp := vo.GetSchemaInfosResp{
		SchemaInfos: schemaInfos,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//@description 上传schema文件
//@router /schema/upload [post]
func (cc *SchemaController)UploadSchemaFile(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.SchemaFileReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	if len(req.SchemaCode) == 0 || len(req.Schema) == 0 {
		c.JSON(http.StatusOK, vo.NewResp(errcode.Wrap(errcode.UploadFileError, errors.New("has no schema code or content")), nil))
		return
	}
	serviceSchema := service.NewSchemeService()
	err := serviceSchema.UploadSchemaFile(ctx, req.SchemaCode, req.Schema)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResp(err, nil))
		return
	}
	resp := vo.SchemaFileResp{
		SchemaCode: req.SchemaCode,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//@description 重置schema文件
//@router /schema/reset [post]
func (cc *SchemaController)ResetInstanceSchemaCode(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.SchemaFileReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	serviceSchema := service.NewSchemeService()
	e := serviceSchema.ResetSchemaCode(ctx, req.SchemaCode, req.Schema)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	resp := vo.SchemaFileResp{
		SchemaCode: req.SchemaCode,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//func UploadSchemaFile(c *gin.Context) {
//	ctx := c.Request.Context()
//	//var req vo.SchemaFileReq
//	//if err := easyBind(c, &req); err != nil {
//	//	return
//	//}
//	fileHeader, e := c.FormFile("file")
//	if e != nil {
//		err := errcode.Wrap(errcode.UploadFileError, e)
//		c.JSON(http.StatusOK, vo.NewResp(err, nil))
//		return
//	}
//	schemaCode, flag := c.GetPostForm("schema_code")
//	if !flag {
//		c.JSON(http.StatusOK, vo.NewResp(errcode.Wrap(errcode.UploadFileError, errors.New("has no schema code")), nil))
//		return
//	}
//	file, e := fileHeader.Open()
//	if e != nil {
//		err := errcode.Wrap(errcode.UploadFileError, e)
//		c.JSON(http.StatusOK, vo.NewResp(err, nil))
//		return
//	}
//	defer file.Close()
//	bytes, e := ioutil.ReadAll(file)
//	if e != nil {
//		err := errcode.Wrap(errcode.UploadFileError, e)
//		c.JSON(http.StatusOK, vo.NewResp(err, nil))
//		return
//	}
//	schemaStr := string(bytes)
//	if len(schemaStr) == 0 {
//		err := errcode.Wrap(errcode.UploadFileError, errors.New("file content is empty"))
//		c.JSON(http.StatusOK, vo.NewResp(err, nil))
//		return
//	}
//	serviceSchema := service.NewSchemeService()
//	err := serviceSchema.UploadSchemaFile(ctx, schemaCode, schemaStr)
//	if err != nil {
//		c.JSON(http.StatusOK, vo.NewResp(err, nil))
//		return
//	}
//	resp := vo.SchemaFileResp{
//		SchemaCode: schemaCode,
//	}
//	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
//}
