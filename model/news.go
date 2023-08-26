package model

import (
	"time"
)

// News news
type News struct {
	Model
	Title   string `gorm:"uniqueIndex:unq_ix;title;type:varchar(50);comment:title" validate:"min=1,max=50" json:"title"`       // title:"title";validate:"min=1,max=50"
	Content string `gorm:"uniqueIndex:unq_ix;content;type:varchar(50);comment:content" validate:"min=1,max=50" json:"content"` // title:"content";validate:"min=1,max=50"
	Creator string `gorm:"creator;type:varchar(100);comment:creator" validate:"min=1,max=100" json:"creator"`                  // title:"creator";validate:"min=1,max=100"

}

// News news
type NewsQuery struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt *time.Time `json:"created_at"`                           // created_at
	UpdatedAt *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt *time.Time `json:"deleted_at"`                           // deleted_at
	Title     *string    `json:"title"`                                // title:"title";validate:"min=1,max=50"
	Content   *string    `json:"content"`                              // title:"content";validate:"min=1,max=50"
	Creator   *string    `json:"creator"`                              // title:"creator";validate:"min=1,max=100"
	PageNum   int        `json:"-" form:"pageNum"`
	PageSize  int        `json:"-" form:"pageSize"`
}

func (t News) TableName() string {
	return "news"
}
