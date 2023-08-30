package vo

import (
	"time"
)

var _ = time.Now()

type ArticleResponse struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt   *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	Sort        *int32     `json:"sort" form:"sort"`                               // 排序
	Status      *int32     `json:"status" form:"status"`                           // 状态
	CateId      *int32     `json:"cate_id" form:"cate_id"`                         // 栏目
	Title       *string    `json:"title" form:"title"`                             // 标题
	Author      *string    `json:"author" form:"author"`                           // 作者
	Source      *string    `json:"source" form:"source"`                           // 来源
	Content     *string    `json:"content" form:"content"`                         // 内容
	Summary     *string    `json:"summary" form:"summary"`                         // 摘要
	Image       *string    `json:"image" form:"image"`                             // 图片
	Images      *string    `json:"images" form:"images"`                           // 图片集
	Download    *string    `json:"download" form:"download"`                       // 文件下载
	Tags        *string    `json:"tags" form:"tags"`                               // TAG
	Hits        *int32     `json:"hits" form:"hits"`                               // 点击次数
	Keywords    *string    `json:"keywords" form:"keywords"`                       // 关键词
	Description *string    `json:"description" form:"description"`                 // 描述
	Template    *string    `json:"template" form:"template"`                       // 模板
	Url         *string    `json:"url" form:"url"`                                 // 跳转地址
	ViewAuth    *int32     `json:"view_auth" form:"view_auth"`                     // 阅读权限

}

// 查询Article article
type Article struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt   *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	Sort        *int32     `json:"sort" form:"sort"`                               // 排序
	Status      *int32     `json:"status" form:"status"`                           // 状态
	CateId      *int32     `json:"cate_id" form:"cate_id"`                         // 栏目
	Title       *string    `json:"title" form:"title"`                             // 标题
	Author      *string    `json:"author" form:"author"`                           // 作者
	Source      *string    `json:"source" form:"source"`                           // 来源
	Content     *string    `json:"content" form:"content"`                         // 内容
	Summary     *string    `json:"summary" form:"summary"`                         // 摘要
	Image       *string    `json:"image" form:"image"`                             // 图片
	Images      *string    `json:"images" form:"images"`                           // 图片集
	Download    *string    `json:"download" form:"download"`                       // 文件下载
	Tags        *string    `json:"tags" form:"tags"`                               // TAG
	Hits        *int32     `json:"hits" form:"hits"`                               // 点击次数
	Keywords    *string    `json:"keywords" form:"keywords"`                       // 关键词
	Description *string    `json:"description" form:"description"`                 // 描述
	Template    *string    `json:"template" form:"template"`                       // 模板
	Url         *string    `json:"url" form:"url"`                                 // 跳转地址
	ViewAuth    *int32     `json:"view_auth" form:"view_auth"`                     // 阅读权限

}

// 查询Article article
type ListArticleResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []Article              `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 查询Article article
type ListArticleRequest struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"  form:"id"` // id
	CreatedAt   *time.Time `json:"created_at"  form:"created_at"`                   // created_at
	UpdatedAt   *time.Time `json:"updated_at"  form:"updated_at"`                   // updated_at
	DeletedAt   *time.Time `json:"deleted_at"  form:"deleted_at"`                   // deleted_at
	Sort        *int32     `json:"sort"  form:"sort"`                               // 排序
	Status      *int32     `json:"status"  form:"status"`                           // 状态
	CateId      *int32     `json:"cate_id"  form:"cate_id"`                         // 栏目
	Title       *string    `json:"title"  form:"title"`                             // 标题
	Author      *string    `json:"author"  form:"author"`                           // 作者
	Source      *string    `json:"source"  form:"source"`                           // 来源
	Content     *string    `json:"content"  form:"content"`                         // 内容
	Summary     *string    `json:"summary"  form:"summary"`                         // 摘要
	Image       *string    `json:"image"  form:"image"`                             // 图片
	Images      *string    `json:"images"  form:"images"`                           // 图片集
	Download    *string    `json:"download"  form:"download"`                       // 文件下载
	Tags        *string    `json:"tags"  form:"tags"`                               // TAG
	Hits        *int32     `json:"hits"  form:"hits"`                               // 点击次数
	Keywords    *string    `json:"keywords"  form:"keywords"`                       // 关键词
	Description *string    `json:"description"  form:"description"`                 // 描述
	Template    *string    `json:"template"  form:"template"`                       // 模板
	Url         *string    `json:"url"  form:"url"`                                 // 跳转地址
	ViewAuth    *int32     `json:"view_auth"  form:"view_auth"`                     // 阅读权限

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

