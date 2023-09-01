package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/global"

	"time"
)

type InstanceRepo interface {
	Get(ctx context.Context, instanceId int) (*model.InstanceTab, error)
	GetList(ctx context.Context, instanceIds []int) ([]*model.InstanceTab, error)
	New(ctx context.Context, instance *model.InstanceTab) error
	Update(ctx context.Context, instance *model.InstanceTab) error
	UpdateSchemaCode(ctx context.Context, instanceIds []int, schemaCode string) error
}

type InstanceDBRepo struct{}

func (i *InstanceDBRepo) New(ctx context.Context, instance *model.InstanceTab) error {
	now := int(time.Now().Unix())
	instance.Ctime = now
	instance.Mtime = now
	return global.Context(ctx).Create(instance).Error
}

func (i *InstanceDBRepo) Get(ctx context.Context, instanceId int) (*model.InstanceTab, error) {
	m := model.InstanceTab{}
	result := global.Context(ctx).Where(&model.InstanceTab{Id: instanceId}).First(&m)
	return &m, result.Error
}

func (i *InstanceDBRepo) GetList(ctx context.Context, instanceIds []int) ([]*model.InstanceTab, error) {
	var ret []*model.InstanceTab
	result := global.Context(ctx).Where("id IN (?)", instanceIds).Find(&ret)
	return ret, result.Error
}

func (i *InstanceDBRepo) Update(ctx context.Context, instance *model.InstanceTab) error {
	instance.Mtime = int(time.Now().Unix())
	db := global.Context(ctx).Save(instance)
	return db.Error
}

func (i *InstanceDBRepo) UpdateSchemaCode(ctx context.Context, instanceIds []int, schemaCode string) error {
	tab := model.InstanceTab{
		SchemeCode: schemaCode,
	}
	db := global.Context(ctx).Where("id IN (?)", instanceIds).Updates(&tab)
	return db.Error
}
