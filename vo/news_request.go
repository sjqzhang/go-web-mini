package vo

import (
    "time"
)


// 查询News 
type News struct {
 
    ID *int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"` 
     
    CreatedAt *time.Time `json:"created_at"` 
     
    UpdatedAt *time.Time `json:"updated_at"` 
     
    DeletedAt *time.Time `json:"deleted_at"` 
     
    Title *string `json:"title"` 
     
    Content *string `json:"content"` 
     
    Creator *string `json:"creator"` 
     
}


// 查询News 
type ListNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"` 
    
     
    Content *string `json:"content"` 
    
     
    Creator *string `json:"creator"` 
    
     
     PageNum  *uint   `json:"pageNum" form:"pageNum"`
     PageSize *uint   `json:"pageSize" form:"pageSize"`
}


// 创建News 
type CreateNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"` 
    
     
    Content *string `json:"content"` 
    
     
    Creator *string `json:"creator"` 
    
     
}


// 更新News 
type UpdateNewsRequest struct {
    ID      *int `json:""`
    
     
     
     
     
    Title *string `json:"title"` 
    
     
    Content *string `json:"content"` 
    
     
    Creator *string `json:"creator"` 
    
     
}

// 删除News 
type DeleteNewsRequest struct {
    ID      int `json:"id" uri:"id" form:"id"`
}

