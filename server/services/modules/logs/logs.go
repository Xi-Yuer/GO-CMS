package logsServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
)

var LogsService = &logsService{}

type logsService struct{}

func (receiver *logsService) GetLogRecords(params *dto.Page) (*dto.HasTotalResponseData, error) {
	return repositories.LogsRepository.GetLogRecords(params)
}
