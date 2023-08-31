package service

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/apps/workflow/repository"

	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var lockCleanerOnce sync.Once

type LockerService struct {
	LockerRepo repository.LockerRepo
}

func NewLockerService() *LockerService {
	return &LockerService{LockerRepo: repository.LockerDBRepo{}}
}

func (l *LockerService) Lock(ctx context.Context, name string, expire time.Duration) (locker *model.LockerTab, err error) {
	locker = new(model.LockerTab)
	locker.LockerName = name
	locker.ExpireTime = int(time.Now().Add(expire).Unix())
	_, err = l.LockerRepo.New(ctx, locker)
	return locker, err
}

func (l *LockerService) Unlock(ctx context.Context, id int) (err error) {
	return l.LockerRepo.Delete(ctx, id)
}

func (l *LockerService) StartLockCleaner(ctx context.Context, duration time.Duration) {
	lockCleanerOnce.Do(
		func() {
			ticker := time.NewTicker(duration)
			go func() {
				defer func() {
					ticker.Stop()
				}()
				for range ticker.C {
					now := int(time.Now().Unix())
					lockers, err := l.LockerRepo.Find(ctx, now)
					if err != nil {
						logrus.Errorf("LockCleaner Find err: %+v", err)
						continue
					}
					for _, locker := range lockers {
						err = l.Unlock(ctx, locker.Id)
						if err != nil {
							logrus.WithFields(logrus.Fields{
								"id":          locker.Id,
								"locker_name": locker.LockerName,
							}).Errorf("LockCleaner unlock err: %+v", err)
						}
					}
				}
			}()
		},
	)
}
