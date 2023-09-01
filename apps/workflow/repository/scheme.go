package repository

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/global"

	"time"
)

type SchemeRepo interface {
	Get(ctx context.Context, schemeCode string) (*model.SchemeTab, error)
	Insert(ctx context.Context, schemaCode string, schemaContent string) error
	Update(ctx context.Context, schemeCode string, schemaContent string) error
}

type SchemeDBRepo struct{}

func (s *SchemeDBRepo) Get(ctx context.Context, schemeCode string) (*model.SchemeTab, error) {
	m := model.SchemeTab{}
	result := global.Context(ctx).Where(&model.SchemeTab{SchemeCode: schemeCode}).First(&m)
	return &m, result.Error
}

func (s *SchemeDBRepo) Insert(ctx context.Context, schemaCode string, schemaContent string) error {
	m := model.SchemeTab{
		Scheme:     schemaContent,
		SchemeCode: schemaCode,
		IsObsolete: 0,
		Ctime:      int(time.Now().Unix()),
		Mtime:      int(time.Now().Unix()),
	}
	err := global.Context(ctx).Create(&m).Error
	return err
}

func (s *SchemeDBRepo) Update(ctx context.Context, schemeCode string, schemaContent string) error {
	m := model.SchemeTab{
		Scheme:     schemaContent,
		SchemeCode: schemeCode,
		IsObsolete: 0,
		Mtime:      int(time.Now().Unix()),
	}
	err := global.Context(ctx).Model(&m).Where("scheme_code = ?", schemeCode).Updates(&m).Error
	return err
}
