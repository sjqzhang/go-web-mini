package vo

import (
	"time"
)

type BranchResponse struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt   *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	Repo        *string    `json:"repo" form:"repo"`                               // repo
	BranchName  *string    `json:"branch_name" form:"branch_name"`                 // branch_name
	BranchType  *int32     `json:"branch_type" form:"branch_type"`                 // '''''branch_type'''''
	IsDev       *int32     `json:"is_dev" form:"is_dev"`                           // '''''is_dev'''''
	JiraKey     *string    `json:"jira_key" form:"jira_key"`                       // jira_key
	CommitId    *string    `json:"commit_id" form:"commit_id"`                     // commit_id
	CommitTitle *string    `json:"commit_title" form:"commit_title"`               // commit_title
	Committer   *string    `json:"committer" form:"committer"`                     // committer
	CommitTime  *int32     `json:"commit_time" form:"commit_time"`                 // '''''commit_time'''''
	Creator     *string    `json:"creator" form:"creator"`                         // creator
	SyncTime    *int32     `json:"sync_time" form:"sync_time"`                     // '''''sync_time'''''

}

// 查询Branch branch
type Branch struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`                   // created_at
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`                   // updated_at
	DeletedAt   *time.Time `json:"deleted_at" form:"deleted_at"`                   // deleted_at
	Repo        *string    `json:"repo" form:"repo"`                               // repo
	BranchName  *string    `json:"branch_name" form:"branch_name"`                 // branch_name
	BranchType  *int32     `json:"branch_type" form:"branch_type"`                 // '''''branch_type'''''
	IsDev       *int32     `json:"is_dev" form:"is_dev"`                           // '''''is_dev'''''
	JiraKey     *string    `json:"jira_key" form:"jira_key"`                       // jira_key
	CommitId    *string    `json:"commit_id" form:"commit_id"`                     // commit_id
	CommitTitle *string    `json:"commit_title" form:"commit_title"`               // commit_title
	Committer   *string    `json:"committer" form:"committer"`                     // committer
	CommitTime  *int32     `json:"commit_time" form:"commit_time"`                 // '''''commit_time'''''
	Creator     *string    `json:"creator" form:"creator"`                         // creator
	SyncTime    *int32     `json:"sync_time" form:"sync_time"`                     // '''''sync_time'''''

}

// 查询Branch branch
type ListBranchResponse struct {
	Total    int64                  `json:"total"`                    //总数
	List     []Branch               `json:"list"`                     //列表
	PageNum  int                    `json:"pageNum" form:"pageNum"`   //第几页
	PageSize int                    `json:"pageSize" form:"pageSize"` //每页多少条
	Extra    map[string]interface{} `json:"extra"`                    //扩展
}

// 查询Branch branch
type ListBranchRequest struct {
	Repo        *string `json:"repo"  form:"repo"`                 // repo
	BranchName  *string `json:"branch_name"  form:"branch_name"`   // branch_name
	BranchType  *int32  `json:"branch_type"  form:"branch_type"`   // '''''branch_type'''''
	IsDev       *int32  `json:"is_dev"  form:"is_dev"`             // '''''is_dev'''''
	JiraKey     *string `json:"jira_key"  form:"jira_key"`         // jira_key
	CommitId    *string `json:"commit_id"  form:"commit_id"`       // commit_id
	CommitTitle *string `json:"commit_title"  form:"commit_title"` // commit_title
	Committer   *string `json:"committer"  form:"committer"`       // committer
	CommitTime  *int32  `json:"commit_time"  form:"commit_time"`   // '''''commit_time'''''
	Creator     *string `json:"creator"  form:"creator"`           // creator
	SyncTime    *int32  `json:"sync_time"  form:"sync_time"`       // '''''sync_time'''''

	PageNum  *uint `json:"pageNum" form:"pageNum"`   //第几页
	PageSize *uint `json:"pageSize" form:"pageSize"` //每页多少条
}

type GetBranchResponse struct {
	BranchResponse
}

// 创建Branch branch
type CreateBranchRequest struct {
	Repo        *string `json:"repo" form:"repo"`                 // repo
	BranchName  *string `json:"branch_name" form:"branch_name"`   // branch_name
	BranchType  *int32  `json:"branch_type" form:"branch_type"`   // '''''branch_type'''''
	IsDev       *int32  `json:"is_dev" form:"is_dev"`             // '''''is_dev'''''
	JiraKey     *string `json:"jira_key" form:"jira_key"`         // jira_key
	CommitId    *string `json:"commit_id" form:"commit_id"`       // commit_id
	CommitTitle *string `json:"commit_title" form:"commit_title"` // commit_title
	Committer   *string `json:"committer" form:"committer"`       // committer
	CommitTime  *int32  `json:"commit_time" form:"commit_time"`   // '''''commit_time'''''
	Creator     *string `json:"creator" form:"creator"`           // creator
	SyncTime    *int32  `json:"sync_time" form:"sync_time"`       // '''''sync_time'''''

}

type CreateBranchResponse struct {
	BranchResponse
}

// 更新Branch branch
type UpdateBranchRequest struct {
	ID          *int    `json:""`
	Repo        *string `json:"repo" form:"repo"`                 // repo
	BranchName  *string `json:"branch_name" form:"branch_name"`   // branch_name
	BranchType  *int32  `json:"branch_type" form:"branch_type"`   // '''''branch_type'''''
	IsDev       *int32  `json:"is_dev" form:"is_dev"`             // '''''is_dev'''''
	JiraKey     *string `json:"jira_key" form:"jira_key"`         // jira_key
	CommitId    *string `json:"commit_id" form:"commit_id"`       // commit_id
	CommitTitle *string `json:"commit_title" form:"commit_title"` // commit_title
	Committer   *string `json:"committer" form:"committer"`       // committer
	CommitTime  *int32  `json:"commit_time" form:"commit_time"`   // '''''commit_time'''''
	Creator     *string `json:"creator" form:"creator"`           // creator
	SyncTime    *int32  `json:"sync_time" form:"sync_time"`       // '''''sync_time'''''

}

type UpdateBranchResponse struct {
	BranchResponse
}

// 删除Branch branch
type DeleteBranchRequest struct {
	Ids []int64 `json:"ids" uri:"ids" form:"ids"` //编号列表
}

// 删除Branch branch
type GetBranchRequest struct {
	ID int64 `json:"id" uri:"id" form:"id"` //编号
}

type DeleteBranchResponse struct {
	Response
	Data int `json:"data"`
}
