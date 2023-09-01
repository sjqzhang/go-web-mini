package service

import (
	"context"
	"go-web-mini/apps/workflow/bpmn_engine"
	"go-web-mini/apps/workflow/errcode"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/apps/workflow/repository"
	"go-web-mini/apps/workflow/spec/BPMN20"

	"github.com/sirupsen/logrus"
	"sync"
)

type SchemeService struct {
	SchemeRepo repository.SchemeRepo
}

func NewSchemeService() *SchemeService {
	return &SchemeService{SchemeRepo: &repository.SchemeDBRepo{}}
}

func (s *SchemeService) Get(ctx context.Context, code string) (*model.SchemeTab, *bpmn_engine.ProcessInfo, error) {
	cache := GetSchemeCache()
	m, p := cache.Get(ctx, code)
	if m != nil && p != nil {
		return m, p, nil
	}

	m, err := s.SchemeRepo.Get(ctx, code)
	if err != nil {
		return nil, nil, err
	}
	engine := bpmn_engine.New(code + "-engine")
	p, err = engine.LoadFromBytes([]byte(m.Scheme))
	if err != nil {
		return nil, nil, err
	}

	cache.Set(ctx, m, p)
	return m, p, nil
}

func (s *SchemeService) GetSchemaTDefinitions(ctx context.Context, schemaCodes []string) ([]BPMN20.TDefinitions, errcode.Exception) {
	var tDefinitions []BPMN20.TDefinitions
	if len(schemaCodes) == 0 {
		return tDefinitions, nil
	}
	schemeService := NewSchemeService()
	for _, schemaCode := range schemaCodes {
		_, processInfo, err := schemeService.Get(ctx, schemaCode)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"scheme_code": schemaCode,
			}).Errorf("get scheme from DB err: %+v", err)
			return nil, errcode.Wrap(errcode.DBError, err)
		}
		processDefinition := processInfo.Definitions()
		tDefinitions = append(tDefinitions, processDefinition)
	}
	return tDefinitions, nil
}

func (s *SchemeService) UploadSchemaFile(ctx context.Context, schemaCode string, schemaContent string) errcode.Exception {
	e := s.SchemeRepo.Insert(ctx, schemaCode, schemaContent)
	if e != nil {
		return errcode.Wrap(errcode.DBError, e)
	}
	return nil
}

func (s *SchemeService) ResetSchemaCode(ctx context.Context, schemaCode string, schemaContent string) errcode.Exception {
	e := s.SchemeRepo.Update(ctx, schemaCode, schemaContent)
	if e != nil {
		return errcode.Wrap(errcode.DBError, e)
	}
	return nil
}

type schemeCache struct {
	lock                 sync.RWMutex
	schemes              map[string]*model.SchemeTab
	schemaMapProcessInfo map[string]*bpmn_engine.ProcessInfo
}

var cache *schemeCache
var cacheOnce sync.Once

func GetSchemeCache() *schemeCache {
	cacheOnce.Do(func() {
		cache = &schemeCache{
			schemes:              make(map[string]*model.SchemeTab),
			schemaMapProcessInfo: make(map[string]*bpmn_engine.ProcessInfo),
		}
	})
	return cache
}

func (s *schemeCache) Get(ctx context.Context, code string) (*model.SchemeTab, *bpmn_engine.ProcessInfo) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if m, exist := s.schemes[code]; exist {
		if processInfo, ok := s.schemaMapProcessInfo[code]; ok {
			return m, processInfo
		}
	}
	return nil, nil
}

func (s *schemeCache) Set(ctx context.Context, m *model.SchemeTab, processInfo *bpmn_engine.ProcessInfo) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.schemes[m.SchemeCode] = m
	s.schemaMapProcessInfo[m.SchemeCode] = processInfo
}

func (s *schemeCache) Remove(ctx context.Context, code string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.schemes, code)
	delete(s.schemaMapProcessInfo, code)
}
