package interfaceRepositoryModules

import (
	"encoding/json"
	"errors"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strings"
)

var InterfaceRepository = &interfaceRepository{}

type interfaceRepository struct{}

func (i *interfaceRepository) CreateInterface(params *dto.CreateInterfaceRequest) error {
	query := `INSERT INTO interfaces (interface_id, interface_name, interface_method, interface_path, interface_page_id, interface_desc,interface_dic) VALUES (?, ?, ?, ?, ?, ?, ?)`
	id := utils.GenID()
	_, err := db.DB.Exec(query, id, params.InterfaceName, params.InterfaceMethod, params.InterfacePath, params.InterfacePageID, params.InterfaceDesc, params.InterfaceDic)
	return err
}

func (i *interfaceRepository) GetInterfaceByPageID(id string) []*dto.GetInterfaceResponse {
	query := `SELECT interface_id, interface_name, interface_method, interface_path, interface_page_id, interface_desc,interface_dic, create_time,update_time FROM interfaces WHERE interface_page_id = ?`
	rows, _ := db.DB.Query(query, id)
	var interfaces []*dto.GetInterfaceResponse
	for rows.Next() {
		var interfaceInfo dto.GetInterfaceResponse
		err := rows.Scan(&interfaceInfo.ID, &interfaceInfo.InterfaceName, &interfaceInfo.InterfaceMethod, &interfaceInfo.InterfacePath, &interfaceInfo.InterfacePageID, &interfaceInfo.InterfaceDesc, &interfaceInfo.InterfaceDic, &interfaceInfo.CreateTime, &interfaceInfo.UpdateTime)
		if err != nil {
			return nil
		}
		interfaces = append(interfaces, &interfaceInfo)
	}
	return interfaces
}

func (i *interfaceRepository) GetAllInterface() ([]*dto.AllInterfaceResponse, error) {
	query := `
			SELECT
				JSON_OBJECT(
						'key', p.page_name,
						'children', JSON_ARRAYAGG(
								JSON_OBJECT(
										'ID', i.interface_id,
										'interfaceName', i.interface_name,
										'interfaceDic', i.interface_dic,
										'interfaceMethod', i.interface_method,
										'interfacePageId', i.interface_page_id,
										'interfacePath', i.interface_path,
										'createTime', i.create_time,
										'updateTime', i.update_time
								)
									  )
				) AS page_interfaces
			FROM
				pages p
					JOIN
				interfaces i ON p.page_id = i.interface_page_id
			GROUP BY
				p.page_name;
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var interfaceResponses []*dto.AllInterfaceResponse
	for rows.Next() {
		var interfaceResponse string
		var interfaceUnmarshalResponse *dto.AllInterfaceResponse
		if err := rows.Scan(&interfaceResponse); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(interfaceResponse), &interfaceUnmarshalResponse); err != nil {
			return nil, err
		}
		interfaceResponses = append(interfaceResponses, interfaceUnmarshalResponse)
	}
	return interfaceResponses, nil
}

func (i *interfaceRepository) GetInterfaceByID(id string) (inter *dto.GetInterfaceResponse, exist bool) {
	query := `SELECT interface_id, interface_name, interface_method, interface_path, interface_page_id, interface_desc,interface_dic, can_edit, create_time,update_time FROM interfaces WHERE interface_id = ?`
	row := db.DB.QueryRow(query, id)
	var interfaceInfo dto.GetInterfaceResponse
	err := row.Scan(&interfaceInfo.ID, &interfaceInfo.InterfaceName, &interfaceInfo.InterfaceMethod, &interfaceInfo.InterfacePath, &interfaceInfo.InterfacePageID, &interfaceInfo.InterfaceDesc, &interfaceInfo.InterfaceDic, &interfaceInfo.CanEdit, &interfaceInfo.CreateTime, &interfaceInfo.UpdateTime)
	if err != nil {
		return nil, false
	}
	return &interfaceInfo, true
}

func (i *interfaceRepository) UpdateInterfaceByID(id string, params *dto.UpdateInterfaceRequest) error {
	query := "UPDATE interfaces SET interface_name = ?, interface_method = ?, interface_path = ?, interface_page_id = ?, interface_desc = ?,interface_dic = ? WHERE interface_id = ?"
	_, err := db.DB.Exec(query, params.InterfaceName, params.InterfaceMethod, params.InterfacePath, params.InterfacePageID, params.InterfaceDic, params.InterfaceDesc, id)
	return err
}

func (i *interfaceRepository) DeleteInterfaceByID(id string) error {
	query := "DELETE FROM interfaces WHERE interface_id = ?"
	_, err := db.DB.Exec(query, id)
	return err
}

func (i *interfaceRepository) CheckInterfacesExistence(interfaceID []string) error {
	// 构建 IN 子句
	var placeholders []string
	var args []interface{}
	for _, id := range interfaceID {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}
	query := "SELECT COUNT(*) FROM interfaces WHERE interface_id IN (" + strings.Join(placeholders, ",") + ") "
	row := db.DB.QueryRow(query, args...)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	// 判断是否所有接口都存在
	if count != len(interfaceID) {
		return errors.New("资源不存在")
	}

	return nil
}
