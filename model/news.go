package model

import (
	"time"
)

// News news
type News struct {
	Model
	Title   string `gorm:"title;type:longtext;comment:'title:"标题";type:"input";validate:"required:min=0,max=35"'" json:"title"` // title:"标题";type:"input";validate:"required:min=0,max=35"
	Content string `gorm:"content;type:longtext;comment:'content'" json:"content"`                                              // content
	Creator string `gorm:"creator;type:longtext;comment:'creator'" json:"creator"`                                              // creator

}

// News news
type NewsQuery struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt *time.Time `json:"created_at"`                           // created_at
	UpdatedAt *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt *time.Time `json:"deleted_at"`                           // deleted_at
	Title     *string    `json:"title"`                                // title:"标题";type:"input";validate:"required:min=0,max=35"
	Content   *string    `json:"content"`                              // content
	Creator   *string    `json:"creator"`                              // creator
	PageNum   int        `json:"-" form:"pageNum"`
	PageSize  int        `json:"-" form:"pageSize"`
}

func (t News) TableName() string {
	return "news"
}
