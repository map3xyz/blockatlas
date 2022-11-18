package main

import (
	"context"

	golibsGin "github.com/trustwallet/golibs/network/gin"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/config"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
)

const (
	defaultPort       = "8430"
	defaultConfigPath = "../../config.yml"
)

var (
	ctx            context.Context
	cancel         context.CancelFunc
	port, confPath string
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	ctx, cancel = context.WithCancel(context.Background())

	internal.InitConfig(confPath)

	engine = internal.InitEngine(config.Default.Gin.Mode)
}

func main() {
	api.SetupMetrics(engine)
	api.SetupKeychainAPI(engine)

	golibsGin.SetupGracefulShutdown(ctx, port, engine)
	cancel()
}
