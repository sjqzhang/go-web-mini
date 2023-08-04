package vo

import (
    "time"
)


// 查询News 
type News struct {
 
    Id *int64 `json:"id" gorm:"primary_key"` 
     
    CreatedAt *time.Time `json:"created_at"` 
     
    UpdatedAt *time.Time `json:"updated_at"` 
     
    DeletedAt *time.Time `json:"deleted_at"` 
     
    Title *string `json:"title"` 
     
    Content *string `json:"content"` 
     
}


// 查询News 
type ListNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"` 
    
     
    Content *string `json:"content"` 
    
     
     PageNum  *uint   `json:"pageNum" form:"pageNum"`
     PageSize *uint   `json:"pageSize" form:"pageSize"`
}


// 创建News 
type CreateNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"` 
    
     
    Content *string `json:"content"` 
    
     
}


// 更新News 
type UpdateNewsRequest struct {
    ID      *int `json:""`
    
     
     
     
     
    Title *string `json:"title"` 
    
     
    Content *string `json:"content"` 
    
     
}

// 删除News 
type DeleteNewsRequest struct {
    ID      int `json:"id" uri:"id" form:"id"`
}

