package model

type SchemeTab struct {
	Id         int    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Scheme     string `gorm:"column:scheme;type:text;NOT NULL" json:"scheme"`
	SchemeCode string `gorm:"column:scheme_code;type:varchar(128);NOT NULL" json:"scheme_code"`
	IsObsolete int    `gorm:"column:is_obsolete;type:tinyint(1);NOT NULL" json:"is_obsolete"`
	Ctime      int    `gorm:"column:ctime;type:int(10) unsigned;NOT NULL" json:"ctime"`
	Mtime      int    `gorm:"column:mtime;type:int(10) unsigned;NOT NULL" json:"mtime"`
}

func (obj *SchemeTab) TableName() string {
	return "bpmn_scheme_tab"
}
