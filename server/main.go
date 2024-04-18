package main

import (
	"context"
	"github.com/Xi-Yuer/cms/bootStrap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	bootStrap.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
