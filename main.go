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
	if err := godotenv.Load("dev.env"); err != nil {
		return
	}
}

func main() {
	// write log
	log, _ := log.WriteLog()

	// Dùng dưới local
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// Dùng trên server heroku
	// redisUrl := os.Getenv("REDIS_URL")

	// connect redis
	client := &db.Redis{
		// Dùng dưới local
		Host: redisHost,
		Port: redisPort,

		// Dùng trên server heroku
		// Url: redisUrl,

		Logger: log,
	}
	client.NewRedis()

	// postgres details
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")

	// connect postgres
	sql := &db.Sql{
		Host:     host,
		Port:     port,
		UserName: username,
		Password: password,
		DbName:   dbname,
		Logger:   log,
	}
	sql.Connect()
	defer sql.Close()

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
