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
// NewsQueryPage news分页查询
func NewsQueryPage(param dto.NewsPageDTO) []vo.NewsVO{

	return []vo.NewsVO{}
}
*/

type NewsService struct {
	newsRepository repository.NewsRepository
}

func (s *NewsService) List(ctx context.Context, req *vo.ListNewsRequest) (*vo.ListNewsResponse, error) {
	var query model.NewsQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListNewsResponse
	objs, err := s.newsRepository.List(ctx, &query)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, objs)
	return &resp, err
}

func (s *NewsService) GetById(ctx context.Context, req *vo.GetNewsRequest) (*vo.NewsResponse, error) {
	obj, err := s.newsRepository.GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp vo.NewsResponse
	err = copier.Copy(&resp, obj)
	return &resp, err
}

func (s *NewsService) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (*vo.CreateNewsResponse, error) {
	var obj model.News
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateNewsResponse
	_, err = s.newsRepository.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *NewsService) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (*vo.UpdateNewsResponse, error) {
	//var obj model.News
	obj, err := s.newsRepository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&obj, req)
	var resp vo.UpdateNewsResponse
	_, err = s.newsRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *NewsService) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (int64, error) {
	return s.newsRepository.Delete(ctx, req.Ids)
}
