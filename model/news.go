package model

import (

    "time"

)


// News 新闻
type News struct {
    Model
    
     
     
     
     
    Title string `json:"title"` // 新闻标题
     
    Content string `json:"content"` // 新闻内容
     
    Creator string `json:"creator"` // 新闻创建者
     
}

// News 新闻
type NewsQuery struct {
 
    ID *int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
    CreatedAt *time.Time `json:"created_at"` // created_at
    UpdatedAt *time.Time `json:"updated_at"` // updated_at
    DeletedAt *time.Time `json:"deleted_at"` // deleted_at
    Title *string `json:"title"` // 新闻标题
    Content *string `json:"content"` // 新闻内容
    Creator *string `json:"creator"` // 新闻创建者

          PageNum  int   `json:"-" form:"pageNum"`
          PageSize int   `json:"-" form:"pageSize"`
}


func (t News)TableName() string {
    return "news"
}

