package infrastructure

import (
	"f1-hex-api/src/drivers/application"
	"github.com/gin-gonic/gin"
)

func ConfigureDriverRoutes(r *gin.Engine) {
	mysql := NewMySQL()

	// Casos de uso
	createDriver := application.NewCreateDriver(mysql)
	listDriver := application.NewListDriver(mysql)
	getDriver := application.NewGetDriver(mysql)
	updateDriver := application.NewUpdateDriver(mysql)
	deleteDriver := application.NewDeleteDriver(mysql)

	// Controladores
	createDriverController := NewCreateDriverController(createDriver)
	listDriverController := NewListDriverController(listDriver)
	getDriverController := NewGetDriverController(getDriver)
	updateDriverController := NewUpdateDriverController(updateDriver)
	deleteDriverController := NewDeleteDriverController(deleteDriver)

	// Rutas
	drivers := r.Group("/api/conductores")
	{
		drivers.POST("/", createDriverController.Execute)
		drivers.GET("/", listDriverController.Execute)
		drivers.GET("/:id", getDriverController.Execute)
		drivers.PUT("/:id", updateDriverController.Execute)
		drivers.DELETE("/:id", deleteDriverController.Execute)
	}
}
