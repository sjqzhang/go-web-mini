package service

import (
	"context"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/apps/workflow/repository"

	"sync"
)

type SchemeService struct {
	SchemeRepo repository.SchemeRepo
}

func NewSchemeService() *SchemeService {
	return &SchemeService{SchemeRepo: &repository.SchemeDBRepo{}}
}

func (s SchemeService) Get(ctx context.Context, code string) (*model.SchemeTab, error) {
	cache := getSchemeCache()
	m := cache.Get(ctx, code)
	if m != nil {
		return m, nil
	}

	m, err := s.SchemeRepo.Get(ctx, code)
	if err != nil {
		return nil, err
	}

	cache.Set(ctx, m)
	return m, nil
}

type schemeCache struct {
	lock    sync.RWMutex
	schemes map[string]*model.SchemeTab
}

var cache *schemeCache
var cacheOnce sync.Once

func getSchemeCache() *schemeCache {
	cacheOnce.Do(func() {
		cache = &schemeCache{
			schemes: make(map[string]*model.SchemeTab),
		}
	})
	return cache
}

func (s *schemeCache) Get(ctx context.Context, code string) *model.SchemeTab {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if m, exist := s.schemes[code]; exist {
		return m
	}
	return nil
}

func (s *schemeCache) Set(ctx context.Context, m *model.SchemeTab) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.schemes[m.SchemeCode] = m
}
