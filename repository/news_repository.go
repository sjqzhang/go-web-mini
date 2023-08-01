package repository

import (
	"context"
	"go-web-mini/common"
	"go-web-mini/model"
)

type ITestRepository interface {
	CreateNews(ctx context.Context, user *model.News) error // 创建用户
	// 清理所有用户信息缓存
}

type TestRepository struct {
}

func (r *TestRepository) CreateNews(ctx context.Context, user *model.News) error {
	db := common.GetDB(ctx)
	return db.Create(user).Error
}
