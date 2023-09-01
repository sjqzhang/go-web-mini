package service

import (
	"context"
	"go-web-mini/apps/workflow/errcode"
	"go-web-mini/apps/workflow/repository"

	"github.com/sirupsen/logrus"
)

type VariablesService struct {
	InstanceRepo repository.InstanceRepo
}

func NewVariablesService() *VariablesService {
	return &VariablesService{InstanceRepo: &repository.InstanceDBRepo{}}
}

func (v *VariablesService) GetInstanceVariables(ctx context.Context, instanceId int, varName string) (name string, value string, e errcode.Exception) {
	instanceTab, err := v.InstanceRepo.Get(ctx, instanceId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": instanceId,
		}).Errorf("get instance from DB err: %+v", err)
		return "", "", errcode.Wrap(errcode.DBError, err)
	}

	name, value = "", ""
	vars := instanceTab.GetVariables()
	for k, v := range vars {
		if k == varName {
			name = k
			value = v.(string)
			break
		}
	}
	return name, value, nil
}

func (v *VariablesService) SetInstanceVariables(ctx context.Context, instanceId int, name, value string) (e errcode.Exception) {
	instanceTab, err := v.InstanceRepo.Get(ctx, instanceId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": instanceId,
		}).Errorf("get instance from DB err: %+v", err)
		return errcode.Wrap(errcode.DBError, err)
	}

	vars := instanceTab.GetVariables()
	if vars != nil {
		vars[name] = value
	} else {
		vars = make(map[string]interface{}, 1)
		vars[name] = value
	}
	instanceTab.SetVariables(vars)

	err = v.InstanceRepo.Update(ctx, instanceTab)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": instanceId,
		}).Errorf("update instance variables err: %+v", err)
		return errcode.Wrap(errcode.DBError, err)
	}
	return nil
}
