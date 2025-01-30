package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListCircuitController struct {
	useCase *application.ListCircuit
}

func NewListCircuitController(useCase *application.ListCircuit) *ListCircuitController {
	return &ListCircuitController{useCase: useCase}
}

func (lc *ListCircuitController) Execute(c *gin.Context) {
	circuits, err := lc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, circuits)
}
