package model

import (
	"time"
)

// Branch branch
type Branch struct {
	Model

	Repo string `json:"repo"` // repo

	BranchName string `json:"branch_name"` // branch_name

	BranchType int32 `json:"branch_type"` // branch_type

	IsDev int32 `json:"is_dev"` // is_dev

	JiraKey string `json:"jira_key"` // jira_key

	CommitId string `json:"commit_id"` // commit_id

	CommitTitle string `json:"commit_title"` // commit_title

	Committer string `json:"committer"` // committer

	CommitTime int32 `json:"commit_time"` // commit_time

	Creator string `json:"creator"` // creator

	SyncTime int32 `json:"sync_time"` // sync_time

}

// Branch branch
type BranchQuery struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	Repo        *string    `json:"repo"`                                 // repo
	BranchName  *string    `json:"branch_name"`                          // branch_name
	BranchType  *int32     `json:"branch_type"`                          // branch_type
	IsDev       *int32     `json:"is_dev"`                               // is_dev
	JiraKey     *string    `json:"jira_key"`                             // jira_key
	CommitId    *string    `json:"commit_id"`                            // commit_id
	CommitTitle *string    `json:"commit_title"`                         // commit_title
	Committer   *string    `json:"committer"`                            // committer
	CommitTime  *int32     `json:"commit_time"`                          // commit_time
	Creator     *string    `json:"creator"`                              // creator
	SyncTime    *int32     `json:"sync_time"`                            // sync_time
	DeletedAt   *time.Time `json:"deleted_at"`                           // deleted_at
	CreatedAt   *time.Time `json:"created_at"`                           // created_at
	UpdatedAt   *time.Time `json:"updated_at"`                           // updated_at

	PageNum  int `json:"-" form:"pageNum"`
	PageSize int `json:"-" form:"pageSize"`
}

func (t Branch) TableName() string {
	return "branch_tab"
}
