package main

import (
	"context"

	golibsGin "github.com/trustwallet/golibs/network/gin"

	"github.com/trustwallet/golibs/network/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	if err := middleware.SetupSentry(config.Default.Sentry.DSN); err != nil {
		log.Error(err)
	}

	engine = internal.InitEngine(config.Default.Gin.Mode)
}

func main() {
	api.SetupMetrics(engine)
	setupAPI(engine)

	golibsGin.SetupGracefulShutdown(ctx, port, engine)
	cancel()
}

func setupAPI(router gin.IRouter) {
	log.Info("Starting keychain service")

	// getAddress(user, customer-apikey, network, asset, addressType (user, memo))
	// -> { address, memo }
	// apikey -> keychainID
	// customer-apikey -> wallet
	// network/asset -> address
	// generate new memo
	// call store api -> (watchAddress/userAddress, memo)
	router.POST("/v1/address")

	// getEvents(customer-apikey)
	// apikey -> keychainID
	// customer-apikey -> wallet
	router.GET("/v1/events")
}
