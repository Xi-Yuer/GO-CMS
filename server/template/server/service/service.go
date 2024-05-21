package TableNameService

import (
	"github.com/Xi-Yuer/cms/template/server/dto"
	"github.com/Xi-Yuer/cms/template/server/repository"
)

type ServiceInterface interface {
	GetTableNameListService(params *TableNameDto.TableNameFindRequestDTO) ([]*TableNameDto.TableNameFindRequestDTO, error)
	CreateTableNameRecordService(params *TableNameDto.TableNameCreateRequestDTO) error
	UpdateTableNameListService(id string, params *TableNameDto.TableNameUpdateRequestDTO) error
	DeleteTableNameRecordService(id string) error
}
type Service struct {
	repo TableNameRepository.RepositoryInterface
}

func NewTableNameService() *Service {
	return &Service{
		repo: TableNameRepository.NewTableNameRepository(),
	}
}

func (s Service) GetTableNameListService(params *TableNameDto.TableNameFindRequestDTO) ([]*TableNameDto.TableNameFindRequestDTO, error) {
	return s.repo.GetTableNameListRepo(params)
}

func (s Service) CreateTableNameRecordService(params *TableNameDto.TableNameCreateRequestDTO) error {
	return s.repo.CreateTableNameRecordRepo(params)
}

func (s Service) UpdateTableNameListService(id string, params *TableNameDto.TableNameUpdateRequestDTO) error {
	return s.repo.UpdateTableNameRecordRepo(id, params)
}

func (s Service) DeleteTableNameRecordService(id string) error {
	return s.repo.DeleteTableNameRecordRepo(id)
}
