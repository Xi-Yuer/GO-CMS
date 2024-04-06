package departmentServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
)

var DepartmentService = &departmentService{}

type departmentService struct{}

func (d *departmentService) CreateDepartment(params *dto.CreateDepartmentRequest) error {
	return repositories.DepartmentRepository.CreateDepartment(params)
}

func (d *departmentService) DeleteDepartment(id string) error {
	return repositories.DepartmentRepository.DeleteDepartment(id)
}

func (d *departmentService) GetDepartments() ([]*dto.DepartmentResponse, error) {
	department, err := repositories.DepartmentRepository.GetDepartments()
	if err != nil {
		return nil, err
	}
	buildDepartment := utils.BuildDepartment(department)
	return buildDepartment, nil
}
