package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"f1-hex-api/src/circuits/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SaveLapTimeController struct {
	useCase *application.SaveLapTime
}

func NewSaveLapTimeController(useCase *application.SaveLapTime) *SaveLapTimeController {
	return &SaveLapTimeController{useCase: useCase}
}

func (sltc *SaveLapTimeController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	var lapTime entities.LapTime
	if err := c.ShouldBindJSON(&lapTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lapTime.CircuitoID = circuitoID

	if err := sltc.useCase.Execute(&lapTime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, lapTime)
}
