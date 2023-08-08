package vo

import (
    "time"
)





// 查询Branch branch
type Branch struct {
 
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

// 查询Branch branch
type PagerBranch struct {
	Total    int64                  `json:"total"`
	List     []Branch          `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}



// 查询Branch branch
type ListBranchRequest struct {
    
     
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


// 创建Branch branch
type CreateBranchRequest struct {
    
     
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


// 更新Branch branch
type UpdateBranchRequest struct {
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

// 删除Branch branch
type DeleteBranchRequest struct {
    Ids      []int64 `json:"ids" uri:"ids" form:"ids"`
}

// 删除Branch branch
type GetBranchRequest struct {
    ID      int64 `json:"id" uri:"id" form:"id"`
}



//以下结构体只用于生成文档
type ListBranchResponse  struct {
    Response
    Data PagerBranch  `json:"data"`
}

type GetBranchResponse  struct {
    Response
    Data Branch  `json:"data"`
}

type CreateBranchResponse  struct {
    Response
    Data Branch  `json:"data"`
}

type UpdateBranchResponse  struct {
    Response
    Data Branch  `json:"data"`
}

type DeleteBranchResponse  struct {
    Response
    Data int `json:"data"`
}

