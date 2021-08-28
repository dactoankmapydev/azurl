package main

import (
	"azurl/db"
	"azurl/handler"
	"azurl/log"
	"azurl/repository/repo_impl"
	"azurl/router"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load("pro.env"); err != nil {
		return
	}
}

func main() {
	// write log
	log, _ := log.WriteLog()

	// redis details
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// connect redis
	client := &db.Redis{
		Host:   redisHost,
		Port:   redisPort,
		Logger: log,
	}
	client.NewRedis()

	e := echo.New()

	shortHandler := handler.ShortHandler{
		ShortRepo: repo_impl.NewRedisRepo(client),
		Logger:    log,
	}

	api := router.API{
		Echo:         e,
		ShortHandler: shortHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
