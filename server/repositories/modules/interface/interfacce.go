package interfaceRepositoryModules

import (
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
)

var InterfaceRepository = &interfaceRepository{}

type interfaceRepository struct{}

func (i *interfaceRepository) CreateInterface(params *dto.CreateInterfaceRequest) error {
	query := `INSERT INTO interfaces (interface_id, interface_name, interface_method, interface_path, interface_page_id, interface_desc) VALUES (?, ?, ?, ?, ?, ?)`
	id := utils.GenID()
	_, err := db.DB.Exec(query, id, params.InterfaceName, params.InterfaceMethod, params.InterfacePath, params.InterfacePageID, params.InterfaceDesc)
	return err
}

func (i *interfaceRepository) GetInterfaceByPageID(id string) []*dto.GetInterfaceResponse {
	query := `SELECT interface_id, interface_name, interface_method, interface_path, interface_page_id, interface_desc,create_time,update_time FROM interfaces WHERE interface_page_id = ?`
	rows, _ := db.DB.Query(query, id)
	var interfaces []*dto.GetInterfaceResponse
	for rows.Next() {
		var interfaceInfo dto.GetInterfaceResponse
		err := rows.Scan(&interfaceInfo.ID, &interfaceInfo.InterfaceName, &interfaceInfo.InterfaceMethod, &interfaceInfo.InterfacePath, &interfaceInfo.InterfacePageID, &interfaceInfo.InterfaceDesc, &interfaceInfo.CreateTime, &interfaceInfo.UpdateTime)
		if err != nil {
			return nil
		}
		interfaces = append(interfaces, &interfaceInfo)
	}
	return interfaces
}

func (i *interfaceRepository) GetInterfaceByID(id string) (inter *dto.GetInterfaceResponse, exist bool) {
	query := `SELECT interface_id, interface_name, interface_method, interface_path, interface_page_id, interface_desc,create_time,update_time FROM interfaces WHERE interface_id = ?`
	row := db.DB.QueryRow(query, id)
	var interfaceInfo dto.GetInterfaceResponse
	err := row.Scan(&interfaceInfo.ID, &interfaceInfo.InterfaceName, &interfaceInfo.InterfaceMethod, &interfaceInfo.InterfacePath, &interfaceInfo.InterfacePageID, &interfaceInfo.InterfaceDesc, &interfaceInfo.CreateTime, &interfaceInfo.UpdateTime)
	if err != nil {
		return nil, false
	}
	return &interfaceInfo, true
}

func (i *interfaceRepository) UpdateInterfaceByID(id string, params *dto.UpdateInterfaceRequest) error {
	query := "UPDATE interfaces SET interface_name = ?, interface_method = ?, interface_path = ?, interface_page_id = ?, interface_desc = ? WHERE interface_id = ?"
	_, err := db.DB.Exec(query, params.InterfaceName, params.InterfaceMethod, params.InterfacePath, params.InterfacePageID, params.InterfaceDesc, id)
	return err
}

func (i *interfaceRepository) DeleteInterfaceByID(id string) error {
	query := "DELETE FROM interfaces WHERE interface_id = ?"
	_, err := db.DB.Exec(query, id)
	return err
}
