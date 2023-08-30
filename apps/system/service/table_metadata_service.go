package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go-web-mini/apps/system/repository"

	"go-web-mini/apps/system/model"
	"go-web-mini/apps/system/vo"
)

/*
// TableMetadataQueryPage table_metadata分页查询
func TableMetadataQueryPage(param dto.TableMetadataPageDTO) []vo.TableMetadataVO{

	return []vo.TableMetadataVO{}
}
*/

type TableMetadataService struct {
	tableMetadataRepository repository.TableMetadataRepository
}

func (s *TableMetadataService) List(ctx context.Context, req *vo.ListTableMetadataRequest) (*vo.ListTableMetadataResponse, error) {
	var query model.TableMetadataQuery
	err := copier.Copy(&query, req)
	if err != nil {
		return nil, err
	}
	var resp vo.ListTableMetadataResponse
	objs, err := s.tableMetadataRepository.List(ctx, &query)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, objs)
	return &resp, err
}

func (s *TableMetadataService) GetById(ctx context.Context, req *vo.GetTableMetadataRequest) (*vo.TableMetadataResponse, error) {
	obj, err := s.tableMetadataRepository.GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp vo.TableMetadataResponse
	err = copier.Copy(&resp, obj)
	return &resp, err
}

func (s *TableMetadataService) Create(ctx *gin.Context, req *vo.CreateTableMetadataRequest) (*vo.CreateTableMetadataResponse, error) {
	var obj model.TableMetadata
	err := copier.Copy(&obj, req)
	if err != nil {
		return nil, err
	}
	var resp vo.CreateTableMetadataResponse
	_, err = s.tableMetadataRepository.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *TableMetadataService) Update(ctx *gin.Context, req *vo.UpdateTableMetadataRequest) (*vo.UpdateTableMetadataResponse, error) {
	//var obj model.TableMetadata
	obj, err := s.tableMetadataRepository.GetById(ctx, int64(*req.ID))
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&obj, req)
	var resp vo.UpdateTableMetadataResponse
	_, err = s.tableMetadataRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, &obj)
	return &resp, err
}

func (s *TableMetadataService) Delete(ctx *gin.Context, req *vo.DeleteTableMetadataRequest) (int64, error) {
	return s.tableMetadataRepository.Delete(ctx, req.Ids)
}
