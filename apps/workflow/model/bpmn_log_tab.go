package model

type LogTab struct {
	Id         int    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	InstanceId int    `gorm:"column:instance_id;type:bigint(20) unsigned;NOT NULL" json:"instance_id"`
	SchemeCode string `gorm:"column:scheme_code;type:varchar(128);NOT NULL" json:"scheme_code"`
	Message    string `gorm:"column:message;type:text;NOT NULL" json:"message"`
	Command    string `gorm:"column:command;type:varchar(256);NOT NULL" json:"command"`
	Vars       string `gorm:"column:vars;type:text;NOT NULL" json:"vars"`
	Ctime      int    `gorm:"column:ctime;type:int(10) unsigned;NOT NULL" json:"ctime"`
}

func (m *LogTab) TableName() string {
	return "bpmn_log_tab"
}
