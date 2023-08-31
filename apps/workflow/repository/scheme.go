package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"

	"go-web-mini/global"
)

type SchemeRepo interface {
	Get(ctx context.Context, schemeCode string) (*model.SchemeTab, error)
}

type SchemeDBRepo struct{}

func (s *SchemeDBRepo) Get(ctx context.Context, schemeCode string) (*model.SchemeTab, error) {
	m := model.SchemeTab{}
	result := global.GetDB(ctx).Where(&model.SchemeTab{SchemeCode: schemeCode}).First(&m)
	return &m, result.Error
}
