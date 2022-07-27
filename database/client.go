package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/restful/entity"
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

// Bu migrate fonksiyonu sayesinde tablomuzu belirlenen structerdaki gibi oluşturuyoruz.
func Migrate(table *entity.Person) {
	Connector.Table("Uyeler").AutoMigrate(&table)
	log.Println("Table migrated")
}
