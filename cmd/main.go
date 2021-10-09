package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"regexp"
	"stonks/internal/config"
	"stonks/internal/infrastructure"
)

var (
	ctx context.Context
	cfg *config.Config
	log *zap.SugaredLogger
)

const projectDirName = "stonkks"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/configs/config.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error loading logger: %s", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()

	loadEnv()

	cfg = config.New()
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

	router.Run()
}
