package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"f1-hex-api/src/circuits/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCircuitController struct {
	useCase *application.CreateCircuit
}

func NewCreateCircuitController(useCase *application.CreateCircuit) *CreateCircuitController {
	return &CreateCircuitController{useCase: useCase}
}

func (cc *CreateCircuitController) Execute(c *gin.Context) {
	var circuit entities.Circuit
	if err := c.ShouldBindJSON(&circuit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.useCase.Execute(&circuit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, circuit)
}
