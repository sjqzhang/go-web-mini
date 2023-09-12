package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-web-mini/apps/system/model"
	"go-web-mini/global"
	"time"
)

type IDictionaryTypeRepository interface {
	ListForPager(ctx context.Context, query *model.DictionaryTypeQuery) (*model.PagerModel, error)
	List(ctx context.Context, query *model.DictionaryTypeQuery) ([]*model.DictionaryType, error)
	Create(ctx context.Context, obj *model.DictionaryType) (*model.DictionaryType, error)
	GetById(ctx context.Context, id int64) (*model.DictionaryType, error)
	Update(ctx context.Context, obj *model.DictionaryType) (*model.DictionaryType, error)
	Delete(ctx context.Context, ids []int64) (int64, error)
}

type DictionaryTypeRepository struct {
}

func (r *DictionaryTypeRepository) ListForPager(ctx context.Context, query *model.DictionaryTypeQuery) (*model.PagerModel, error) {
	db := global.GetDB(ctx)
	var list []*model.DictionaryType
	var obj model.DictionaryType
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
	if query.PageSize <= 0 {
		query.PageSize = 100
	}
	err = db.Preload("Dictionaries").Where(where, values...).Where("deleted_at is null").Offset((query.PageNum - 1) * query.PageSize).Limit(query.PageSize).Find(&list).Error
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

func (r *DictionaryTypeRepository) List(ctx context.Context, query *model.DictionaryTypeQuery) ([]*model.DictionaryType, error) {
	db := global.GetDB(ctx)
	var list []*model.DictionaryType
	var obj model.DictionaryType
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

func (r *DictionaryTypeRepository) Create(ctx context.Context, obj *model.DictionaryType) (*model.DictionaryType, error) {
	db := global.GetDB(ctx)
	return obj, db.Create(obj).Error
}

func (r *DictionaryTypeRepository) GetById(ctx context.Context, id int64) (*model.DictionaryType, error) {
	db := global.GetDB(ctx)
	var obj model.DictionaryType
	err := db.Model(obj).Where("id=?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func (r *DictionaryTypeRepository) Update(ctx context.Context, obj *model.DictionaryType) (*model.DictionaryType, error) {
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

func (r *DictionaryTypeRepository) Delete(ctx context.Context, ids []int64) (int64, error) {
	db := global.GetDB(ctx)
	//软删除
	return db.Model(model.DictionaryType{}).Where("id in (?)", ids).UpdateColumn("deleted_at", time.Now()).RowsAffected, nil
}
