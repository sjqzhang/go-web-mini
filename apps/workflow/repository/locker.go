package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/global"

	"time"
)

type LockerRepo interface {
	New(ctx context.Context, locker *model.LockerTab) (rowsAffected int64, e error)
	Delete(ctx context.Context, id int) error
	Find(ctx context.Context, lteExpireTime int) (lockers []model.LockerTab, err error)
}

type LockerDBRepo struct {}

func (l LockerDBRepo) New(ctx context.Context, locker *model.LockerTab) (rowsAffected int64, e error) {
	now := int(time.Now().Unix())
	locker.Ctime = now
	db := global.GetDB(ctx).Create(locker)
	return db.RowsAffected, db.Error
}

// The deleted value needs to have primary key or it will trigger a Batch Delete
func (l LockerDBRepo) Delete(ctx context.Context, id int) error {
	db := global.GetDB(ctx).Delete(&model.LockerTab{Id: id})
	return db.Error
}

func (l LockerDBRepo) Find(ctx context.Context, lteExpireTime int) (lockers []model.LockerTab, err error) {
	db := global.GetDB(ctx).Where("expire_time <= ?", lteExpireTime).Find(&lockers)
	return lockers, db.Error
}
