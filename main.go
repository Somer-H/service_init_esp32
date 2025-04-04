package main

import (
	"status_esp32_service/src/esp32/infraestructure/dependencies"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Authorization"},
		MaxAge:        12 * time.Hour,
	}))


	dependencies.InitEsp32(r)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}