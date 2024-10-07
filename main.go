package main

import (
	"imd/src/db"
	"log"
)

func main() {
    // Init the DB
    log.Println("Initalizing IMD database")
    db := db.Init()

    err := db.Create("test", "hello")
    if err != nil {
        log.Fatalln(err.Error())
    }

    test, err := db.Read("test")

    if err != nil {
        log.Fatalln(err.Error())
        return
    }

    log.Println(test)
}
