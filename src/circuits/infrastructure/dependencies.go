package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
)

func ConfigureCircuitRoutes(r *gin.Engine) {
	mysql := NewMySQL()

	// Casos de uso
	createCircuit := application.NewCreateCircuit(mysql)
	listCircuit := application.NewListCircuit(mysql)
	getCircuit := application.NewGetCircuit(mysql)
	updateCircuit := application.NewUpdateCircuit(mysql)
	deleteCircuit := application.NewDeleteCircuit(mysql)
	inscribirPiloto := application.NewInscribirPiloto(mysql)
	getDriversByCircuit := application.NewGetDriversByCircuit(mysql)
	monitorLapTimes := application.NewMonitorLapTimes(mysql)
	monitorLapRecords := application.NewMonitorLapRecords(mysql)

	saveLapTime := application.NewSaveLapTime(mysql)
	waitForIncidents := application.NewWaitForIncidents(mysql)

	// Controladores
	createCircuitController := NewCreateCircuitController(createCircuit)
	listCircuitController := NewListCircuitController(listCircuit)
	getCircuitController := NewGetCircuitController(getCircuit)
	updateCircuitController := NewUpdateCircuitController(updateCircuit)
	deleteCircuitController := NewDeleteCircuitController(deleteCircuit)
	inscribirPilotoController := NewInscribirPilotoController(inscribirPiloto)
	getDriversByCircuitController := NewGetDriversByCircuitController(getDriversByCircuit)
	monitorLapTimesController := NewMonitorLapTimesController(monitorLapTimes)
	monitorLapRecordsController := NewMonitorLapRecordsController(monitorLapRecords)

	saveLapTimeController := NewSaveLapTimeController(saveLapTime)
	waitForIncidentsController := NewWaitForIncidentsController(waitForIncidents)

	// Rutas
	circuits := r.Group("/api/circuitos")
	{
		circuits.POST("/", createCircuitController.Execute)
		circuits.GET("/", listCircuitController.Execute)
		circuits.GET("/:id", getCircuitController.Execute) // Nueva ruta
		circuits.PUT("/:id", updateCircuitController.Execute)
		circuits.DELETE("/:id", deleteCircuitController.Execute)
		circuits.POST("/:id/pilotos", inscribirPilotoController.Execute)
		circuits.GET("/:id/pilotos", getDriversByCircuitController.Execute)
		circuits.GET("/:id/tiempos", monitorLapTimesController.Execute)
		circuits.GET("/:id/records", monitorLapRecordsController.Execute)
		circuits.POST("/:id/tiempos", saveLapTimeController.Execute)
		circuits.GET("/:id/incidentes", waitForIncidentsController.Execute)
	}
}
