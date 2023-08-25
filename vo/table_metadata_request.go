package vo

import (
	"time"
)

type TableMetadataResponse struct {
	ID                 *int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"`   // id
	TableAlias         *string    `json:"table_alias" form:"table_alias"`                   // table_alias
	ColumnName         *string    `json:"column_name" form:"column_name"`                   // column_name
	ColumnComment      *string    `json:"column_comment" form:"column_comment"`             // column_comment
	IsNullable         *string    `json:"is_nullable" form:"is_nullable"`                   // is_nullable
	ColumnType         *string    `json:"column_type" form:"column_type"`                   // column_type
	DataType           *string    `json:"data_type" form:"data_type"`                       // data_type
	CharacterMaxLength *string    `json:"character_max_length" form:"character_max_length"` // character_max_length
	UpdatedAt          *time.Time `json:"updated_at" form:"updated_at"`                     // updated_at
	CreatedAt          *time.Time `json:"created_at" form:"created_at"`                     // created_at
	DeletedAt          *time.Time `json:"deleted_at" form:"deleted_at"`                     // deleted_at

}

// 查询TableMetadata table_metadata
type TableMetadata struct {
	ID                 *int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"`   // id
	TableAlias         *string    `json:"table_alias" form:"table_alias"`                   // table_alias
	ColumnName         *string    `json:"column_name" form:"column_name"`                   // column_name
	ColumnComment      *string    `json:"column_comment" form:"column_comment"`             // column_comment
	IsNullable         *string    `json:"is_nullable" form:"is_nullable"`                   // is_nullable
	ColumnType         *string    `json:"column_type" form:"column_type"`                   // column_type
	DataType           *string    `json:"data_type" form:"data_type"`                       // data_type
	CharacterMaxLength *string    `json:"character_max_length" form:"character_max_length"` // character_max_length
	UpdatedAt          *time.Time `json:"updated_at" form:"updated_at"`                     // updated_at
	CreatedAt          *time.Time `json:"created_at" form:"created_at"`                     // created_at
	DeletedAt          *time.Time `json:"deleted_at" form:"deleted_at"`                     // deleted_at

}

// 查询TableMetadata table_metadata
type ListTableMetadataResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []TableMetadata        `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 查询TableMetadata table_metadata
type ListTableMetadataRequest struct {
	TableAlias         *string `json:"table_alias"  form:"table_alias"`                   // table_alias
	ColumnName         *string `json:"column_name"  form:"column_name"`                   // column_name
	ColumnComment      *string `json:"column_comment"  form:"column_comment"`             // column_comment
	IsNullable         *string `json:"is_nullable"  form:"is_nullable"`                   // is_nullable
	ColumnType         *string `json:"column_type"  form:"column_type"`                   // column_type
	DataType           *string `json:"data_type"  form:"data_type"`                       // data_type
	CharacterMaxLength *string `json:"character_max_length"  form:"character_max_length"` // character_max_length

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

type GetTableMetadataResponse struct {
	TableMetadataResponse
}

// 创建TableMetadata table_metadata
type CreateTableMetadataRequest struct {
	TableAlias         *string `json:"table_alias" form:"table_alias" validate:""`                   // table_alias
	ColumnName         *string `json:"column_name" form:"column_name" validate:""`                   // column_name
	ColumnComment      *string `json:"column_comment" form:"column_comment" validate:""`             // column_comment
	IsNullable         *string `json:"is_nullable" form:"is_nullable" validate:""`                   // is_nullable
	ColumnType         *string `json:"column_type" form:"column_type" validate:""`                   // column_type
	DataType           *string `json:"data_type" form:"data_type" validate:""`                       // data_type
	CharacterMaxLength *string `json:"character_max_length" form:"character_max_length" validate:""` // character_max_length

}

type CreateTableMetadataResponse struct {
	TableMetadataResponse
}

// 更新TableMetadata table_metadata
type UpdateTableMetadataRequest struct {
	ID                 *int    `json:""`
	TableAlias         *string `json:"table_alias" validate:"" form:"table_alias"`                   // table_alias
	ColumnName         *string `json:"column_name" validate:"" form:"column_name"`                   // column_name
	ColumnComment      *string `json:"column_comment" validate:"" form:"column_comment"`             // column_comment
	IsNullable         *string `json:"is_nullable" validate:"" form:"is_nullable"`                   // is_nullable
	ColumnType         *string `json:"column_type" validate:"" form:"column_type"`                   // column_type
	DataType           *string `json:"data_type" validate:"" form:"data_type"`                       // data_type
	CharacterMaxLength *string `json:"character_max_length" validate:"" form:"character_max_length"` // character_max_length

}

type UpdateTableMetadataResponse struct {
	TableMetadataResponse
}

// 删除TableMetadata table_metadata
type DeleteTableMetadataRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除TableMetadata table_metadata
type GetTableMetadataRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteTableMetadataResponse struct {
	Response
	Data int `json:"data"`
}
