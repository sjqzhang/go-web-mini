package model

import (
	"time"
)

var _ = time.Now()

// DictionaryType dictionary_type
type DictionaryType struct {
	Model
	DictName string `gorm:"dict_name;type:char(100);comment:字典名称" validate:"" json:"dict_name"`   // 字典名称
	Status   int32  `gorm:"status;type:tinyint(4) unsigned;comment:状态" validate:"" json:"status"` // 状态
	Remark   string `gorm:"remark;type:varchar(200);comment:备注" validate:"" json:"remark"`        // 备注
	Sort     int32  `gorm:"sort;type:int(5) unsigned;comment:排序" validate:"" json:"sort"`         // 排序
	//Dictionaries []Dictionary `gorm:"-;foreignKey:ID" json:"-"`
	Dictionaries []Dictionary `gorm:"foreignKey:DictTypeId;" json:"dictionaries"`

}

// DictionaryType dictionary_type
type DictionaryTypeQuery struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt *time.Time `json:"created_at"`                           // created_at
	UpdatedAt *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt *time.Time `json:"deleted_at"`                           // deleted_at
	DictName  *string    `json:"dict_name"`                            // 字典名称
	Status    *int32     `json:"status"`                               // 状态
	Remark    *string    `json:"remark"`                               // 备注
	Sort      *int32     `json:"sort"`                                 // 排序
	PageNum   int        `json:"-" form:"pageNum"`
	PageSize  int        `json:"-" form:"pageSize"`
}

func (t DictionaryType) TableName() string {
	return "sys_dictionary_type"
}
