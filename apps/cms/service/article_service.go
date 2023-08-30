package service

import (
	"context"
	"github.com/jinzhu/copier"

	"go-web-mini/apps/cms/model"
	"go-web-mini/apps/cms/repository"
	"go-web-mini/apps/cms/vo"
)

/*
// ArticleQueryPage article分页查询
func ArticleQueryPage(param dto.ArticlePageDTO) []vo.ArticleVO{

	return []vo.ArticleVO{}
}
*/

type ArticleService struct {
	articleRepository repository.ArticleRepository
}

func (s *ArticleService) List(ctx context.Context, req *vo.ListArticleRequest) (*vo.ListArticleResponse, error) {
	var query model.ArticleQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListArticleResponse
	objs, err := s.articleRepository.List(ctx, &query)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, objs)
	return &resp, err
}

func (s *ArticleService) GetById(ctx context.Context, req *vo.GetArticleRequest) (*vo.ArticleResponse, error) {
	obj, err := s.articleRepository.GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp vo.ArticleResponse
	err = copier.Copy(&resp, obj)
	return &resp, err
}

func (s *ArticleService) Create(ctx context.Context, req *vo.CreateArticleRequest) (*vo.CreateArticleResponse, error) {
	var obj model.Article
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateArticleResponse
	_, err = s.articleRepository.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *ArticleService) Update(ctx context.Context, req *vo.UpdateArticleRequest) (*vo.UpdateArticleResponse, error) {
	//var obj model.Article
	obj, err := s.articleRepository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&obj, req)
	var resp vo.UpdateArticleResponse
	_, err = s.articleRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *ArticleService) Delete(ctx context.Context, req *vo.DeleteArticleRequest) (int64, error) {
	return s.articleRepository.Delete(ctx, req.Ids)
}
