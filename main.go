package main

import (
	"github.com/aaalik/ke-jepang/bootstrap"
	"github.com/aaalik/ke-jepang/helper"
	routers "github.com/aaalik/ke-jepang/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bootstrap.InitLogger()
	helper.Log = bootstrap.Log

	db := bootstrap.Database{}
	db.Connect()
	defer db.CloseConnection()

	routers.SetupRouter()

	forever := make(chan bool)
	<-forever
}
