package infrastructure

import (
	"f1-hex-api/src/drivers/application"
	"f1-hex-api/src/drivers/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateDriverController struct {
	useCase *application.CreateDriver
}

func NewCreateDriverController(useCase *application.CreateDriver) *CreateDriverController {
	return &CreateDriverController{useCase: useCase}
}

func (cdc *CreateDriverController) Execute(c *gin.Context) {
	var driver entities.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cdc.useCase.Execute(&driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, driver)
}
