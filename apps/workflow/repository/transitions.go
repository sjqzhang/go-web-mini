package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/global"

	"time"
)

type TransitionsRepo interface {
	QueryByInstanceId(ctx context.Context, instanceId int) ([]model.TransitionsTab, error)
	QueryByIds(ctx context.Context, ids []int) ([]model.TransitionsTab, error)
	New(ctx context.Context, transition *model.TransitionsTab) error
}

type TransitionsDBRepo struct{}

func (t TransitionsDBRepo) QueryByInstanceId(ctx context.Context, instanceId int) ([]model.TransitionsTab, error) {
	// TODO: offset and limit support
	var transitions []model.TransitionsTab
	db := global.Context(ctx).Where("instance_id = ?", instanceId).Find(&transitions)
	return transitions, db.Error
}

func (t TransitionsDBRepo) QueryByIds(ctx context.Context, ids []int) ([]model.TransitionsTab, error) {
	var transitions []model.TransitionsTab
	db := global.Context(ctx).Where("id IN (?)", ids).Find(&transitions)
	return transitions, db.Error
}

func (t TransitionsDBRepo) New(ctx context.Context, transition *model.TransitionsTab) error {
	now := int(time.Now().Unix())
	transition.Ctime = now
	return global.Context(ctx).Create(transition).Error
}
