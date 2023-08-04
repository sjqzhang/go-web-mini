package model

import (
	"time"
)

type Model struct {
	ID        int64      `gorm:"primarykey" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

//type queryModel struct {
//	PageNum  uint                     `json:"pageNum" form:"pageNum"`
//	PageSize uint                     `json:"pageSize" form:"pageSize"`
//	QueryMap []map[string]interface{} `json:"queryMap" form:"queryMap"`
//}
//
//
//func QueryModel(queryObject interface{}) *queryModel {
//	json.Marshal()
//	return &queryModel{
//		QueryMap: []map[string]interface{}{},
//	}
//}
