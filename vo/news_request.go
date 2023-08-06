package vo

import (
    "time"
)


// 查询News news
type News struct {
 
    ID *int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
     
    CreatedAt *time.Time `json:"created_at" form:"created_at"` // created_at
     
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"` // updated_at
     
    DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"` // deleted_at
     
    Title *string `json:"title" form:"title"` // 新闻标题
     
    Content *string `json:"content" form:"content"` // 新闻内容
     
    Creator *string `json:"creator" form:"creator"` // 新闻创建者
     
}

// 查询News news
type PagerNews struct {
	Total    int64                  `json:"total"`
	List     []News          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询News news
type ListNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"  form:"title"` // 新闻标题
    
     
    Content *string `json:"content"  form:"content"` // 新闻内容
    
     
    Creator *string `json:"creator"  form:"creator"` // 新闻创建者
    
     
     PageNum  *uint   `json:"pageNum" form:"pageNum"`
     PageSize *uint   `json:"pageSize" form:"pageSize"`
}


// 创建News news
type CreateNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title" form:"title"` // 新闻标题
    
     
    Content *string `json:"content" form:"content"` // 新闻内容
    
     
    Creator *string `json:"creator" form:"creator"` // 新闻创建者
    
     
}


// 更新News news
type UpdateNewsRequest struct {
    ID      *int `json:""`
    
     
     
     
     
    Title *string `json:"title" form:"title"` // 新闻标题
    
     
    Content *string `json:"content" form:"content"` // 新闻内容
    
     
    Creator *string `json:"creator" form:"creator"` // 新闻创建者
    
     
}

// 删除News news
type DeleteNewsRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"`
}

// 删除News news
type GetNewsRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"`
}

