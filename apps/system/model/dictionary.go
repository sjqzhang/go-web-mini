package model

import (
	"time"
)

var _ = time.Now()

// Dictionary dictionary
type Dictionary struct {
	Model
	DictLabel  string `gorm:"dict_label;type:varchar(100);comment:字典标签" validate:"" json:"dict_label"`   // 字典标签
	DictValue  string `gorm:"dict_value;type:varchar(255);comment:字典键值" validate:"" json:"dict_value"`   // 字典键值
	DictTypeId int64  `gorm:"dict_type_id;type:bigint(20);comment:字典类型" validate:"" json:"dict_type_id"` // 字典类型
	Remark     string `gorm:"remark;type:varchar(200);comment:备注" validate:"" json:"remark"`             // 备注
	Sort       int32  `gorm:"sort;type:int(5) unsigned;comment:排序" validate:"" json:"sort"`              // 排序
	Status     int32  `gorm:"status;type:tinyint(4) unsigned;comment:状态" validate:"" json:"status"`      // 状态
	DictType DictionaryType `gorm:"foreignKey:DictTypeId;references:ID" json:"-"`

}

// Dictionary dictionary
type DictionaryQuery struct {
	ID         *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt  *time.Time `json:"created_at"`                           // created_at
	UpdatedAt  *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt  *time.Time `json:"deleted_at"`                           // deleted_at
	DictLabel  *string    `json:"dict_label"`                           // 字典标签
	DictValue  *string    `json:"dict_value"`                           // 字典键值
	DictTypeId *int64     `json:"dict_type_id"`                         // 字典类型
	Remark     *string    `json:"remark"`                               // 备注
	Sort       *int32     `json:"sort"`                                 // 排序
	Status     *int32     `json:"status"`                               // 状态
	PageNum    int        `json:"-" form:"pageNum"`
	PageSize   int        `json:"-" form:"pageSize"`
}

func (t Dictionary) TableName() string {
	return "sys_dictionary"
}
