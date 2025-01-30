package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetCircuitController struct {
	useCase *application.GetCircuit
}

func NewGetCircuitController(useCase *application.GetCircuit) *GetCircuitController {
	return &GetCircuitController{useCase: useCase}
}

func (gc *GetCircuitController) Execute(c *gin.Context) {
	// Obtener el ID de los parámetros de la URL
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Ejecutar el caso de uso
	circuit, err := gc.useCase.Execute(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, circuit)
}
