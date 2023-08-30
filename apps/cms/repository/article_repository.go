package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-web-mini/apps/cms/model"
	"go-web-mini/global"
	"time"
)

type ArticleRepository struct {
}

func (r *ArticleRepository) List(ctx context.Context, query *model.ArticleQuery) (*model.PagerModel, error) {
	db := global.GetDB(ctx)
	var list []*model.Article
	var obj model.Article
	copier.CopyWithOption(&obj, &query, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	var total int64
	where, values, _ := model.BuildWhere(obj)
	err := db.Debug().Model(&obj).Where(where, values...).Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&obj).Debug().Where(where, values...).Where("deleted_at is null").Offset((query.PageNum - 1) * query.PageSize).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var pagerModel model.PagerModel
	pagerModel.List = list
	pagerModel.Total = total
	pagerModel.PageNum = query.PageNum
	pagerModel.PageSize = query.PageSize
	return &pagerModel, err
}

func (r *ArticleRepository) Create(ctx context.Context, obj *model.Article) (*model.Article, error) {
	db := global.GetDB(ctx)
	return obj, db.Create(obj).Error
}

func (r *ArticleRepository) GetById(ctx context.Context, id int64) (*model.Article, error) {
	db := global.GetDB(ctx)
	var obj model.Article
	err := db.Model(obj).Where("id=?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func (r *ArticleRepository) Update(ctx context.Context, obj *model.Article) (*model.Article, error) {
	db := global.GetDB(ctx)
	if obj.ID == 0 {
		return nil, fmt.Errorf("id is empty")
	}
	count := db.Model(obj).Where("id=?", obj.ID).Save(obj).RowsAffected
	if count > 0 {
		return obj, db.First(obj).Error
	}
	return nil, fmt.Errorf("not found")
}

func (r *ArticleRepository) Delete(ctx context.Context, ids []int64) (int64, error) {
	db := global.GetDB(ctx)
	//软删除
	return db.Model(model.Article{}).Where("id in (?)", ids).UpdateColumn("deleted_at", time.Now()).RowsAffected, nil
}
