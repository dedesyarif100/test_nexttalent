package main

import (
	"fmt"
	"log"
	employmenttype "rest-api-golang-v3/employment_type"
	"rest-api-golang-v3/handler"
	position "rest-api-golang-v3/position"
	positionlevel "rest-api-golang-v3/position_level"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@(127.0.0.1:43306)/nawadata?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR")
	}
	fmt.Println("SUCCESS", db)

	employmentTypeRepository := employmenttype.NewRepository(db)
	employmentTypeService := employmenttype.NewService(employmentTypeRepository)
	employmentTypeHandler := handler.NewEmploymentTypeHandler(employmentTypeService)

	positionLevelRepository := positionlevel.NewRepository(db)
	positionLevelService := positionlevel.NewService(positionLevelRepository)
	positionLevelHandler := handler.NewPositionLevelHandler(positionLevelService)

	positionRepository := position.NewRepository(db)
	positionService := position.NewService(positionRepository)
	positionHandler := handler.NewPositionHandler(positionService)

	router := gin.Default()
	v3 := router.Group("/v3")

	v3.GET("/employment-types", employmentTypeHandler.GetEmploymentTypes)
	v3.GET("/employment-type/:id", employmentTypeHandler.GetDetailEmploymentType)

	v3.GET("/position-levels", positionLevelHandler.GetPositionLevel)
	v3.GET("/position-level/:id", positionLevelHandler.GetDetailPositionLevel)

	v3.GET("/positions", positionHandler.GetPositions)
	v3.GET("/position/:id", positionHandler.GetDetailPosition)

	router.Run(":9999")
}
