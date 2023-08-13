package vo

import (
	"time"
)

type ModuleResponse struct {
	ID           *int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // 编号
	ModuleName   *string    `json:"module_name" form:"module_name"`                 // 模块名称
	Tablename    *string    `json:"tablename" form:"tablename"`                     // 表名称
	ModelName    *string    `json:"model_name" form:"model_name"`                   // 模型名称
	TableComment *string    `json:"table_comment" form:"table_comment"`             // 表注释
	TableType    *string    `json:"table_type" form:"table_type"`                   // 表类型
	Pk           *string    `json:"pk" form:"pk"`                                   // 主键
	ListFields   *string    `json:"list_fields" form:"list_fields"`                 // 查询字段
	Remark       *string    `json:"remark" form:"remark"`                           // 备注
	Sort         *int32     `json:"sort" form:"sort"`                               // ''排序''
	IsSort       *int32     `json:"is_sort" form:"is_sort"`                         // ''排序字段''
	IsStatus     *int32     `json:"is_status" form:"is_status"`                     // ''状态字段''
	TopButton    *string    `json:"top_button" form:"top_button"`                   // 顶部按钮
	RightButton  *string    `json:"right_button" form:"right_button"`               // 右侧按钮
	IsSingle     *int32     `json:"is_single" form:"is_single"`                     // ''单页模式''
	ShowAll      *int32     `json:"show_all" form:"show_all"`                       // ''查看全部''
	AddParam     *string    `json:"add_param" form:"add_param"`                     // 添加参数
	CreatedAt    *time.Time `json:"created_at" form:"created_at"`                   // created_at
	DeletedAt    *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	UpdatedAt    *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at

}

// 查询Module 模块配置表
type Module struct {
	ID           *int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // 编号
	ModuleName   *string    `json:"module_name" form:"module_name"`                 // 模块名称
	Tablename    *string    `json:"tablename" form:"tablename"`                     // 表名称
	ModelName    *string    `json:"model_name" form:"model_name"`                   // 模型名称
	TableComment *string    `json:"table_comment" form:"table_comment"`             // 表注释
	TableType    *string    `json:"table_type" form:"table_type"`                   // 表类型
	Pk           *string    `json:"pk" form:"pk"`                                   // 主键
	ListFields   *string    `json:"list_fields" form:"list_fields"`                 // 查询字段
	Remark       *string    `json:"remark" form:"remark"`                           // 备注
	Sort         *int32     `json:"sort" form:"sort"`                               // ''排序''
	IsSort       *int32     `json:"is_sort" form:"is_sort"`                         // ''排序字段''
	IsStatus     *int32     `json:"is_status" form:"is_status"`                     // ''状态字段''
	TopButton    *string    `json:"top_button" form:"top_button"`                   // 顶部按钮
	RightButton  *string    `json:"right_button" form:"right_button"`               // 右侧按钮
	IsSingle     *int32     `json:"is_single" form:"is_single"`                     // ''单页模式''
	ShowAll      *int32     `json:"show_all" form:"show_all"`                       // ''查看全部''
	AddParam     *string    `json:"add_param" form:"add_param"`                     // 添加参数
	CreatedAt    *time.Time `json:"created_at" form:"created_at"`                   // created_at
	DeletedAt    *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	UpdatedAt    *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at

}

