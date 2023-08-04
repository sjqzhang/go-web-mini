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
	{{.Table.Uri}}Repository repository.{{.Table.TableName}}Repository
}

func (s *{{.Table.TableName}}Service) List(ctx context.Context, req *vo.List{{.Table.TableName}}Request) ([]*vo.{{.Table.TableName}}, error) {
	var query model.{{.Table.TableName}}Query
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
    var resp []*vo.News
    objs,err:= s.{{.Table.Uri}}Repository.List(ctx, &query)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, objs)
    return resp, err
}

func (s *{{.Table.TableName}}Service) Create(ctx *gin.Context, req *vo.Create{{.Table.TableName}}Request) (*vo.{{.Table.TableName}}, error) {
	var obj model.{{.Table.TableName}}
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp *vo.{{.Table.TableName}}
	_,err= s.{{.Table.Uri}}Repository.Create(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return resp, err
}


func (s *{{.Table.TableName}}Service) Update(ctx *gin.Context, req *vo.Update{{.Table.TableName}}Request) (*vo.{{.Table.TableName}}, error) {
	var obj model.News
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp *vo.{{.Table.TableName}}
	_,err= s.{{.Table.Uri}}Repository.Update(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return resp, err
}


func (s *{{.Table.TableName}}Service) Delete(ctx *gin.Context, req *vo.Delete{{.Table.TableName}}Request) (int64, error) {
	var obj model.{{.Table.TableName}}
	err := copier.Copy(&obj, req)
	if err != nil {
		return 0, err
	}
	return s.{{.Table.Uri}}Repository.Delete(ctx, &obj)
}


{{end}}

