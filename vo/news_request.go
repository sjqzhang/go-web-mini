package vo

import (
	"time"
)

type NewsResponse struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	Title     *string    `json:"title" form:"title"`                             // title:"title";validate:"min=1,max=50"
	Content   *string    `json:"content" form:"content"`                         // title:"content";validate:"min=1,max=50"
	Creator   *string    `json:"creator" form:"creator"`                         // title:"creator";validate:"min=1,max=100"

}

// 查询News news
type News struct {
	ID        *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	Title     *string    `json:"title" form:"title"`                             // title:"title";validate:"min=1,max=50"
	Content   *string    `json:"content" form:"content"`                         // title:"content";validate:"min=1,max=50"
	Creator   *string    `json:"creator" form:"creator"`                         // title:"creator";validate:"min=1,max=100"

}

// 查询News news
type ListNewsResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []News                 `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 查询News news
type ListNewsRequest struct {
	Title   *string `json:"title"  form:"title"`     // title:"title";validate:"min=1,max=50"
	Content *string `json:"content"  form:"content"` // title:"content";validate:"min=1,max=50"
	Creator *string `json:"creator"  form:"creator"` // title:"creator";validate:"min=1,max=100"

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

type GetNewsResponse struct {
	NewsResponse
}

// 创建News news
type CreateNewsRequest struct {
	Title   *string `json:"title" form:"title" validate:"min=1,max=50"`      // title:"title";validate:"min=1,max=50"
	Content *string `json:"content" form:"content" validate:"min=1,max=50"`  // title:"content";validate:"min=1,max=50"
	Creator *string `json:"creator" form:"creator" validate:"min=1,max=100"` // title:"creator";validate:"min=1,max=100"

}

type CreateNewsResponse struct {
	NewsResponse
}

// 更新News news
type UpdateNewsRequest struct {
	ID      *int    `json:""`
	Title   *string `json:"title" validate:"min=1,max=50" form:"title"`      // title:"title";validate:"min=1,max=50"
	Content *string `json:"content" validate:"min=1,max=50" form:"content"`  // title:"content";validate:"min=1,max=50"
	Creator *string `json:"creator" validate:"min=1,max=100" form:"creator"` // title:"creator";validate:"min=1,max=100"

}

type UpdateNewsResponse struct {
	NewsResponse
}

// 删除News news
type DeleteNewsRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除News news
type GetNewsRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteNewsResponse struct {
	Response
	Data int `json:"data"`
}
