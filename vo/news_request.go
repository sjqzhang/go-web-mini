package vo

import (
    "time"
)





// 查询News 新闻
type News struct {
 
    ID *int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
     
    CreatedAt *time.Time `json:"created_at" form:"created_at"` // created_at
     
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"` // updated_at
     
    DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"` // deleted_at
     
    Title *string `json:"title" form:"title"` // 新闻标题
     
    Content *string `json:"content" form:"content"` // 新闻内容
     
    Creator *string `json:"creator" form:"creator"` // 新闻创建者
     
}

// 查询News 新闻
type PagerNews struct {
	Total    int64                  `json:"total"`
	List     []News          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询News 新闻
type ListNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title"  form:"title"` // 新闻标题
    
     
    Content *string `json:"content"  form:"content"` // 新闻内容
    
     
    Creator *string `json:"creator"  form:"creator"` // 新闻创建者
    
     
     PageNum  *uint   `json:"pageNum" form:"pageNum"` //第几页
     PageSize *uint   `json:"pageSize" form:"pageSize"` //每页多少条
}


// 创建News 新闻
type CreateNewsRequest struct {
    
     
     
     
     
    Title *string `json:"title" form:"title"` // 新闻标题
    
     
    Content *string `json:"content" form:"content"` // 新闻内容
    
     
    Creator *string `json:"creator" form:"creator"` // 新闻创建者
    
     
}


// 更新News 新闻
type UpdateNewsRequest struct {
    ID      *int `json:""`
    
     
     
     
     
    Title *string `json:"title" form:"title"` // 新闻标题
    
     
    Content *string `json:"content" form:"content"` // 新闻内容
    
     
    Creator *string `json:"creator" form:"creator"` // 新闻创建者
    
     
}

// 删除News 新闻
type DeleteNewsRequest struct {
    Ids      []int64 `json:"ids" uri:"ids" form:"ids"` //待编号
}

// 删除News 新闻
type GetNewsRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"` //待编号
}



//以下结构体只用于生成swagger文档
type ListNewsResponse  struct {
    Response
    Data PagerNews  `json:"data"`
}

type GetNewsResponse  struct {
    Response
    Data News  `json:"data"`
}

type CreateNewsResponse  struct {
    Response
    Data News  `json:"data"`
}

type UpdateNewsResponse  struct {
    Response
    Data News  `json:"data"`
}

type DeleteNewsResponse  struct {
    Response
    Data int `json:"data"`
}