// 查询Module 模块配置表
type ListModuleResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []Module               `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 查询Module 模块配置表
type ListModuleRequest struct {
	ModuleName   *string `json:"module_name"  form:"module_name"`     // 模块名称
	Tablename    *string `json:"tablename"  form:"tablename"`         // 表名称
	ModelName    *string `json:"model_name"  form:"model_name"`       // 模型名称
	TableComment *string `json:"table_comment"  form:"table_comment"` // 表注释
	TableType    *string `json:"table_type"  form:"table_type"`       // 表类型
	Pk           *string `json:"pk"  form:"pk"`                       // 主键
	ListFields   *string `json:"list_fields"  form:"list_fields"`     // 查询字段
	Remark       *string `json:"remark"  form:"remark"`               // 备注
	Sort         *int32  `json:"sort"  form:"sort"`                   // ''排序''
	IsSort       *int32  `json:"is_sort"  form:"is_sort"`             // ''排序字段''
	IsStatus     *int32  `json:"is_status"  form:"is_status"`         // ''状态字段''
	TopButton    *string `json:"top_button"  form:"top_button"`       // 顶部按钮
	RightButton  *string `json:"right_button"  form:"right_button"`   // 右侧按钮
	IsSingle     *int32  `json:"is_single"  form:"is_single"`         // ''单页模式''
	ShowAll      *int32  `json:"show_all"  form:"show_all"`           // ''查看全部''
	AddParam     *string `json:"add_param"  form:"add_param"`         // 添加参数

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

type GetModuleResponse struct {
	ModuleResponse
}

// 创建Module 模块配置表
type CreateModuleRequest struct {
	ModuleName   *string `json:"module_name" form:"module_name"`     // 模块名称
	Tablename    *string `json:"tablename" form:"tablename"`         // 表名称
	ModelName    *string `json:"model_name" form:"model_name"`       // 模型名称
	TableComment *string `json:"table_comment" form:"table_comment"` // 表注释
	TableType    *string `json:"table_type" form:"table_type"`       // 表类型
	Pk           *string `json:"pk" form:"pk"`                       // 主键
	ListFields   *string `json:"list_fields" form:"list_fields"`     // 查询字段
	Remark       *string `json:"remark" form:"remark"`               // 备注
	Sort         *int32  `json:"sort" form:"sort"`                   // ''排序''
	IsSort       *int32  `json:"is_sort" form:"is_sort"`             // ''排序字段''
	IsStatus     *int32  `json:"is_status" form:"is_status"`         // ''状态字段''
	TopButton    *string `json:"top_button" form:"top_button"`       // 顶部按钮
	RightButton  *string `json:"right_button" form:"right_button"`   // 右侧按钮
	IsSingle     *int32  `json:"is_single" form:"is_single"`         // ''单页模式''
	ShowAll      *int32  `json:"show_all" form:"show_all"`           // ''查看全部''
	AddParam     *string `json:"add_param" form:"add_param"`         // 添加参数

}

type CreateModuleResponse struct {
	ModuleResponse
}

// 更新Module 模块配置表
type UpdateModuleRequest struct {
	ID           *int    `json:""`
	ModuleName   *string `json:"module_name" form:"module_name"`     // 模块名称
	Tablename    *string `json:"tablename" form:"tablename"`         // 表名称
	ModelName    *string `json:"model_name" form:"model_name"`       // 模型名称
	TableComment *string `json:"table_comment" form:"table_comment"` // 表注释
	TableType    *string `json:"table_type" form:"table_type"`       // 表类型
	Pk           *string `json:"pk" form:"pk"`                       // 主键
	ListFields   *string `json:"list_fields" form:"list_fields"`     // 查询字段
	Remark       *string `json:"remark" form:"remark"`               // 备注
	Sort         *int32  `json:"sort" form:"sort"`                   // ''排序''
	IsSort       *int32  `json:"is_sort" form:"is_sort"`             // ''排序字段''
	IsStatus     *int32  `json:"is_status" form:"is_status"`         // ''状态字段''
	TopButton    *string `json:"top_button" form:"top_button"`       // 顶部按钮
	RightButton  *string `json:"right_button" form:"right_button"`   // 右侧按钮
	IsSingle     *int32  `json:"is_single" form:"is_single"`         // ''单页模式''
	ShowAll      *int32  `json:"show_all" form:"show_all"`           // ''查看全部''
	AddParam     *string `json:"add_param" form:"add_param"`         // 添加参数

}

type UpdateModuleResponse struct {
	ModuleResponse
}

// 删除Module 模块配置表
type DeleteModuleRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除Module 模块配置表
type GetModuleRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteModuleResponse struct {
	Response
	Data int `json:"data"`
}
