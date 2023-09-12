package service

import (
	"context"
	"github.com/jinzhu/copier"

	"go-web-mini/apps/system/model"
	"go-web-mini/apps/system/repository"
	"go-web-mini/apps/system/vo"
)

/*
// DictionaryTypeQueryPage dictionary_type分页查询
func DictionaryTypeQueryPage(param dto.DictionaryTypePageDTO) []vo.DictionaryTypeVO{

	return []vo.DictionaryTypeVO{}
}
*/

type IDictionaryTypeService interface {
	ListForPager(ctx context.Context, req *vo.ListForPagerDictionaryTypeRequest) (*vo.ListForPagerDictionaryTypeResponse, error)
	List(ctx context.Context, req *vo.ListDictionaryTypeRequest) (*vo.ListDictionaryTypeResponse, error)
	GetById(ctx context.Context, req *vo.GetDictionaryTypeRequest) (*vo.DictionaryTypeResponse, error)
	Create(ctx context.Context, req *vo.CreateDictionaryTypeRequest) (*vo.CreateDictionaryTypeResponse, error)
	Update(ctx context.Context, req *vo.UpdateDictionaryTypeRequest) (*vo.UpdateDictionaryTypeResponse, error)
	Delete(ctx context.Context, req *vo.DeleteDictionaryTypeRequest) (int64, error)
}

type DictionaryTypeService struct {
	dictionaryTypeRepository repository.DictionaryTypeRepository
}

func (s *DictionaryTypeService) ListForPager(ctx context.Context, req *vo.ListForPagerDictionaryTypeRequest) (*vo.ListForPagerDictionaryTypeResponse, error) {
	var query model.DictionaryTypeQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListForPagerDictionaryTypeResponse
	objs, err := s.dictionaryTypeRepository.ListForPager(ctx, &query)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, objs)
	return &resp, err
}

func (s *DictionaryTypeService) List(ctx context.Context, req *vo.ListDictionaryTypeRequest) (*vo.ListDictionaryTypeResponse, error) {
	var query model.DictionaryTypeQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListDictionaryTypeResponse
	objs, err := s.dictionaryTypeRepository.List(ctx, &query)
	if err != nil {
		return nil, err
	}
	if err = copier.Copy(&resp.List, objs); err != nil {
		return nil, err
	}
	return &resp, err
}

func (s *DictionaryTypeService) GetById(ctx context.Context, req *vo.GetDictionaryTypeRequest) (*vo.DictionaryTypeResponse, error) {
	obj, err := s.dictionaryTypeRepository.GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp vo.DictionaryTypeResponse
	err = copier.Copy(&resp, obj)
	return &resp, err
}

func (s *DictionaryTypeService) Create(ctx context.Context, req *vo.CreateDictionaryTypeRequest) (*vo.CreateDictionaryTypeResponse, error) {
	var obj model.DictionaryType
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateDictionaryTypeResponse
	_, err = s.dictionaryTypeRepository.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *DictionaryTypeService) Update(ctx context.Context, req *vo.UpdateDictionaryTypeRequest) (*vo.UpdateDictionaryTypeResponse, error) {
	//var obj model.DictionaryType
	obj, err := s.dictionaryTypeRepository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&obj, req)
	var resp vo.UpdateDictionaryTypeResponse
	_, err = s.dictionaryTypeRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *DictionaryTypeService) Delete(ctx context.Context, req *vo.DeleteDictionaryTypeRequest) (int64, error) {
	return s.dictionaryTypeRepository.Delete(ctx, req.Ids)
}
