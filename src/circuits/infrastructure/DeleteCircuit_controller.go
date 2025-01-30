package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteCircuitController struct {
	useCase *application.DeleteCircuit
}

func NewDeleteCircuitController(useCase *application.DeleteCircuit) *DeleteCircuitController {
	return &DeleteCircuitController{useCase: useCase}
}

func (dc *DeleteCircuitController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := dc.useCase.Execute(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Circuito eliminado correctamente"})
}
