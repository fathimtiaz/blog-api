package main

import (
	"blog-api/api"
	"blog-api/config"
	"blog-api/internal/repository/mysql"
	"blog-api/internal/service"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadDefault()

	repo, err := mysql.NewSqlDB("mysql", cfg.DB.ConnStr.String())
	if err != nil {
		log.Fatal(err)
	}

	userService := service.NewUserService(cfg, repo)

	api.Router(userService).Run(":8080")
}
