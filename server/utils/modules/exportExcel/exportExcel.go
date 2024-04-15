package exportExcel

import (
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ExportExcel(context *gin.Context, data [][]interface{}, fileName string) error {
	file := excelize.NewFile()
	// 创建一个工作表
	sheet := "sheet1"

	if err := file.SetSheetName(sheet, fileName); err != nil {
		return err
	}
	if err := file.SetColWidth(sheet, "A", "Z", 20); err != nil {
		return err
	}
	// 设置标题样式
	titleStyle, err := file.NewStyle(&excelize.Style{Font: &excelize.Font{Bold: true, VertAlign: "center"}, Fill: excelize.Fill{Color: []string{"#cccccc"}, Type: "pattern", Pattern: 1}})

	// 设置数据样式
	dataStyle, err := file.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 12}, Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center", WrapText: true}})

	// 将数据写入 Excel 文件
	for rowIndex, row := range data {
		for colIndex, cell := range row {
			cellIndex, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			err := file.SetCellValue(sheet, cellIndex, cell)
			if err != nil {
				return err
			}
			if rowIndex == 0 {
				_ = file.SetCellStyle(sheet, cellIndex, cellIndex, titleStyle)
			} else {
				_ = file.SetCellStyle(sheet, cellIndex, cellIndex, dataStyle)
			}
		}
	}

	// 将 Excel 文件内容写入 HTTP 响应
	context.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	context.Header("Content-Disposition", "attachment; filename="+fileName+".xlsx")
	context.Header("Content-Transfer-Encoding", "binary")
	err = file.Write(context.Writer)
	return err
}
