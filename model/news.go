package model

import (

    "time"

)


// News 
type News struct {
    Model
    
     
     
     
     
    Title string `json:"title"` 
     
    Content string `json:"content"` 
     
}

// News 
type NewsQuery struct {
 
    Id *int64 `json:"id" gorm:"primary_key"` 
    CreatedAt *time.Time `json:"created_at"` 
    UpdatedAt *time.Time `json:"updated_at"` 
    DeletedAt *time.Time `json:"deleted_at"` 
    Title *string `json:"title"` 
    Content *string `json:"content"` 

          PageNum  int   `json:"pageNum" form:"pageNum"`
          PageSize int   `json:"pageSize" form:"pageSize"`
}

