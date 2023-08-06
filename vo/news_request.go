package vo

import (
    "time"
)


// 查询News 
type News struct {
 
    ID *int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` 
     
    CreatedAt *time.Time `json:"created_at" form:"created_at"` 
     
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"` 
     
    DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"` 
     
    Title *string `json:"title" form:"title"` 
     
    Content *string `json:"content" form:"content"` 
     
    Creator *string `json:"creator" form:"creator"` 
     
}

// 查询News 
type PagerNews struct {
	Total    int64                  `json:"total"`
	List     []News          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询News 
type ListNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"  form:"title"` 
    
     
    Content *string `json:"content"  form:"content"` 
    
     
    Creator *string `json:"creator"  form:"creator"` 
    
     
     PageNum  *uint   `json:"pageNum" form:"pageNum"`
     PageSize *uint   `json:"pageSize" form:"pageSize"`
}


// 创建News 
type CreateNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title" form:"title"` 
    
     
    Content *string `json:"content" form:"content"` 
    
     
    Creator *string `json:"creator" form:"creator"` 
    
     
}


// 更新News 
type UpdateNewsRequest struct {
    ID      *int `json:""`
    
     
     
     
     
    Title *string `json:"title" form:"title"` 
    
     
    Content *string `json:"content" form:"content"` 
    
     
    Creator *string `json:"creator" form:"creator"` 
    
     
}

// 删除News 
type DeleteNewsRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"`
}

// 删除News 
type GetNewsRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"`
}

