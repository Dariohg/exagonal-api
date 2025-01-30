package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"f1-hex-api/src/circuits/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateCircuitController struct {
	useCase *application.UpdateCircuit
}

func NewUpdateCircuitController(useCase *application.UpdateCircuit) *UpdateCircuitController {
	return &UpdateCircuitController{useCase: useCase}
}

func (uc *UpdateCircuitController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var circuit entities.Circuit
	if err := c.ShouldBindJSON(&circuit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	circuit.ID = idInt

	if err := uc.useCase.Execute(&circuit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, circuit)
}
