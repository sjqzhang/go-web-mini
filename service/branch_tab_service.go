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
// BranchTabQueryPage branch_tab分页查询
func BranchTabQueryPage(param dto.BranchTabPageDTO) []vo.BranchTabVO{

	return []vo.BranchTabVO{}
}
*/

type BranchTabService struct {
	branchTabRepository repository.BranchTabRepository
}

func (s *BranchTabService) List(ctx context.Context, req *vo.ListBranchTabRequest) (*vo.PagerBranchTab, error) {
	var query model.BranchTabQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
    var resp vo.PagerBranchTab
    objs,err:= s.branchTabRepository.List(ctx, &query)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, objs)
    return &resp, err
}

func (s *BranchTabService) GetById(ctx context.Context, req *vo.GetBranchTabRequest) (*vo.BranchTab, error) {
    obj,err:= s.branchTabRepository.GetById(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    var resp vo.BranchTab
    err = copier.Copy(&resp, obj)
    return &resp, err
}

func (s *BranchTabService) Create(ctx *gin.Context, req *vo.CreateBranchTabRequest) (*vo.BranchTab, error) {
	var obj model.BranchTab
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.BranchTab
	_,err= s.branchTabRepository.Create(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return &resp, err
}


func (s *BranchTabService) Update(ctx *gin.Context, req *vo.UpdateBranchTabRequest) (*vo.BranchTab, error) {
	var obj model.BranchTab
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.BranchTab
	_,err= s.branchTabRepository.Update(ctx, &obj)
	if err != nil {
        return nil, err
    }
    err = copier.Copy(&resp, &obj)
    return &resp, err
}


func (s *BranchTabService) Delete(ctx *gin.Context, req *vo.DeleteBranchTabRequest) (int64, error) {
	return s.branchTabRepository.Delete(ctx, req.Ids)
}


