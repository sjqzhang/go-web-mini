package model

import (
	"time"
)

// Branch branch
type Branch struct {
	Model
	Repo        string `gorm:"repo;type:longtext;comment:'repo'" json:"repo"`                               // repo
	BranchName  string `gorm:"branch_name;type:longtext;comment:'branch_name'" json:"branch_name"`          // branch_name
	BranchType  int32  `gorm:"branch_type;type:int(11);comment:''''''branch_type''''''" json:"branch_type"` // '''''branch_type'''''
	IsDev       int32  `gorm:"is_dev;type:int(11);comment:''''''is_dev''''''" json:"is_dev"`                // '''''is_dev'''''
	JiraKey     string `gorm:"jira_key;type:longtext;comment:'jira_key'" json:"jira_key"`                   // jira_key
	CommitId    string `gorm:"commit_id;type:longtext;comment:'commit_id'" json:"commit_id"`                // commit_id
	CommitTitle string `gorm:"commit_title;type:longtext;comment:'commit_title'" json:"commit_title"`       // commit_title
	Committer   string `gorm:"committer;type:longtext;comment:'committer'" json:"committer"`                // committer
	CommitTime  int32  `gorm:"commit_time;type:int(11);comment:''''''commit_time''''''" json:"commit_time"` // '''''commit_time'''''
	Creator     string `gorm:"creator;type:longtext;comment:'creator'" json:"creator"`                      // creator
	SyncTime    int32  `gorm:"sync_time;type:int(11);comment:''''''sync_time''''''" json:"sync_time"`       // '''''sync_time'''''

}

// Branch branch
type BranchQuery struct {
	ID          *int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // id
	CreatedAt   *time.Time `json:"created_at"`                           // created_at
	UpdatedAt   *time.Time `json:"updated_at"`                           // updated_at
	DeletedAt   *time.Time `json:"deleted_at"`                           // deleted_at
	Repo        *string    `json:"repo"`                                 // repo
	BranchName  *string    `json:"branch_name"`                          // branch_name
	BranchType  *int32     `json:"branch_type"`                          // '''''branch_type'''''
	IsDev       *int32     `json:"is_dev"`                               // '''''is_dev'''''
	JiraKey     *string    `json:"jira_key"`                             // jira_key
	CommitId    *string    `json:"commit_id"`                            // commit_id
	CommitTitle *string    `json:"commit_title"`                         // commit_title
	Committer   *string    `json:"committer"`                            // committer
	CommitTime  *int32     `json:"commit_time"`                          // '''''commit_time'''''
	Creator     *string    `json:"creator"`                              // creator
	SyncTime    *int32     `json:"sync_time"`                            // '''''sync_time'''''
	PageNum     int        `json:"-" form:"pageNum"`
	PageSize    int        `json:"-" form:"pageSize"`
}

func (t Branch) TableName() string {
	return "branch_tab"
}
