package interfaceServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
)

var InterfaceService = &interfaceService{}

type interfaceService struct{}

func (i *interfaceService) CreateInterface(params *dto.CreateInterfaceRequest) error {
	return repositories.InterfaceRepository.CreateInterface(params)
}

func (i *interfaceService) GetInterfaceByPageID(id string) []*dto.GetInterfaceResponse {
	return repositories.InterfaceRepository.GetInterfaceByPageID(id)
}

func (i *interfaceService) UpdateInterfaceByID(id string, params *dto.UpdateInterfaceRequest) error {
	if _, exist := repositories.InterfaceRepository.GetInterfaceByID(id); !exist {
		return errors.New("接口不存在")
	}
	return repositories.InterfaceRepository.UpdateInterfaceByID(id, params)
}

func (i *interfaceService) DeleteInterfaceByID(id string) error {
	if _, exist := repositories.InterfaceRepository.GetInterfaceByID(id); !exist {
		return errors.New("接口不存在")
	}
	return repositories.InterfaceRepository.DeleteInterfaceByID(id)
}
