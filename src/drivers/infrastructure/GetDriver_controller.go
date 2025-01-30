package infrastructure

import (
	"f1-hex-api/src/drivers/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetDriverController struct {
	useCase *application.GetDriver
}

func NewGetDriverController(useCase *application.GetDriver) *GetDriverController {
	return &GetDriverController{useCase: useCase}
}

func (gdc *GetDriverController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	driver, err := gdc.useCase.Execute(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}
