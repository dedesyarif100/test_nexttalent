package handler

import (
	"net/http"
	employmenttype "rest-api-golang-v3/employment_type"
	"strconv"

	"github.com/gin-gonic/gin"
)

type employmentTypeHandler struct {
	employmentTypeService employmenttype.Service
}

func NewEmploymentTypeHandler(employmentTypeService employmenttype.Service) *employmentTypeHandler {
	return &employmentTypeHandler{employmentTypeService}
}

func (handler *employmentTypeHandler) GetEmploymentTypes(c *gin.Context) {
	employmentTypes, err := handler.employmentTypeService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []employmenttype.EmploymentTypeResponse
	for _, employment := range employmentTypes {
		bookResponse := convertToEmploymentTypeResponse(&employment)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   booksResponse,
	})
}

func (handler *employmentTypeHandler) GetDetailEmploymentType(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	employeeType, err := handler.employmentTypeService.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	employeeTypeResponse := convertToEmploymentTypeResponse(employeeType)

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   employeeTypeResponse,
	})
}

func convertToEmploymentTypeResponse(employeeType *employmenttype.EmploymentType) employmenttype.EmploymentTypeResponse {
	return employmenttype.EmploymentTypeResponse{
		ID:   employeeType.ID,
		Name: employeeType.Name,
	}
}