type GetArticleResponse struct {
	ArticleResponse
}

// 创建Article article
type CreateArticleRequest struct {
	Sort        *int32  `json:"sort" form:"sort" validate:""`               // 排序
	Status      *int32  `json:"status" form:"status" validate:""`           // 状态
	CateId      *int32  `json:"cate_id" form:"cate_id" validate:""`         // 栏目
	Title       *string `json:"title" form:"title" validate:""`             // 标题
	Author      *string `json:"author" form:"author" validate:""`           // 作者
	Source      *string `json:"source" form:"source" validate:""`           // 来源
	Content     *string `json:"content" form:"content" validate:""`         // 内容
	Summary     *string `json:"summary" form:"summary" validate:""`         // 摘要
	Image       *string `json:"image" form:"image" validate:""`             // 图片
	Images      *string `json:"images" form:"images" validate:""`           // 图片集
	Download    *string `json:"download" form:"download" validate:""`       // 文件下载
	Tags        *string `json:"tags" form:"tags" validate:""`               // TAG
	Hits        *int32  `json:"hits" form:"hits" validate:""`               // 点击次数
	Keywords    *string `json:"keywords" form:"keywords" validate:""`       // 关键词
	Description *string `json:"description" form:"description" validate:""` // 描述
	Template    *string `json:"template" form:"template" validate:""`       // 模板
	Url         *string `json:"url" form:"url" validate:""`                 // 跳转地址
	ViewAuth    *int32  `json:"view_auth" form:"view_auth" validate:""`     // 阅读权限

}

type CreateArticleResponse struct {
	ArticleResponse
}

// 更新Article article
type UpdateArticleRequest struct {
	ID          *int    `json:""`
	Sort        *int32  `json:"sort" validate:"" form:"sort"`               // 排序
	Status      *int32  `json:"status" validate:"" form:"status"`           // 状态
	CateId      *int32  `json:"cate_id" validate:"" form:"cate_id"`         // 栏目
	Title       *string `json:"title" validate:"" form:"title"`             // 标题
	Author      *string `json:"author" validate:"" form:"author"`           // 作者
	Source      *string `json:"source" validate:"" form:"source"`           // 来源
	Content     *string `json:"content" validate:"" form:"content"`         // 内容
	Summary     *string `json:"summary" validate:"" form:"summary"`         // 摘要
	Image       *string `json:"image" validate:"" form:"image"`             // 图片
	Images      *string `json:"images" validate:"" form:"images"`           // 图片集
	Download    *string `json:"download" validate:"" form:"download"`       // 文件下载
	Tags        *string `json:"tags" validate:"" form:"tags"`               // TAG
	Hits        *int32  `json:"hits" validate:"" form:"hits"`               // 点击次数
	Keywords    *string `json:"keywords" validate:"" form:"keywords"`       // 关键词
	Description *string `json:"description" validate:"" form:"description"` // 描述
	Template    *string `json:"template" validate:"" form:"template"`       // 模板
	Url         *string `json:"url" validate:"" form:"url"`                 // 跳转地址
	ViewAuth    *int32  `json:"view_auth" validate:"" form:"view_auth"`     // 阅读权限

}

type UpdateArticleResponse struct {
	ArticleResponse
}

// 删除Article article
type DeleteArticleRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除Article article
type GetArticleRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteArticleResponse struct {
	Response
	Data int `json:"data"`
}
