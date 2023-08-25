package model

import (
	"time"
)

// TableMetadata table_metadata
type TableMetadata struct {
	Model
	//unique_index
	TableAlias         string `gorm:"uniqueIndex:uniq_name;table_alias;type:varchar(255);comment:'table_alias'" validate:"" json:"table_alias"`                          // table_alias
	ColumnName         string `gorm:"uniqueIndex:uniq_name;column_name;type:varchar(255);comment:'column_name'" validate:"" json:"column_name"`                          // column_name
	ColumnComment      string `gorm:"column_comment;type:text;comment:'column_comment'" validate:"" json:"column_comment"`                         // column_comment
	IsNullable         string `gorm:"is_nullable;type:varchar(3);comment:'is_nullable'" validate:"" json:"is_nullable"`                            // is_nullable
	ColumnType         string `gorm:"column_type;type:varchar(128);comment:'column_type'" validate:"" json:"column_type"`                          // column_type
	DataType           string `gorm:"data_type;type:varchar(255);comment:'data_type'" validate:"" json:"data_type"`                                // data_type
	CharacterMaxLength string `gorm:"character_max_length;type:mediumtext;comment:'character_max_length'" validate:"" json:"character_max_length"` // character_max_length

}

// TableMetadata table_metadata
type TableMetadataQuery struct {
	ID                 *int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	TableAlias         *string    `json:"table_alias"`                          // table_alias
	ColumnName         *string    `json:"column_name"`                          // column_name
	ColumnComment      *string    `json:"column_comment"`                       // column_comment
	IsNullable         *string    `json:"is_nullable"`                          // is_nullable
	ColumnType         *string    `json:"column_type"`                          // column_type
	DataType           *string    `json:"data_type"`                            // data_type
	CharacterMaxLength *string    `json:"character_max_length"`                 // character_max_length
	UpdatedAt          *time.Time `json:"updated_at"`                           // updated_at
	CreatedAt          *time.Time `json:"created_at"`                           // created_at
	DeletedAt          *time.Time `json:"deleted_at"`                           // deleted_at
	PageNum            int        `json:"-" form:"pageNum"`
	PageSize           int        `json:"-" form:"pageSize"`
}

func (t TableMetadata) TableName() string {
	return "table_metadata"
}
