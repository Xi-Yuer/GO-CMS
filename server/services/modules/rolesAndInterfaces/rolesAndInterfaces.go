package rolesAndInterfacesServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
)

var RolesAndInterfacesService = &rolesAndInterfacesService{}

type rolesAndInterfacesService struct{}

// GetInterfacesByRoleID 根据角色ID获取接口
func (r *rolesAndInterfacesService) GetInterfacesByRoleID(roleId string) ([]*dto.GetInterfaceResponse, error) {
	interIDs, err := repositories.RolesAndInterfacesRepository.GetRecordByRoleID(roleId)
	if err != nil {
		return nil, err
	}
	var interfaceDic []*dto.GetInterfaceResponse
	for _, id := range interIDs {
		inter, e := repositories.InterfaceRepository.GetInterfaceByID(id)
		if !e {
			continue
		}
		interfaceDic = append(interfaceDic, inter)
	}
	return interfaceDic, nil
}
