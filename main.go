package main

import (
	"fmt"
	"log"
	"os"

	// queries "github.com/softtacos/retroBot/queries"
	server "github.com/softtacos/retroBot/server"
	"github.com/softtacos/retroBot/service"
	util "github.com/softtacos/retroBot/util"
)

//postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
func main() {
	dburl := os.Getenv("DBURL")
	port := ":9001"
	errChan := make(chan error, 1)

	db, err := util.NewGoPgDb(dburl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db)
	}

	retroService := service.NewRetroService(db)
	log.Println(retroService)
	retroServer := server.NewRetroServer(retroService, port, errChan)
	retroServer.Start()
	log.Println("RETRO SERVICE LISTENING ON PORT :", port)
	for {
		log.Println(<-errChan)
	}
}
