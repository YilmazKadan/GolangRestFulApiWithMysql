package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Bağlantı değişkeni CRUD işlemlerini yapabilmek için
var Connector *gorm.DB

// Veritabanı bağlantısını oluşturuyoruz

func Connect(connectionString string) error {
	var err error

	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Bağlantı başarılı bir şekilde gerçekleşti")

	return nil
}
