package main

import (
	"fmt"
	"log"
	"os"

	// queries "github.com/softtacos/retroBot/queries"
	"github.com/softtacos/retroBot/service"
	util "github.com/softtacos/retroBot/util"
)

//postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
func main() {
	dburl := os.Getenv("DBURL")
	db, err := util.NewGoPgDb(dburl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db)
	}

	retroService := service.NewRetroService(db)
	log.Println(retroService)
	// server:=server.NewHttpRetroServer(
}
