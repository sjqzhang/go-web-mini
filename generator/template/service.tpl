{{define "service"}}package service

import (

    "context"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/copier"

    "{{.ModuleName}}/model"
    "{{.ModuleName}}/repository"
    "go-web-mini/vo"
)

/*
// {{.Table.TableName}}QueryPage {{.Table.TableComment}}分页查询
func {{.Table.TableName}}QueryPage(param dto.{{.Table.TableName}}PageDTO) []vo.{{.Table.TableName}}VO{

	return []vo.{{.Table.TableName}}VO{}
}
*/

type {{.Table.TableName}}Service struct {
	{{.Table.TableName}}Repository repository.{{.Table.TableName}}Repository
}

func (s *{{.Table.TableName}}Service) List(ctx context.Context, req *vo.List{{.Table.TableName}}Request) ([]*model.{{.Table.TableName}}, error) {
	var query model.{{.Table.TableName}}Query
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	return s.{{.Table.TableName}}Repository.List(ctx, &query)
}

func (s *{{.Table.TableName}}Service) Create(ctx *gin.Context, req *vo.Create{{.Table.TableName}}Request) (*model.{{.Table.TableName}}, error) {
	var obj model.{{.Table.TableName}}
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	return s.{{.Table.TableName}}Repository.Create(ctx, &obj)
}


func (s *{{.Table.TableName}}Service) Update(ctx *gin.Context, req *vo.Update{{.Table.TableName}}Request) (*model.{{.Table.TableName}}, error) {
	var obj model.News
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	return s.{{.Table.TableName}}Repository.Update(ctx, &obj)
}


func (s *{{.Table.TableName}}Service) Delete(ctx *gin.Context, req *vo.Delete{{.Table.TableName}}Request) (int64, error) {
	var obj model.{{.Table.TableName}}
	err := copier.Copy(&obj, req)
	if err != nil {
		return 0, err
	}
	return s.{{.Table.TableName}}Repository.Delete(ctx, &obj)
}


{{end}}

