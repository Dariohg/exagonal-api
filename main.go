package main

import (
	"f1-hex-api/src/circuits/infrastructure"
	drivers "f1-hex-api/src/drivers/infrastructure"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	infrastructure.ConfigureCircuitRoutes(r)

	drivers.ConfigureDriverRoutes(r)

	puerto := os.Getenv("PUERTO")
	if puerto == "" {
		puerto = "8080"
	}

	log.Printf("Servidor iniciando en el puerto %s", puerto)
	if err := r.Run(":" + puerto); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
