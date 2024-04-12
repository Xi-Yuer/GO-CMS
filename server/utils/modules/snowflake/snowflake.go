package snowflake

import (
	"fmt"
	logsModules "github.com/Xi-Yuer/cms/utils/modules/logs"
	"github.com/bwmarrin/snowflake"
	"time"
)

var Node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		logsModules.Log.Error(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	Node, err = snowflake.NewNode(machineID)
	if err != nil {
		logsModules.Log.Error(err)
		return
	}

	return
}

func init() {
	if err := Init("2024-01-01", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return
	}
}

// GenID 生成 64 位的 雪花 ID
func GenID() int64 {
	return Node.Generate().Int64()
}
