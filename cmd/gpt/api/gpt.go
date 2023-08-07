package main

import (
	"flag"
	"fmt"

	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/config"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/handler"
	"github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	_ "github.com/ptonlix/drinkGPT/cmd/gpt/api/internal/prompt"
)

var configFile = flag.String("f", "etc/gpt.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
