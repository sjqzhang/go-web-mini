package vo

import (
    "time"
)


// 查询BranchTab branch_tab
type BranchTab struct {
 
    ID *int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"` // id
     
    Repo *string `json:"repo" form:"repo"` // repo
     
    BranchName *string `json:"branch_name" form:"branch_name"` // branch_name
     
    BranchType *int32 `json:"branch_type" form:"branch_type"` // branch_type
     
    IsDev *int32 `json:"is_dev" form:"is_dev"` // is_dev
     
    JiraKey *string `json:"jira_key" form:"jira_key"` // jira_key
     
    CommitId *string `json:"commit_id" form:"commit_id"` // commit_id
     
    CommitTitle *string `json:"commit_title" form:"commit_title"` // commit_title
     
    Committer *string `json:"committer" form:"committer"` // committer
     
    CommitTime *int32 `json:"commit_time" form:"commit_time"` // commit_time
     
    Creator *string `json:"creator" form:"creator"` // creator
     
    SyncTime *int32 `json:"sync_time" form:"sync_time"` // sync_time
     
    DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"` // deleted_at
     
    CreatedAt *time.Time `json:"created_at" form:"created_at"` // created_at
     
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at"` // updated_at
     
}

// 查询BranchTab branch_tab
type PagerBranchTab struct {
	Total    int64                  `json:"total"`
	List     []BranchTab          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询BranchTab branch_tab
type ListBranchTabRequest struct {
    
     
    Repo *string `json:"repo"  form:"repo"` // repo
    
     
    BranchName *string `json:"branch_name"  form:"branch_name"` // branch_name
    
     
    BranchType *int32 `json:"branch_type"  form:"branch_type"` // branch_type
    
     
    IsDev *int32 `json:"is_dev"  form:"is_dev"` // is_dev
    
     
    JiraKey *string `json:"jira_key"  form:"jira_key"` // jira_key
    
     
    CommitId *string `json:"commit_id"  form:"commit_id"` // commit_id
    
     
    CommitTitle *string `json:"commit_title"  form:"commit_title"` // commit_title
    
     
    Committer *string `json:"committer"  form:"committer"` // committer
    
     
    CommitTime *int32 `json:"commit_time"  form:"commit_time"` // commit_time
    
     
    Creator *string `json:"creator"  form:"creator"` // creator
    
     
    SyncTime *int32 `json:"sync_time"  form:"sync_time"` // sync_time
    
     
     
     
     
     PageNum  *uint   `json:"pageNum" form:"pageNum"`
     PageSize *uint   `json:"pageSize" form:"pageSize"`
}


// 创建BranchTab branch_tab
type CreateBranchTabRequest struct {
    
     
    Repo *string `json:"repo" form:"repo"` // repo
    
     
    BranchName *string `json:"branch_name" form:"branch_name"` // branch_name
    
     
    BranchType *int32 `json:"branch_type" form:"branch_type"` // branch_type
    
     
    IsDev *int32 `json:"is_dev" form:"is_dev"` // is_dev
    
     
    JiraKey *string `json:"jira_key" form:"jira_key"` // jira_key
    
     
    CommitId *string `json:"commit_id" form:"commit_id"` // commit_id
    
     
    CommitTitle *string `json:"commit_title" form:"commit_title"` // commit_title
    
     
    Committer *string `json:"committer" form:"committer"` // committer
    
     
    CommitTime *int32 `json:"commit_time" form:"commit_time"` // commit_time
    
     
    Creator *string `json:"creator" form:"creator"` // creator
    
     
    SyncTime *int32 `json:"sync_time" form:"sync_time"` // sync_time
    
     
     
     
     
}


// 更新BranchTab branch_tab
type UpdateBranchTabRequest struct {
    ID      *int `json:""`
    
     
    Repo *string `json:"repo" form:"repo"` // repo
    
     
    BranchName *string `json:"branch_name" form:"branch_name"` // branch_name
    
     
    BranchType *int32 `json:"branch_type" form:"branch_type"` // branch_type
    
     
    IsDev *int32 `json:"is_dev" form:"is_dev"` // is_dev
    
     
    JiraKey *string `json:"jira_key" form:"jira_key"` // jira_key
    
     
    CommitId *string `json:"commit_id" form:"commit_id"` // commit_id
    
     
    CommitTitle *string `json:"commit_title" form:"commit_title"` // commit_title
    
     
    Committer *string `json:"committer" form:"committer"` // committer
    
     
    CommitTime *int32 `json:"commit_time" form:"commit_time"` // commit_time
    
     
    Creator *string `json:"creator" form:"creator"` // creator
    
     
    SyncTime *int32 `json:"sync_time" form:"sync_time"` // sync_time
    
     
     
     
     
}

// 删除BranchTab branch_tab
type DeleteBranchTabRequest struct {
    Ids      []int64 `json:"ids" uri:"ids" form:"ids"`
}

// 删除BranchTab branch_tab
type GetBranchTabRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"`
}

