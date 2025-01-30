package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
)

func ConfigureCircuitRoutes(r *gin.Engine) {
	// Repositorio compartido
	mysql := NewMySQL()

	// Casos de uso
	createCircuit := application.NewCreateCircuit(mysql)
	listCircuit := application.NewListCircuit(mysql)
	getCircuit := application.NewGetCircuit(mysql)
	updateCircuit := application.NewUpdateCircuit(mysql)
	deleteCircuit := application.NewDeleteCircuit(mysql)

	// Controladores
	createCircuitController := NewCreateCircuitController(createCircuit)
	listCircuitController := NewListCircuitController(listCircuit)
	getCircuitController := NewGetCircuitController(getCircuit)
	updateCircuitController := NewUpdateCircuitController(updateCircuit)
	deleteCircuitController := NewDeleteCircuitController(deleteCircuit)

	// Rutas
	circuits := r.Group("/api/circuitos")
	{
		circuits.POST("/", createCircuitController.Execute)
		circuits.GET("/", listCircuitController.Execute)
		circuits.GET("/:id", getCircuitController.Execute) // Nueva ruta
		circuits.PUT("/:id", updateCircuitController.Execute)
		circuits.DELETE("/:id", deleteCircuitController.Execute)
	}
}
