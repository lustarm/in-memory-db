package main

import (
	"imd/src/api"
	"imd/src/db"
	"log"
)

func main() {
    // Init the DB
    log.Println("Initalizing IMD database")
    db.Init()

    log.Println("Starting API")
    api.StartApi()
}
