package model

import (
	"time"
)

// News news
type News struct {
	Model
	Title   string `gorm:"title;type:longtext;comment:'标题'" validate:"required,checkMobile" json:"title"` // title:"标题";type:"input";validate:"required,checkMobile"
	Content string `gorm:"content;type:longtext;comment:'content'" validate:"" json:"content"`            // '内容001'
	Creator string `gorm:"creator;type:longtext;comment:'创建人'" validate:"" json:"creator"`                // title:"创建人";type:"input";validate:"

}

// News news
type NewsQuery struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt *time.Time `json:"created_at"`                           // created_at
	UpdatedAt *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt *time.Time `json:"deleted_at"`                           // deleted_at
	Title     *string    `json:"title"`                                // title:"标题";type:"input";validate:"required,checkMobile"
	Content   *string    `json:"content"`                              // '内容001'
	Creator   *string    `json:"creator"`                              // title:"创建人";type:"input";validate:"
	PageNum   int        `json:"-" form:"pageNum"`
	PageSize  int        `json:"-" form:"pageSize"`
}

func (t News) TableName() string {
	return "news"
}
