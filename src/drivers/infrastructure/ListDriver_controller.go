package infrastructure

import (
	"f1-hex-api/src/drivers/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListDriverController struct {
	useCase *application.ListDriver
}

func NewListDriverController(useCase *application.ListDriver) *ListDriverController {
	return &ListDriverController{useCase: useCase}
}

func (ldc *ListDriverController) Execute(c *gin.Context) {
	drivers, err := ldc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}
