package service

import (
	"context"
	"github.com/jinzhu/copier"

	"go-web-mini/apps/system/model"
	"go-web-mini/apps/system/repository"
	"go-web-mini/apps/system/vo"
)

/*
// DictionaryQueryPage dictionary分页查询
func DictionaryQueryPage(param dto.DictionaryPageDTO) []vo.DictionaryVO{

	return []vo.DictionaryVO{}
}
*/

type IDictionaryService interface {
	ListForPager(ctx context.Context, req *vo.ListForPagerDictionaryRequest) (*vo.ListForPagerDictionaryResponse, error)
	List(ctx context.Context, req *vo.ListDictionaryRequest) (*vo.ListDictionaryResponse, error)
	GetById(ctx context.Context, req *vo.GetDictionaryRequest) (*vo.DictionaryResponse, error)
	Create(ctx context.Context, req *vo.CreateDictionaryRequest) (*vo.CreateDictionaryResponse, error)
	Update(ctx context.Context, req *vo.UpdateDictionaryRequest) (*vo.UpdateDictionaryResponse, error)
	Delete(ctx context.Context, req *vo.DeleteDictionaryRequest) (int64, error)
}

type DictionaryService struct {
	dictionaryRepository repository.DictionaryRepository
}

func (s *DictionaryService) ListForPager(ctx context.Context, req *vo.ListForPagerDictionaryRequest) (*vo.ListForPagerDictionaryResponse, error) {
	var query model.DictionaryQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListForPagerDictionaryResponse
	objs, err := s.dictionaryRepository.ListForPager(ctx, &query)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, objs)
	return &resp, err
}

func (s *DictionaryService) List(ctx context.Context, req *vo.ListDictionaryRequest) (*vo.ListDictionaryResponse, error) {
	var query model.DictionaryQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListDictionaryResponse
	objs, err := s.dictionaryRepository.List(ctx, &query)
	if err != nil {
		return nil, err
	}
	if err = copier.Copy(&resp.List, objs); err != nil {
		return nil, err
	}
	return &resp, err
}

func (s *DictionaryService) GetById(ctx context.Context, req *vo.GetDictionaryRequest) (*vo.DictionaryResponse, error) {
	obj, err := s.dictionaryRepository.GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp vo.DictionaryResponse
	err = copier.Copy(&resp, obj)
	return &resp, err
}

func (s *DictionaryService) Create(ctx context.Context, req *vo.CreateDictionaryRequest) (*vo.CreateDictionaryResponse, error) {
	var obj model.Dictionary
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateDictionaryResponse
	_, err = s.dictionaryRepository.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *DictionaryService) Update(ctx context.Context, req *vo.UpdateDictionaryRequest) (*vo.UpdateDictionaryResponse, error) {
	//var obj model.Dictionary
	obj, err := s.dictionaryRepository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&obj, req)
	var resp vo.UpdateDictionaryResponse
	_, err = s.dictionaryRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *DictionaryService) Delete(ctx context.Context, req *vo.DeleteDictionaryRequest) (int64, error) {
	return s.dictionaryRepository.Delete(ctx, req.Ids)
}
