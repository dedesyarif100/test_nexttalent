package handler

import (
	"net/http"
	position "rest-api-golang-v3/position"
	"strconv"

	"github.com/gin-gonic/gin"
)

type positionHandler struct {
	positionService position.Service
}

func NewPositionHandler(position position.Service) *positionHandler {
	return &positionHandler{position}
}

func (handler *positionHandler) GetPositions(c *gin.Context) {
	positions, err := handler.positionService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": err.Error(),
		})
		return
	}

	var positionsResponse []position.PositionResponse
	for _, position := range positions {
		positionResponse := convertToPositionResponse(&position)
		positionsResponse = append(positionsResponse, positionResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   positionsResponse,
	})
}

func (handler *positionHandler) GetDetailPosition(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	position, err := handler.positionService.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": err.Error(),
		})
		return
	}

	bookResponse := convertToPositionResponse(position)

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "success get data",
		"result":   bookResponse,
	})
}

func convertToPositionResponse(p *position.Position) position.PositionResponse {
	return position.PositionResponse{
		ID:                   p.ID,
		Name:                 p.Name,
		JobDescription:       p.JobDescription,
		MinimumQualification: p.MinimumQualification,
		MinimumExperience:    p.MinimumExperience,
		Skills:               p.Skills,
		Benefit:              p.Benefit,
	}
}
