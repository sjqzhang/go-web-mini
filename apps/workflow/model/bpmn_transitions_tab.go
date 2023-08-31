package model

import "encoding/json"

type TransitionsTab struct {
	Id               int    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	InstanceId       int    `gorm:"column:instance_id;type:bigint(20) unsigned;NOT NULL" json:"instance_id"`
	IdentityId       string `gorm:"column:identity_id;type:varchar(256);NOT NULL" json:"identity_id"`
	FromActivityId   string `gorm:"column:from_activity_id;type:varchar(256);NOT NULL" json:"from_activity_id"`
	FromActivityName string `gorm:"column:from_activity_name;type:varchar(256);NOT NULL" json:"from_activity_name"`
	ToActivityId     string `gorm:"column:to_activity_id;type:varchar(256);NOT NULL" json:"to_activity_id"`
	ToActivityName   string `gorm:"column:to_activity_name;type:varchar(256);NOT NULL" json:"to_activity_name"`
	CommandId        string `gorm:"column:command_id;type:varchar(256);NOT NULL" json:"command_id"`
	CommandName      string `gorm:"column:command_name;type:varchar(256);NOT NULL" json:"command_name"`
	Remark           string `gorm:"column:remark;type:varchar(256);NOT NULL" json:"remark"`
	Vars             string `gorm:"column:vars;type:text" json:"vars"`
	Ctime            int    `gorm:"column:ctime;type:int(10) unsigned;NOT NULL" json:"ctime"`
}

func (m *TransitionsTab) TableName() string {
	return "bpmn_transitions_tab"
}

func (m *TransitionsTab) SetVariables(variables map[string]interface{}) {
	if variables == nil {
		m.Vars = ""
		return
	}
	s, _ := json.Marshal(variables)
	m.Vars = string(s)
}

func (m *TransitionsTab) GetVariables() map[string]interface{} {
	if m.Vars == "" {
		return nil
	}
	v := make(map[string]interface{})
	_ = json.Unmarshal([]byte(m.Vars), &v)
	return v
}
