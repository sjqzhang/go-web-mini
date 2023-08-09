package service

import (

    "context"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/copier"

    "go-web-mini/model"
    "go-web-mini/repository"
    "go-web-mini/vo"
)

/*
// BranchQueryPage branch分页查询
func BranchQueryPage(param dto.BranchPageDTO) []vo.BranchVO{

	return []vo.BranchVO{}
}
*/

type BranchService struct {
	branchRepository repository.BranchRepository
}

func (s *BranchService) List(ctx context.Context, req *vo.ListBranchRequest) (*vo.ListBranchResponse, error) {
	var query model.BranchQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
    var resp vo.ListBranchResponse
    objs,err:= s.branchRepository.List(ctx, &query)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, objs)
    return &resp, err
}

func (s *BranchService) GetById(ctx context.Context, req *vo.GetBranchRequest) (*vo.BranchResponse, error) {
    obj,err:= s.branchRepository.GetById(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    var resp vo.BranchResponse
    err = copier.Copy(&resp, obj)
    return &resp, err
}

func (s *BranchService) Create(ctx *gin.Context, req *vo.CreateBranchRequest) (*vo.CreateBranchResponse, error) {
	var obj model.Branch
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateBranchResponse
	_,err= s.branchRepository.Create(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return &resp, err
}


func (s *BranchService) Update(ctx *gin.Context, req *vo.UpdateBranchRequest) (*vo.UpdateBranchResponse, error) {
	var obj model.Branch
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.UpdateBranchResponse
	_,err= s.branchRepository.Update(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return &resp, err
}


func (s *BranchService) Delete(ctx *gin.Context, req *vo.DeleteBranchRequest) (int64, error) {
	return s.branchRepository.Delete(ctx, req.Ids)
}


