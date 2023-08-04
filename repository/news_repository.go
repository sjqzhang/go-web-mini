package repository

import (
	"context"
	"fmt"
	"go-web-mini/common"
	"go-web-mini/model"
	"time"
)




type NewsRepository struct {
}

func (r *NewsRepository) List(ctx context.Context, query *model.NewsQuery) ([]*model.News, error) {
	db := common.GetDB(ctx)
	var list []*model.News
	err := db.Debug().Offset(query.PageNum * query.PageSize).Limit(query.PageSize).Find(&list).Error
	return list, err
}

func (r *NewsRepository) Create(ctx context.Context, obj *model.News) (*model.News, error) {
	db := common.GetDB(ctx)
	return obj, db.Create(obj).Error
}

func (r *NewsRepository) Update(ctx context.Context, obj *model.News) (*model.News, error) {
	db := common.GetDB(ctx)
	if obj.ID==0  {
		return nil, fmt.Errorf("id is empty")
	}
	count := db.Model(obj).Updates(obj).RowsAffected
	if count > 0 {
		return obj, db.First(obj).Error
	}
	return nil, fmt.Errorf("not found")
}

func (r *NewsRepository) Delete(ctx context.Context, obj *model.News) (int64, error) {
	db := common.GetDB(ctx)
	//软删除
	return db.Model(obj).UpdateColumn("deleted_at", time.Now()).Where("id = ?", obj.ID).RowsAffected, nil
}


