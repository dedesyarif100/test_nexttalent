package main

import (
	"fmt"
	"log"
	country "rest-api-golang-v3/country"
	"rest-api-golang-v3/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@(127.0.0.1:43306)/nexttalent_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR")
	}
	fmt.Println("SUCCESS", db)

	countryRepository := country.NewRepository(db)
	countryService := country.NewService(countryRepository)
	countryHandler := handler.NewCountryHandler(countryService)

	router := gin.Default()
	router.GET("/GetCountry/:name", countryHandler.GetDetailCountry)
	router.GET("/GetCurrentTime", countryHandler.GetCurrentTime)
	router.Run(":8080")
}
