package TableNameController

import (
	"github.com/Xi-Yuer/cms/template/server/dto"
	"github.com/Xi-Yuer/cms/template/server/service"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	GetTableNameListRepo(context *gin.Context)
	CreateTableNameRecordRepo(context *gin.Context)
	UpdateTableNameRecordRepo(context *gin.Context)
	DeleteTableNameRecordRepo(context *gin.Context)
}

func NewTableNameController() ControllerInterface {
	return &TableNameController{
		service: TableNameService.NewTableNameService(),
	}
}

type TableNameController struct {
	service TableNameService.ServiceInterface
}

func (t *TableNameController) GetTableNameListRepo(context *gin.Context) {
	var params TableNameDto.TableNameFindRequestDTO
	if err := context.ShouldBind(&params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	service, err := t.service.GetTableNameListService(&params)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"data": service})
}

func (t *TableNameController) CreateTableNameRecordRepo(context *gin.Context) {
	var params TableNameDto.TableNameCreateRequestDTO
	if err := context.ShouldBind(&params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := t.service.CreateTableNameRecordService(&params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"data": "success"})
}

func (t *TableNameController) UpdateTableNameRecordRepo(context *gin.Context) {
	var params TableNameDto.TableNameUpdateRequestDTO
	id := context.Param("id")
	if id == "" {
		context.JSON(400, gin.H{"error": "id is null"})
		return
	}
	if err := context.ShouldBind(&params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := t.service.UpdateTableNameListService(id, &params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"data": "success"})
}

func (t *TableNameController) DeleteTableNameRecordRepo(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(400, gin.H{"error": "id is null"})
		return
	}
	if err := t.service.DeleteTableNameRecordService(id); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	}
	context.JSON(200, gin.H{"data": "success"})
}
