package main

import (
	"flag"
	"fmt"
	"net/http"

	"micro/api/test/internal/config"
	"micro/api/test/internal/handler"
	"micro/api/test/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		// 自定义处理返回
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
