package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/global"

	"time"
)

type InstanceRepo interface {
	Get(ctx context.Context, instanceId int) (*model.InstanceTab, error)
	New(ctx context.Context, instance *model.InstanceTab) error
	Update(ctx context.Context, instance *model.InstanceTab) error
}

type InstanceDBRepo struct{}

func (i *InstanceDBRepo) New(ctx context.Context, instance *model.InstanceTab) error {
	now := int(time.Now().Unix())
	instance.Ctime = now
	instance.Mtime = now
	return global.GetDB(ctx).Create(instance).Error
}

func (i *InstanceDBRepo) Get(ctx context.Context, instanceId int) (*model.InstanceTab, error) {
	m := model.InstanceTab{}
	result := global.GetDB(ctx).Where(&model.InstanceTab{Id: instanceId}).First(&m)
	return &m, result.Error
}

func (i *InstanceDBRepo) Update(ctx context.Context, instance *model.InstanceTab) error {
	instance.Mtime = int(time.Now().Unix())
	db := global.GetDB(ctx).Save(instance)
	return db.Error
}
