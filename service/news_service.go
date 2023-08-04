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
// NewsQueryPage 分页查询
func NewsQueryPage(param dto.NewsPageDTO) []vo.NewsVO{

	return []vo.NewsVO{}
}
*/

type NewsService struct {
	NewsRepository repository.NewsRepository
}

func (s *NewsService) List(ctx context.Context, req *vo.ListNewsRequest) ([]*model.News, error) {
	var query model.NewsQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	return s.NewsRepository.List(ctx, &query)
}

func (s *NewsService) Create(ctx *gin.Context, req *vo.CreateNewsRequest) (*model.News, error) {
	var obj model.News
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	return s.NewsRepository.Create(ctx, &obj)
}


func (s *NewsService) Update(ctx *gin.Context, req *vo.UpdateNewsRequest) (*model.News, error) {
	var obj model.News
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	return s.NewsRepository.Update(ctx, &obj)
}


func (s *NewsService) Delete(ctx *gin.Context, req *vo.DeleteNewsRequest) (int64, error) {
	var obj model.News
	err := copier.Copy(&obj, req)
	if err != nil {
		return 0, err
	}
	return s.NewsRepository.Delete(ctx, &obj)
}


