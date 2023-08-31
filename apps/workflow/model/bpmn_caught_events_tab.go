package model

type CaughtEventsTab struct {
	Id         int    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	InstanceId int    `gorm:"column:instance_id;type:bigint(20) unsigned;NOT NULL" json:"instance_id"`
	EventName  string `gorm:"column:event_name;type:varchar(128);NOT NULL" json:"event_name"`
	IsConsumed int    `gorm:"column:is_consumed;type:tinyint(1);NOT NULL" json:"is_consumed"`
	Ctime      int    `gorm:"column:ctime;type:int(10) unsigned;NOT NULL" json:"ctime"`
	Mtime      int    `gorm:"column:mtime;type:int(10) unsigned;NOT NULL" json:"mtime"`
}

func (m *CaughtEventsTab) TableName() string {
	return "bpmn_caught_events_tab"
}

func (m *CaughtEventsTab) HasConsumed() bool {
	return m.IsConsumed == 1
}
