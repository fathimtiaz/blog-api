package main

import (
	"blog-api/api"
	"blog-api/internal/repository/mysql"
	"log"
)

func main() {
	_, err := mysql.NewSqlDB("mysql", "cfg.DB.ConnStr.String()")
	if err != nil {
		log.Fatal(err)
	}

	api.Router().Run(":8080")
}
