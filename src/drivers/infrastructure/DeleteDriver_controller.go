package infrastructure

import (
	"f1-hex-api/src/drivers/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteDriverController struct {
	useCase *application.DeleteDriver
}

func NewDeleteDriverController(useCase *application.DeleteDriver) *DeleteDriverController {
	return &DeleteDriverController{useCase: useCase}
}

func (ddc *DeleteDriverController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ddc.useCase.Execute(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Conductor eliminado correctamente"})
}
