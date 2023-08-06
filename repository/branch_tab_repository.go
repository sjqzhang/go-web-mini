package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-web-mini/common"
	"go-web-mini/model"
	"time"
)




type BranchTabRepository struct {
}

func (r *BranchTabRepository) List(ctx context.Context, query *model.BranchTabQuery) (*model.PagerModel, error) {
	db := common.GetDB(ctx)
	var list []*model.BranchTab
	var obj model.BranchTab
	copier.CopyWithOption(&obj, &query, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	var total int64
	where,values,_:=model.BuildWhere(obj)
	err := db.Debug().Model(&obj).Where(where,values...).Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&obj).Debug().Where(where,values...).Where("deleted_at is null").Offset((query.PageNum-1) * query.PageSize).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var pagerModel model.PagerModel
	pagerModel.List=list
	pagerModel.Total=total
	pagerModel.PageNum=query.PageNum
	pagerModel.PageSize=query.PageSize
	return &pagerModel, err
}

func (r *BranchTabRepository) Create(ctx context.Context, obj *model.BranchTab) (*model.BranchTab, error) {
	db := common.GetDB(ctx)
	return obj, db.Create(obj).Error
}


func (r *BranchTabRepository) GetById(ctx context.Context,  id int64) (*model.BranchTab, error) {
	db := common.GetDB(ctx)
	var obj model.BranchTab
    err:=db.Model(obj).Where("id=?",id).First(&obj).Error
	if err != nil {
	    return nil,err
	}
	return &obj,err
}

func (r *BranchTabRepository) Update(ctx context.Context, obj *model.BranchTab) (*model.BranchTab, error) {
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

func (r *BranchTabRepository) Delete(ctx context.Context, ids []int64) (int64, error) {
	db := common.GetDB(ctx)
	//软删除
	return db.Model(model.BranchTab{}).Where("id in (?)", ids).UpdateColumn("deleted_at", time.Now()).RowsAffected, nil
}


