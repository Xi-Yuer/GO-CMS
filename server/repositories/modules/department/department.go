package departmentRepositoryModules

import (
	"database/sql"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"time"
)

var DepartmentRepository = &departmentRepository{}

type departmentRepository struct{}

func (d *departmentRepository) CreateDepartment(params *dto.CreateDepartmentRequest) error {
	var (
		query  string
		values []interface{}
	)
	id := utils.GenID()

	if params.ParentDepartment == "" {
		query = "INSERT INTO department (id,department_name,department_order) VALUES (?,?,?)"
		values = append(values, id, params.DepartmentName, params.DepartmentOrder)
	} else {
		query = "INSERT INTO department (id,department_name,parent_department,department_order) VALUES (?,?,?,?)"
		values = append(values, id, params.DepartmentName, params.ParentDepartment, params.DepartmentOrder)
	}

	if _, err := db.DB.Exec(query, values...); err != nil {
		return err
	}
	return nil
}

func (d *departmentRepository) DeleteDepartment(id string) error {
	query := "UPDATE department SET delete_time = ? WHERE id = ?"
	_, err := db.DB.Exec(query, time.Now(), id)
	return err
}

func (d *departmentRepository) GetDepartments() ([]*dto.DepartmentResponse, error) {
	query := "SELECT id,department_name,parent_department,department_order,create_time,update_time FROM department WHERE delete_time IS NULL"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)
	var departments []*dto.DepartmentResponse
	for rows.Next() {
		var department dto.DepartmentResponse
		err := rows.Scan(&department.ID, &department.DepartmentName, &department.ParentDepartment, &department.DepartmentOrder, &department.CreateTime, &department.UpdateTime)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}
	return departments, nil
}
