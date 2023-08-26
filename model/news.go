package model

import (
	"time"
)

// News news
type News struct {
	Model
	Title   string `gorm:"title;type:varchar(50);comment:标题" validate:"required" json:"title"`                       // title:"标题";type:"input";validate:"required"
	Content string `gorm:"content;type:varchar(50);comment:content" validate:"min=1,max=4294967295" json:"content"`  // title:"content";type:"input";validate:"min=1,max=4294967295"
	Creator string `gorm:"creator;type:varchar(100);comment:creator" validate:"min=1,max=4294967295" json:"creator"` // title:"creator";type:"input";validate:"min=1,max=4294967295"

}

// News news
type NewsQuery struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt *time.Time `json:"created_at"`                           // created_at
	UpdatedAt *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt *time.Time `json:"deleted_at"`                           // deleted_at
	Title     *string    `json:"title"`                                // title:"标题";type:"input";validate:"required"
	Content   *string    `json:"content"`                              // title:"content";type:"input";validate:"min=1,max=4294967295"
	Creator   *string    `json:"creator"`                              // title:"creator";type:"input";validate:"min=1,max=4294967295"
	PageNum   int        `json:"-" form:"pageNum"`
	PageSize  int        `json:"-" form:"pageSize"`
}

func (t News) TableName() string {
	return "news"
}
