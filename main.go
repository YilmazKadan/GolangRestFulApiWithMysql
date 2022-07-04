package main

import (
	"github.com/restful/database"
	// Altre (-) işareti sadece yan etkileri için kullanılacak paktler için yazılmalı.
	_"github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "golangRestFul",
		}

	connectionString := database.BaglantiCumlesiGetir(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
