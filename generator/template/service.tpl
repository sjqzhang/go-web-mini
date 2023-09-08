{{define "service"}}package service

import (

    "context"
    "github.com/jinzhu/copier"

    "{{.ModuleName}}/apps/{{.AppName}}/model"
    "{{.ModuleName}}/apps/{{.AppName}}/repository"
    "{{.ModuleName}}/apps/{{.AppName}}/vo"
)

/*
// {{.Table.TableName}}QueryPage {{.Table.TableComment}}分页查询
func {{.Table.TableName}}QueryPage(param dto.{{.Table.TableName}}PageDTO) []vo.{{.Table.TableName}}VO{

	return []vo.{{.Table.TableName}}VO{}
}
*/

type  I{{.Table.TableName}}Service interface {
    ListForPager(ctx context.Context, req *vo.ListForPager{{.Table.TableName}}Request) (*vo.ListForPager{{.Table.TableName}}Response, error)
    List(ctx context.Context, req *vo.List{{.Table.TableName}}Request) (*vo.List{{.Table.TableName}}Response, error)
    GetById(ctx context.Context, req *vo.Get{{.Table.TableName}}Request) (*vo.{{.Table.TableName}}Response, error)
    Create(ctx context.Context, req *vo.Create{{.Table.TableName}}Request) (*vo.Create{{.Table.TableName}}Response, error)
    Update(ctx context.Context, req *vo.Update{{.Table.TableName}}Request) (*vo.Update{{.Table.TableName}}Response, error)
    Delete(ctx context.Context, req *vo.Delete{{.Table.TableName}}Request) (int64, error)
}

type {{.Table.TableName}}Service struct {
	{{.Table.Uri}}Repository repository.{{.Table.TableName}}Repository
}

func (s *{{.Table.TableName}}Service) ListForPager(ctx context.Context, req *vo.ListForPager{{.Table.TableName}}Request) (*vo.ListForPager{{.Table.TableName}}Response, error) {
	var query model.{{.Table.TableName}}Query
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
    var resp vo.ListForPager{{.Table.TableName}}Response
    objs,err:= s.{{.Table.Uri}}Repository.ListForPager(ctx, &query)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, objs)
    return &resp, err
}

func (s *{{.Table.TableName}}Service) List(ctx context.Context, req *vo.List{{.Table.TableName}}Request) (*vo.List{{.Table.TableName}}Response, error) {
    var query model.{{.Table.TableName}}Query
    err := copier.Copy(&query, req)
    if err != nil {
        return nil, err
    }
    var resp vo.List{{.Table.TableName}}Response
    objs,err:= s.{{.Table.Uri}}Repository.List(ctx, &query)
    if err != nil {
        return nil, err
    }
    if err = copier.Copy(&resp.List, objs); err != nil {
        return nil, err
    }
    return &resp, err
}

func (s *{{.Table.TableName}}Service) GetById(ctx context.Context, req *vo.Get{{.Table.TableName}}Request) (*vo.{{.Table.TableName}}Response, error) {
    obj,err:= s.{{.Table.Uri}}Repository.GetById(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    var resp vo.{{.Table.TableName}}Response
    err = copier.Copy(&resp, obj)
    return &resp, err
}

func (s *{{.Table.TableName}}Service) Create(ctx context.Context, req *vo.Create{{.Table.TableName}}Request) (*vo.Create{{.Table.TableName}}Response, error) {
	var obj model.{{.Table.TableName}}
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.Create{{.Table.TableName}}Response
	_,err= s.{{.Table.Uri}}Repository.Create(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return &resp, err
}


func (s *{{.Table.TableName}}Service) Update(ctx context.Context, req *vo.Update{{.Table.TableName}}Request) (*vo.Update{{.Table.TableName}}Response, error) {
	//var obj model.{{.Table.TableName}}
    obj,err:=s.{{.Table.Uri}}Repository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
    err = copier.Copy(&obj, req)
	var resp vo.Update{{.Table.TableName}}Response
	_,err= s.{{.Table.Uri}}Repository.Update(ctx, obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return &resp, err
}


func (s *{{.Table.TableName}}Service) Delete(ctx context.Context, req *vo.Delete{{.Table.TableName}}Request) (int64, error) {
	return s.{{.Table.Uri}}Repository.Delete(ctx, req.Ids)
}


{{end}}

