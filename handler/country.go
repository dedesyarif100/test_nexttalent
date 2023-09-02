package handler

import (
	"net/http"
	country "rest-api-golang-v3/country"

	"github.com/gin-gonic/gin"
)

type countryHandler struct {
	countryService country.Service
}

func NewCountryHandler(country country.Service) *countryHandler {
	return &countryHandler{country}
}

func (handler *countryHandler) GetDetailCountry(c *gin.Context) {
	name := c.Param("name")
	country, err := handler.countryService.FindById(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": err.Error(),
		})
		return
	}

	response := convertToCountryResponse(country)
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   response,
	})
}

func convertToCountryResponse(c *country.Country) country.CountryResponse {
	return country.CountryResponse{
		Country: c.Country,
	}
}
