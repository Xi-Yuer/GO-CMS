package TableNameRoute

import (
	"github.com/Xi-Yuer/cms/template/server/controller"
	"github.com/gin-gonic/gin"
)

func UseTableNameRoutes(r *gin.RouterGroup) {
	group := r.Group("/TableName")
	c := TableNameController.NewTableNameController()
	{
		group.GET("", c.GetTableNameListRepo)
		group.POST("", c.CreateTableNameRecordRepo)
		group.PUT("/:id", c.UpdateTableNameRecordRepo)
		group.DELETE("/:id", c.DeleteTableNameRecordRepo)
	}
}
