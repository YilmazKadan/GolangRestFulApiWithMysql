// Bu dosya içerisinde veritabanı bağlantı ayarları yapılmaktadır.

package database

import "fmt"

// Veritabanı ayarlamaları

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

// Bu fonksiyon bize veritabanı bağlantı cümlesini döndürecek
var baglantiCumlesiGetir = func(config Config) string {
	// Sprintf geriye işlenen stringi döndürür.
	baglantiCumlesi := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&collation=utf8_turkish_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return baglantiCumlesi
}
