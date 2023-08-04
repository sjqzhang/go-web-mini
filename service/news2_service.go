package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go-web-mini/model"
	"go-web-mini/repository"
	"go-web-mini/vo"
)

type News2Service struct {
	newsRepository repository.News2Repository
}

func (s *News2Service) List(ctx context.Context, news *vo.ListNewsRequest) ([]*model.News, error) {
	var query model.NewsQuery
	err := copier.Copy(&query, news)
	if err != nil {
		return nil, err
	}
	return s.newsRepository.List(ctx, &query)
}

func (s *News2Service) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (*model.News, error) {
	var news model.News
	err := copier.Copy(&news, req)
	if err != nil {
		return nil, err
	}
	return s.newsRepository.Create(ctx, &news)
}


func (s *News2Service) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (*model.News, error) {
	var news model.News
	err := copier.Copy(&news, req)
	if err != nil {
		return nil, err
	}
	return s.newsRepository.Update(ctx, &news)
}


func (s *News2Service) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (int64, error) {
	var news model.News
	err := copier.Copy(&news, req)
	if err != nil {
		return 0, err
	}
	return s.newsRepository.Delete(ctx, &news)
}
