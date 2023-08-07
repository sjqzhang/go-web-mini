package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-web-mini/common"
	"go-web-mini/model"
	"time"
)




type BranchRepository struct {
}

func (r *BranchRepository) List(ctx context.Context, query *model.BranchQuery) (*model.PagerModel, error) {
	db := common.GetDB(ctx)
	var list []*model.Branch
	var obj model.Branch
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

func (r *BranchRepository) Create(ctx context.Context, obj *model.Branch) (*model.Branch, error) {
	db := common.GetDB(ctx)
	return obj, db.Create(obj).Error
}


func (r *BranchRepository) GetById(ctx context.Context,  id int64) (*model.Branch, error) {
	db := common.GetDB(ctx)
	var obj model.Branch
    err:=db.Model(obj).Where("id=?",id).First(&obj).Error
	if err != nil {
	    return nil,err
	}
	return &obj,err
}

func (r *BranchRepository) Update(ctx context.Context, obj *model.Branch) (*model.Branch, error) {
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

func (r *BranchRepository) Delete(ctx context.Context, ids []int64) (int64, error) {
	db := common.GetDB(ctx)
	//软删除
	return db.Model(model.Branch{}).Where("id in (?)", ids).UpdateColumn("deleted_at", time.Now()).RowsAffected, nil
}


