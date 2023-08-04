package repository

import (
	"context"
	"fmt"
	"go-web-mini/common"
	"go-web-mini/model"
	"time"
)


type News2Repository struct {
}

func (r *News2Repository) List(ctx context.Context, query *model.NewsQuery) ([]*model.News, error) {
	db := common.GetDB(ctx)
	var list []*model.News
	err := db.Debug().Offset(query.PageNum * query.PageSize).Limit(query.PageSize).Find(&list).Error
	return list, err
}

func (r *News2Repository) Create(ctx context.Context, news *model.News) (*model.News, error) {
	db := common.GetDB(ctx)
	return news, db.Create(news).Error
}

func (r *News2Repository) Update(ctx context.Context, news *model.News) (*model.News, error) {
	db := common.GetDB(ctx)
	if news.ID==0  {
		return nil, fmt.Errorf("id is empty")
	}
	count := db.Model(news).Updates(news).RowsAffected
	if count > 0 {
		return news, db.First(news).Error
	}
	return nil, fmt.Errorf("not found")
}

func (r *News2Repository) Delete(ctx context.Context, news *model.News) (int64, error) {
	db := common.GetDB(ctx)
	//软删除
	return db.Model(news).UpdateColumn("deleted_at", time.Now()).Where("id = ?", news.ID).RowsAffected, nil
}
