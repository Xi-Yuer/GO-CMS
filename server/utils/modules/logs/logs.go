package logsModules

import (
	"fmt"
	"os"
	"time"
)

var Log = &log{}

type log struct {
}

func baseInfo() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s]", t)
}

func (l *log) Info(v ...interface{}) {
	fmt.Printf("【信息】%s %s\n", baseInfo(), v)
}

func (l *log) Warn(v ...interface{}) {
	fmt.Printf("【警告】%s %s\n", baseInfo(), v)
}
func (l *log) Error(v ...interface{}) {
	fmt.Printf("【错误】%s %s\n", baseInfo(), v)
}
func (l *log) Panic(v ...interface{}) {
	fmt.Printf("【系统错误】%s %s\n", baseInfo(), v)
	os.Exit(0)
}
