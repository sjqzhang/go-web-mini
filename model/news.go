package model

import (
	"time"
)

// News news
type News struct {
	Model
	Title   string `gorm:"title;type:longtext;comment:'标题'" validate:"required,min=2,max=35" json:"title"` // title:"标题";type:"input";validate:"required,min=2,max=35"
	Content string `gorm:"content;type:longtext;comment:'内容'" validate:"" json:"content"`                  // title:"内容";type:"input";validate:"
	Creator string `gorm:"creator;type:longtext;comment:'创建人'" validate:"" json:"creator"`                 // title:"创建人";type:"input";validate:"

}

// News news
type NewsQuery struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt *time.Time `json:"created_at"`                           // created_at
	UpdatedAt *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt *time.Time `json:"deleted_at"`                           // deleted_at
	Title     *string    `json:"title"`                                // title:"标题";type:"input";validate:"required,min=2,max=35"
	Content   *string    `json:"content"`                              // title:"内容";type:"input";validate:"
	Creator   *string    `json:"creator"`                              // title:"创建人";type:"input";validate:"
	PageNum   int        `json:"-" form:"pageNum"`
	PageSize  int        `json:"-" form:"pageSize"`
}

func (t News) TableName() string {
	return "news"
}
