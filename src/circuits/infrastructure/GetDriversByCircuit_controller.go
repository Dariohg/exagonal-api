package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetDriversByCircuitController struct {
	useCase *application.GetDriversByCircuit
}

func NewGetDriversByCircuitController(useCase *application.GetDriversByCircuit) *GetDriversByCircuitController {
	return &GetDriversByCircuitController{useCase: useCase}
}

func (gdcc *GetDriversByCircuitController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	drivers, err := gdcc.useCase.Execute(circuitoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"circuito_id": circuitoID,
		"pilotos":     drivers,
	})
}
