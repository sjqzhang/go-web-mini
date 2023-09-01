package model

import "encoding/json"

type InstanceTab struct {
	Id            int    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	SchemeCode    string `gorm:"column:scheme_code;type:varchar(128);NOT NULL" json:"scheme_code"`
	State         string `gorm:"column:state;type:varchar(64) unsigned;NOT NULL" json:"state"`
	Vars          string `gorm:"column:vars;type:text" json:"vars"`
	CurrentTaskId string `gorm:"column:current_task_id;type:varchar(128);NOT NULL" json:"current_task_id"`
	Ctime         int    `gorm:"column:ctime;type:int(10) unsigned;NOT NULL" json:"ctime"`
	Mtime         int    `gorm:"column:mtime;type:int(10) unsigned;NOT NULL" json:"mtime"`
}

func (obj *InstanceTab) TableName() string {
	return "bpmn_instance_tab"
}

func (obj *InstanceTab) GetVariables() map[string]interface{} {
	if obj.Vars == "" {
		return nil
	}
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(obj.Vars), &m)
	return m
}

func (obj *InstanceTab) SetVariables(variables map[string]interface{}) {
	if variables == nil {
		obj.Vars = ""
		return
	}
	s, _ := json.Marshal(variables)
	obj.Vars = string(s)
}
