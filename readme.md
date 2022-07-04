# Mysql ve Golang ile RestFul API

## Kullanılan paketler

`
github.com/gorilla/mux 
github.com/jinzhu/gorm
github.com/go-sql-driver/mysql
`

1. **gorilla/mux** paketi endpointlerimiz için rota ve yakalayıcı tanımlamak için
2. **jinzhu/gorm** paketi MYSQL'de ORM yapısı ile veritabanı CRUD işlemlerini yapabilmek için
3. **go-sql-driver/mysql** MYSQL sürücüsü


## Yapılandırma

Yapılandırma için main.go dosyasına gidip kendi veritabanı bilgilerinizi girmeniz lazım.