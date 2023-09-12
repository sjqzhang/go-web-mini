package vo

import (
	"time"
)

var _ = time.Now()

type DictionaryResponse struct {
	ID         *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt  *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt  *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt  *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	DictLabel  *string    `json:"dict_label" form:"dict_label"`                   // 字典标签
	DictValue  *string    `json:"dict_value" form:"dict_value"`                   // 字典键值
	DictTypeId *int64     `json:"dict_type_id" form:"dict_type_id"`               // 字典类型
	Remark     *string    `json:"remark" form:"remark"`                           // 备注
	Sort       *int32     `json:"sort" form:"sort"`                               // 排序
	Status     *int32     `json:"status" form:"status"`                           // 状态

}

// 查询Dictionary dictionary
type Dictionary struct {
	ID         *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt  *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt  *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt  *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	DictLabel  *string    `json:"dict_label" form:"dict_label"`                   // 字典标签
	DictValue  *string    `json:"dict_value" form:"dict_value"`                   // 字典键值
	DictTypeId *int64     `json:"dict_type_id" form:"dict_type_id"`               // 字典类型
	Remark     *string    `json:"remark" form:"remark"`                           // 备注
	Sort       *int32     `json:"sort" form:"sort"`                               // 排序
	Status     *int32     `json:"status" form:"status"`                           // 状态
	DictType DictionaryType ` json:"dict_type"`
}

// 查询Dictionary dictionary
type ListForPagerDictionaryResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []Dictionary           `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 分页查询Dictionary dictionary
type ListForPagerDictionaryRequest struct {
	ID         *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"  form:"id"` // id
	CreatedAt  *time.Time `json:"created_at"  form:"created_at"`                   // created_at
	UpdatedAt  *time.Time `json:"updated_at"  form:"updated_at"`                   // updated_at
	DeletedAt  *time.Time `json:"deleted_at"  form:"deleted_at"`                   // deleted_at
	DictLabel  *string    `json:"dict_label"  form:"dict_label"`                   // 字典标签
	DictValue  *string    `json:"dict_value"  form:"dict_value"`                   // 字典键值
	DictTypeId *int64     `json:"dict_type_id"  form:"dict_type_id"`               // 字典类型
	Remark     *string    `json:"remark"  form:"remark"`                           // 备注
	Sort       *int32     `json:"sort"  form:"sort"`                               // 排序
	Status     *int32     `json:"status"  form:"status"`                           // 状态


	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

// 查询Dictionary dictionary
type ListDictionaryRequest struct {
	ID         *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"  form:"id"` // id
	CreatedAt  *time.Time `json:"created_at"  form:"created_at"`                   // created_at
	UpdatedAt  *time.Time `json:"updated_at"  form:"updated_at"`                   // updated_at
	DeletedAt  *time.Time `json:"deleted_at"  form:"deleted_at"`                   // deleted_at
	DictLabel  *string    `json:"dict_label"  form:"dict_label"`                   // 字典标签
	DictValue  *string    `json:"dict_value"  form:"dict_value"`                   // 字典键值
	DictTypeId *int64     `json:"dict_type_id"  form:"dict_type_id"`               // 字典类型
	Remark     *string    `json:"remark"  form:"remark"`                           // 备注
	Sort       *int32     `json:"sort"  form:"sort"`                               // 排序
	Status     *int32     `json:"status"  form:"status"`                           // 状态

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

// 查询Dictionary dictionary
type ListDictionaryResponse struct {
	List []Dictionary `json:"list"` //列表
}

type GetDictionaryResponse struct {
	DictionaryResponse
}

// 创建Dictionary dictionary
type CreateDictionaryRequest struct {
	DictLabel  *string `json:"dict_label" form:"dict_label" validate:""`     // 字典标签
	DictValue  *string `json:"dict_value" form:"dict_value" validate:""`     // 字典键值
	DictTypeId *int64  `json:"dict_type_id" form:"dict_type_id" validate:""` // 字典类型
	Remark     *string `json:"remark" form:"remark" validate:""`             // 备注
	Sort       *int32  `json:"sort" form:"sort" validate:""`                 // 排序
	Status     *int32  `json:"status" form:"status" validate:""`             // 状态

}

type CreateDictionaryResponse struct {
	DictionaryResponse
}

// 更新Dictionary dictionary
type UpdateDictionaryRequest struct {
	ID         *int    `json:""`
	DictLabel  *string `json:"dict_label" validate:"" form:"dict_label"`     // 字典标签
	DictValue  *string `json:"dict_value" validate:"" form:"dict_value"`     // 字典键值
	DictTypeId *int64  `json:"dict_type_id" validate:"" form:"dict_type_id"` // 字典类型
	Remark     *string `json:"remark" validate:"" form:"remark"`             // 备注
	Sort       *int32  `json:"sort" validate:"" form:"sort"`                 // 排序
	Status     *int32  `json:"status" validate:"" form:"status"`             // 状态

}

type UpdateDictionaryResponse struct {
	DictionaryResponse
}

// 删除Dictionary dictionary
type DeleteDictionaryRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除Dictionary dictionary
type GetDictionaryRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteDictionaryResponse struct {
	Response
	Data int `json:"data"`
}
