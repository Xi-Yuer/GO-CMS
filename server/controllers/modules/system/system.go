package systemControllerModules

import (
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

var SystemController = &systemController{}

type systemController struct{}

type MemInfo struct {
	Total       uint64  `json:"total"`       // 总内存 (bytes)
	Available   uint64  `json:"available"`   // 可用内存 (bytes)
	Used        uint64  `json:"used"`        // 已使用内存 (bytes)
	UsedPercent float64 `json:"usedPercent"` // 内存使用率 (%)
	Free        uint64  `json:"free"`        // 空闲内存 (bytes)
	Active      uint64  `json:"active"`      // 活跃内存 (bytes)
	Inactive    uint64  `json:"inactive"`    // 非活跃内存 (bytes)
	Wired       uint64  `json:"wired"`       // 已分配内存 (bytes)
	Laundry     uint64  `json:"laundry"`     // 被回收内存 (bytes)
}

var response []systemInfo

func init() {
	initSystemInfo()
}
func initSystemInfo() {
	go func() {
		ticker := time.Tick(time.Second)
		for {
			select {
			case <-ticker:
				// CPU使用率
				CPURate, _ := cpu.Percent(time.Second, false)
				// 内存信息
				MemInfo, _ := mem.VirtualMemory()
				info := systemInfo{
					CPUUsage: CPURate[0],
					MemUsage: MemInfo,
				}
				response = append(response, info)
				if len(response) > 20 {
					response = response[1:]
				}
			}
		}
	}()
}

type systemInfo struct {
	CPUUsage float64                `json:"cpuUsage"`
	MemUsage *mem.VirtualMemoryStat `json:"memUsage"`
}

// GetSystemInfo 获取系统运行信息
// @Summary 获取系统运行信息
// @Description 获取系统运行信息
// @Tags 系统管理
// @Accept json
// @Produce json
// @Router /system [get]
func (s *systemController) GetSystemInfo(context *gin.Context) {
	utils.Response.Success(context, response)
}
