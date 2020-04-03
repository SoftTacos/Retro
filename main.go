package main

import (
	"fmt"
	"os"

	gopg "github.com/go-pg/pg"
)

func main() {
	dburl := os.Getenv("DBURL")
	db, err := gopg.ParseURL(dburl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db)
	}
	fmt.Println(dburl)

}
