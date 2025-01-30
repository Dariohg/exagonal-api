package infrastructure

import (
	"f1-hex-api/src/drivers/application"
	"f1-hex-api/src/drivers/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateDriverController struct {
	useCase *application.UpdateDriver
}

func NewUpdateDriverController(useCase *application.UpdateDriver) *UpdateDriverController {
	return &UpdateDriverController{useCase: useCase}
}

func (udc *UpdateDriverController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var driver entities.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver.ID = idInt

	if err := udc.useCase.Execute(&driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}
