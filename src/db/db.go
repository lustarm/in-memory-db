package db

import (
	"errors"
	"log"
)

type Database map[string]string

var Db Database

func Init() {
    Db = make(Database)
}

func (db Database) Create(key string, value string) error {
    if _, found := db[key]; found {
        return errors.New("Key is already has a value set, please use Update")
    }

    db[key] = value
    return nil
}

func (db Database) Update(key string, value string) error {
    if _, found := db[key]; !found {
        return errors.New("Failed to find key value " + key)
    }

    db[key] = value
    log.Println("Updated " + key + " with " + value)
    return nil
}

func (db Database) Read(key string) (string, error) {
    if _, found := db[key]; !found {
        return "", errors.New("Failed to find key value " + key)
    }

    return db[key], nil
}
