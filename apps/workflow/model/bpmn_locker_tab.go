package model

type LockerTab struct {
	Id         int    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	LockerName string `gorm:"column:locker_name;type:varchar(128);NOT NULL" json:"locker_name"`
	ExpireTime int    `gorm:"column:expire_time;type:int(10) unsigned;NOT NULL" json:"expire_time"`
	Ctime      int    `gorm:"column:ctime;type:int(10) unsigned;NOT NULL" json:"ctime"`
}

func (m *LockerTab) TableName() string {
	return "bpmn_locker_tab"
}
