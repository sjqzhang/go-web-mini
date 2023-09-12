package vo

import (
	"time"
)

var _ = time.Now()

type DictionaryTypeResponse struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	DictName  *string    `json:"dict_name" form:"dict_name"`                     // 字典名称
	Status    *int32     `json:"status" form:"status"`                           // 状态
	Remark    *string    `json:"remark" form:"remark"`                           // 备注
	Sort      *int32     `json:"sort" form:"sort"`                               // 排序

}

// 查询DictionaryType dictionary_type
type DictionaryType struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	DictName  *string    `json:"dict_name" form:"dict_name"`                     // 字典名称
	Status    *int32     `json:"status" form:"status"`                           // 状态
	Remark    *string    `json:"remark" form:"remark"`                           // 备注
	Sort      *int32     `json:"sort" form:"sort"`                               // 排序
	Dictionaries []Dictionary ` json:"dictionaries"`

}

// 查询DictionaryType dictionary_type
type ListForPagerDictionaryTypeResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []DictionaryType       `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 分页查询DictionaryType dictionary_type
type ListForPagerDictionaryTypeRequest struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"  form:"id"` // id
	CreatedAt *time.Time `json:"created_at"  form:"created_at"`                   // created_at
	UpdatedAt *time.Time `json:"updated_at"  form:"updated_at"`                   // updated_at
	DeletedAt *time.Time `json:"deleted_at"  form:"deleted_at"`                   // deleted_at
	DictName  *string    `json:"dict_name"  form:"dict_name"`                     // 字典名称
	Status    *int32     `json:"status"  form:"status"`                           // 状态
	Remark    *string    `json:"remark"  form:"remark"`                           // 备注
	Sort      *int32     `json:"sort"  form:"sort"`                               // 排序

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

// 查询DictionaryType dictionary_type
type ListDictionaryTypeRequest struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"  form:"id"` // id
	CreatedAt *time.Time `json:"created_at"  form:"created_at"`                   // created_at
	UpdatedAt *time.Time `json:"updated_at"  form:"updated_at"`                   // updated_at
	DeletedAt *time.Time `json:"deleted_at"  form:"deleted_at"`                   // deleted_at
	DictName  *string    `json:"dict_name"  form:"dict_name"`                     // 字典名称
	Status    *int32     `json:"status"  form:"status"`                           // 状态
	Remark    *string    `json:"remark"  form:"remark"`                           // 备注
	Sort      *int32     `json:"sort"  form:"sort"`                               // 排序

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

// 查询DictionaryType dictionary_type
type ListDictionaryTypeResponse struct {
	List []DictionaryType `json:"list"` //列表
}

type GetDictionaryTypeResponse struct {
	DictionaryTypeResponse
}

// 创建DictionaryType dictionary_type
type CreateDictionaryTypeRequest struct {
	DictName *string `json:"dict_name" form:"dict_name" validate:""` // 字典名称
	Status   *int32  `json:"status" form:"status" validate:""`       // 状态
	Remark   *string `json:"remark" form:"remark" validate:""`       // 备注
	Sort     *int32  `json:"sort" form:"sort" validate:""`           // 排序

}

type CreateDictionaryTypeResponse struct {
	DictionaryTypeResponse
}

// 更新DictionaryType dictionary_type
type UpdateDictionaryTypeRequest struct {
	ID       *int    `json:""`
	DictName *string `json:"dict_name" validate:"" form:"dict_name"` // 字典名称
	Status   *int32  `json:"status" validate:"" form:"status"`       // 状态
	Remark   *string `json:"remark" validate:"" form:"remark"`       // 备注
	Sort     *int32  `json:"sort" validate:"" form:"sort"`           // 排序

}

type UpdateDictionaryTypeResponse struct {
	DictionaryTypeResponse
}

// 删除DictionaryType dictionary_type
type DeleteDictionaryTypeRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除DictionaryType dictionary_type
type GetDictionaryTypeRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteDictionaryTypeResponse struct {
	Response
	Data int `json:"data"`
}
