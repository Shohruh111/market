package main

import (
	"fmt"
	"market/api"
	"market/config"
	"market/pkg/logger"
	"market/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("market", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	pgconn, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())
	api.NewApi(r, &cfg, pgconn, log)

	fmt.Println("Listening server...", cfg.ServerHost+cfg.HTTPPort)
	err = r.Run(cfg.ServerHost + cfg.HTTPPort)
	if err != nil {
		panic(err)
	}

}
