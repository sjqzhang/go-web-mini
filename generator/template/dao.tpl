{{define "repository"}}package repository

import (
	"context"
	"fmt"
	"{{.ModuleName}}/common"
	"{{.ModuleName}}/model"
	"time"
)




type {{.Table.TableName}}Repository struct {
}

func (r *{{.Table.TableName}}Repository) List(ctx context.Context, query *model.{{.Table.TableName}}Query) ([]*model.{{.Table.TableName}}, error) {
	db := common.GetDB(ctx)
	var list []*model.{{.Table.TableName}}
	err := db.Debug().Offset(query.PageNum * query.PageSize).Limit(query.PageSize).Find(&list).Error
	return list, err
}

func (r *{{.Table.TableName}}Repository) Create(ctx context.Context, obj *model.{{.Table.TableName}}) (*model.{{.Table.TableName}}, error) {
	db := common.GetDB(ctx)
	return obj, db.Create(obj).Error
}

func (r *{{.Table.TableName}}Repository) Update(ctx context.Context, obj *model.{{.Table.TableName}}) (*model.{{.Table.TableName}}, error) {
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

func (r *{{.Table.TableName}}Repository) Delete(ctx context.Context, obj *model.{{.Table.TableName}}) (int64, error) {
	db := common.GetDB(ctx)
	//软删除
	return db.Model(obj).UpdateColumn("deleted_at", time.Now()).Where("id = ?", obj.ID).RowsAffected, nil
}


{{end}}
