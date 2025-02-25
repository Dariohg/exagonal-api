package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"f1-hex-api/src/circuits/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InscribirPilotoController struct {
	useCase *application.InscribirPiloto
}

func NewInscribirPilotoController(useCase *application.InscribirPiloto) *InscribirPilotoController {
	return &InscribirPilotoController{useCase: useCase}
}

func (ipc *InscribirPilotoController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req domain.InscripcionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ipc.useCase.Execute(circuitoID, req.ConductorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Piloto inscrito correctamente"})
}
