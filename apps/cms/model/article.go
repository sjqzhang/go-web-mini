package model

import (
	"time"
)

var _ = time.Now()

// Article article
type Article struct {
	Model
	Sort        int32  `gorm:"sort;type:mediumint(8);comment:排序" validate:"" json:"sort"`                    // 排序
	Status      int32  `gorm:"status;type:tinyint(1);comment:状态" validate:"" json:"status"`                  // 状态
	CateId      int32  `gorm:"cate_id;type:tinyint(4) unsigned;comment:栏目" validate:"" json:"cate_id"`       // 栏目
	Title       string `gorm:"title;type:varchar(255);comment:标题" validate:"" json:"title"`                  // 标题
	Author      string `gorm:"author;type:varchar(255);comment:作者" validate:"" json:"author"`                // 作者
	Source      string `gorm:"source;type:varchar(255);comment:来源" validate:"" json:"source"`                // 来源
	Content     string `gorm:"content;type:text;comment:内容" validate:"" json:"content"`                      // 内容
	Summary     string `gorm:"summary;type:text;comment:摘要" validate:"" json:"summary"`                      // 摘要
	Image       string `gorm:"image;type:varchar(80);comment:图片" validate:"" json:"image"`                   // 图片
	Images      string `gorm:"images;type:text;comment:图片集" validate:"" json:"images"`                       // 图片集
	Download    string `gorm:"download;type:varchar(80);comment:文件下载" validate:"" json:"download"`           // 文件下载
	Tags        string `gorm:"tags;type:varchar(255);comment:TAG" validate:"" json:"tags"`                   // TAG
	Hits        int32  `gorm:"hits;type:int(10) unsigned;comment:点击次数" validate:"" json:"hits"`              // 点击次数
	Keywords    string `gorm:"keywords;type:varchar(255);comment:关键词" validate:"" json:"keywords"`           // 关键词
	Description string `gorm:"description;type:varchar(255);comment:描述" validate:"" json:"description"`      // 描述
	Template    string `gorm:"template;type:varchar(30);comment:模板" validate:"" json:"template"`             // 模板
	Url         string `gorm:"url;type:varchar(255);comment:跳转地址" validate:"" json:"url"`                    // 跳转地址
	ViewAuth    int32  `gorm:"view_auth;type:tinyint(4) unsigned;comment:阅读权限" validate:"" json:"view_auth"` // 阅读权限

}

// Article article
type ArticleQuery struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt   *time.Time `json:"created_at"`                           // created_at
	UpdatedAt   *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt   *time.Time `json:"deleted_at"`                           // deleted_at
	Sort        *int32     `json:"sort"`                                 // 排序
	Status      *int32     `json:"status"`                               // 状态
	CateId      *int32     `json:"cate_id"`                              // 栏目
	Title       *string    `json:"title"`                                // 标题
	Author      *string    `json:"author"`                               // 作者
	Source      *string    `json:"source"`                               // 来源
	Content     *string    `json:"content"`                              // 内容
	Summary     *string    `json:"summary"`                              // 摘要
	Image       *string    `json:"image"`                                // 图片
	Images      *string    `json:"images"`                               // 图片集
	Download    *string    `json:"download"`                             // 文件下载
	Tags        *string    `json:"tags"`                                 // TAG
	Hits        *int32     `json:"hits"`                                 // 点击次数
	Keywords    *string    `json:"keywords"`                             // 关键词
	Description *string    `json:"description"`                          // 描述
	Template    *string    `json:"template"`                             // 模板
	Url         *string    `json:"url"`                                  // 跳转地址
	ViewAuth    *int32     `json:"view_auth"`                            // 阅读权限
	PageNum     int        `json:"-" form:"pageNum"`
	PageSize    int        `json:"-" form:"pageSize"`
}

func (t Article) TableName() string {
	return "article"
}
