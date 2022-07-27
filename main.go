package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restful/controllers"
	"github.com/restful/database"
	"github.com/restful/entity"

	// Altre (-) işareti sadece yan etkileri için kullanılacak paktler için yazılmalı.
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	dbBaglan()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))

}

// Veritabanı bağlantısını burada başlatıyoruz.
func dbBaglan() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "123456",
			DB:         "golangRestFul",
		}

	connectionString := database.BaglantiCumlesiGetir(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})
}

// Bu fonksiyon sayesinde router yönlendirmelerini tanımlıyoruz
func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}
