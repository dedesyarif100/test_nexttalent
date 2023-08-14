package handler

import (
	"net/http"
	positionlevel "rest-api-golang-v3/position_level"
	"strconv"

	"github.com/gin-gonic/gin"
)

type positionLevelHandler struct {
	positionLevelService positionlevel.Service
}

func NewPositionLevelHandler(positionLevelService positionlevel.Service) *positionLevelHandler {
	return &positionLevelHandler{positionLevelService}
}

func (handler *positionLevelHandler) GetPositionLevel(c *gin.Context) {
	positionLevels, err := handler.positionLevelService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var positionLevelsResponse []positionlevel.PositionLevelResponse
	for _, position := range positionLevels {
		positionLevelResponse := convertToPositionLevelResponse(&position)
		positionLevelsResponse = append(positionLevelsResponse, positionLevelResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   positionLevelsResponse,
	})
}

func (handler *positionLevelHandler) GetDetailPositionLevel(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	positionLevel, err := handler.positionLevelService.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	positionLevelResponse := convertToPositionLevelResponse(positionLevel)

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   positionLevelResponse,
	})
}

func convertToPositionLevelResponse(positionLevel *positionlevel.PositionLevel) positionlevel.PositionLevelResponse {
	return positionlevel.PositionLevelResponse{
		ID:   positionLevel.ID,
		Name: positionLevel.Name,
	}
}
