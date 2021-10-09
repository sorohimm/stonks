package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"stonks/internal/config"
	"stonks/internal/infrastructure"
)

var (
	ctx context.Context
	cfg *config.Config
	log *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error loading logger: %s", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config init error: %s", err)
	}
	log.Infof("Config loaded:\n%+v", cfg)


	ctx = context.Background()
}

func main() {
	injector, _ := infrastructure.Injector(log, cfg)
	newsController := injector.InjectNewsController()

	router := gin.Default()


	v1 := router.Group("/stonks/v1")
	{
		v1.GET("/news", newsController.GetNews)
	}

	log.Fatal(router.Run())
}
