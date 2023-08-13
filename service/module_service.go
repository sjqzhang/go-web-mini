package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"go-web-mini/model"
	"go-web-mini/repository"
	"go-web-mini/vo"
)

/*
// ModuleQueryPage 模块配置表分页查询
func ModuleQueryPage(param dto.ModulePageDTO) []vo.ModuleVO{

	return []vo.ModuleVO{}
}
*/

type ModuleService struct {
	moduleRepository repository.ModuleRepository
}

func (s *ModuleService) List(ctx context.Context, req *vo.ListModuleRequest) (*vo.ListModuleResponse, error) {
	var query model.ModuleQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListModuleResponse
	objs, err := s.moduleRepository.List(ctx, &query)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, objs)
	return &resp, err
}

func (s *ModuleService) GetById(ctx context.Context, req *vo.GetModuleRequest) (*vo.ModuleResponse, error) {
	obj, err := s.moduleRepository.GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp vo.ModuleResponse
	err = copier.Copy(&resp, obj)
	return &resp, err
}

func (s *ModuleService) Create(ctx *gin.Context, req *vo.CreateModuleRequest) (*vo.CreateModuleResponse, error) {
	var obj model.Module
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateModuleResponse
	_, err = s.moduleRepository.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *ModuleService) Update(ctx *gin.Context, req *vo.UpdateModuleRequest) (*vo.UpdateModuleResponse, error) {
	//var obj model.Module
	obj, err := s.moduleRepository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&obj, req)
	var resp vo.UpdateModuleResponse
	_, err = s.moduleRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *ModuleService) Delete(ctx *gin.Context, req *vo.DeleteModuleRequest) (int64, error) {
	return s.moduleRepository.Delete(ctx, req.Ids)
}
