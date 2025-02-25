package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"f1-hex-api/src/circuits/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SavePositionController struct {
	useCase *application.SavePosition
}

func NewSavePositionController(useCase *application.SavePosition) *SavePositionController {
	return &SavePositionController{useCase: useCase}
}

func (spc *SavePositionController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	var position entities.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position.CircuitoID = circuitoID

	if err := spc.useCase.Execute(&position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, position)
}
