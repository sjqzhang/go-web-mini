package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-web-mini/apps/system/model"
	"go-web-mini/global"
	"time"
)

type IDictionaryRepository interface {
	ListForPager(ctx context.Context, query *model.DictionaryQuery) (*model.PagerModel, error)
	List(ctx context.Context, query *model.DictionaryQuery) ([]*model.Dictionary, error)
	Create(ctx context.Context, obj *model.Dictionary) (*model.Dictionary, error)
	GetById(ctx context.Context, id int64) (*model.Dictionary, error)
	Update(ctx context.Context, obj *model.Dictionary) (*model.Dictionary, error)
	Delete(ctx context.Context, ids []int64) (int64, error)
}

type DictionaryRepository struct {
}

func (r *DictionaryRepository) ListForPager(ctx context.Context, query *model.DictionaryQuery) (*model.PagerModel, error) {
	db := global.GetDB(ctx)
	var list []*model.Dictionary
	var obj model.Dictionary
	err := copier.CopyWithOption(&obj, &query, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}
	var total int64
	where, values, _ := model.BuildWhere(obj)
	err = db.Model(&obj).Where(where, values...).Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, err
	}
	if query.PageSize > 1000 {
		query.PageSize = 1000
	}
	err = db.Preload("DictType").Where(where, values...).Where("deleted_at is null").Offset((query.PageNum - 1) * query.PageSize).Limit(query.PageSize).Find(&list).Error
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

func (r *DictionaryRepository) List(ctx context.Context, query *model.DictionaryQuery) ([]*model.Dictionary, error) {
	db := global.GetDB(ctx)
	var list []*model.Dictionary
	var obj model.Dictionary
	err := copier.CopyWithOption(&obj, &query, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}
	where, values, _ := model.BuildWhere(obj)
	if query.PageSize > 1000 {
		query.PageSize = 1000
	} else if query.PageSize <= 0 {
		query.PageSize = 100
	}
	err = db.Model(&obj).Where(where, values...).Where("deleted_at is null").Offset((query.PageNum - 1) * query.PageSize).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *DictionaryRepository) Create(ctx context.Context, obj *model.Dictionary) (*model.Dictionary, error) {
	db := global.GetDB(ctx)
	return obj, db.Create(obj).Error
}

func (r *DictionaryRepository) GetById(ctx context.Context, id int64) (*model.Dictionary, error) {
	db := global.GetDB(ctx)
	var obj model.Dictionary
	err := db.Model(obj).Where("id=?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func (r *DictionaryRepository) Update(ctx context.Context, obj *model.Dictionary) (*model.Dictionary, error) {
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

func (r *DictionaryRepository) Delete(ctx context.Context, ids []int64) (int64, error) {
	db := global.GetDB(ctx)
	//软删除
	return db.Model(model.Dictionary{}).Where("id in (?)", ids).UpdateColumn("deleted_at", time.Now()).RowsAffected, nil
}
