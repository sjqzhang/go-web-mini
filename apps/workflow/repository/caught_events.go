package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/global"

	"time"
)

type CaughtEventsRepo interface {
	QueryByInstanceId(ctx context.Context, instanceId int) ([]model.CaughtEventsTab, error)
	New(ctx context.Context, event *model.CaughtEventsTab) error
	Update(ctx context.Context, event *model.CaughtEventsTab) error
}

type CaughtEventsDBRepo struct{}

func (c CaughtEventsDBRepo) QueryByInstanceId(ctx context.Context, instanceId int) ([]model.CaughtEventsTab, error) {
	var events []model.CaughtEventsTab

	db := global.GetDB(ctx).Where("instance_id = ?", instanceId).Find(&events)
	return events, db.Error
}

func (c CaughtEventsDBRepo) New(ctx context.Context, event *model.CaughtEventsTab) error {
	now := int(time.Now().Unix())
	event.Ctime = now
	event.Mtime = now
	db :=global.GetDB(ctx).Create(event)
	return db.Error
}

func (c CaughtEventsDBRepo) Update(ctx context.Context, event *model.CaughtEventsTab) error {
	event.Mtime = int(time.Now().Unix())
	db := global.GetDB(ctx).Save(event)
	return db.Error
}
