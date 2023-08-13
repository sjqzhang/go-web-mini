package model

import (
	"time"
)

// Module 模块配置表
type Module struct {
	Model
	ModuleName   string `gorm:"module_name;type:varchar(100);comment:'模块名称'" json:"module_name"`        // 模块名称
	Tablename    string `gorm:"tablename;type:varchar(50);comment:'表名称'" json:"tablename"`              // 表名称
	ModelName    string `gorm:"model_name;type:varchar(50);comment:'模型名称'" json:"model_name"`           // 模型名称
	TableComment string `gorm:"table_comment;type:varchar(200);comment:'表注释'" json:"table_comment"`     // 表注释
	TableType    string `gorm:"table_type;type:varchar(10);comment:'表类型'" json:"table_type"`            // 表类型
	Pk           string `gorm:"pk;type:varchar(50);comment:'主键'" json:"pk"`                             // 主键
	ListFields   string `gorm:"list_fields;type:varchar(255);comment:'查询字段'" json:"list_fields"`        // 查询字段
	Remark       string `gorm:"remark;type:text;comment:'备注'" json:"remark"`                            // 备注
	Sort         int32  `gorm:"sort;type:smallint(3) unsigned;comment:'''排序'''" json:"sort"`            // ''排序''
	IsSort       int32  `gorm:"is_sort;type:tinyint(4) unsigned;comment:'''排序字段'''" json:"is_sort"`     // ''排序字段''
	IsStatus     int32  `gorm:"is_status;type:tinyint(4) unsigned;comment:'''状态字段'''" json:"is_status"` // ''状态字段''
	TopButton    string `gorm:"top_button;type:varchar(255);comment:'顶部按钮'" json:"top_button"`          // 顶部按钮
	RightButton  string `gorm:"right_button;type:varchar(255);comment:'右侧按钮'" json:"right_button"`      // 右侧按钮
	IsSingle     int32  `gorm:"is_single;type:tinyint(4) unsigned;comment:'''单页模式'''" json:"is_single"` // ''单页模式''
	ShowAll      int32  `gorm:"show_all;type:tinyint(4) unsigned;comment:'''查看全部'''" json:"show_all"`   // ''查看全部''
	AddParam     string `gorm:"add_param;type:varchar(100);comment:'添加参数'" json:"add_param"`            // 添加参数

}

// Module 模块配置表
type ModuleQuery struct {
	ID           *int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 编号
	ModuleName   *string    `json:"module_name"`                          // 模块名称
	Tablename    *string    `json:"tablename"`                            // 表名称
	ModelName    *string    `json:"model_name"`                           // 模型名称
	TableComment *string    `json:"table_comment"`                        // 表注释
	TableType    *string    `json:"table_type"`                           // 表类型
	Pk           *string    `json:"pk"`                                   // 主键
	ListFields   *string    `json:"list_fields"`                          // 查询字段
	Remark       *string    `json:"remark"`                               // 备注
	Sort         *int32     `json:"sort"`                                 // ''排序''
	IsSort       *int32     `json:"is_sort"`                              // ''排序字段''
	IsStatus     *int32     `json:"is_status"`                            // ''状态字段''
	TopButton    *string    `json:"top_button"`                           // 顶部按钮
	RightButton  *string    `json:"right_button"`                         // 右侧按钮
	IsSingle     *int32     `json:"is_single"`                            // ''单页模式''
	ShowAll      *int32     `json:"show_all"`                             // ''查看全部''
	AddParam     *string    `json:"add_param"`                            // 添加参数
	CreatedAt    *time.Time `json:"created_at"`                           // created_at
	DeletedAt    *time.Time `json:"deleted_at"`                           // deleted_at
	UpdatedAt    *time.Time `json:"updated_at"`                           // updated_at
	PageNum      int        `json:"-" form:"pageNum"`
	PageSize     int        `json:"-" form:"pageSize"`
}

func (t Module) TableName() string {
	return "module"
}
