package dependencies

import (
	"log"
	"os"
	"status_esp32_service/src/esp32/application"
	"status_esp32_service/src/esp32/infraestructure/adapters"
	"status_esp32_service/src/esp32/infraestructure/controllers"
	"status_esp32_service/src/esp32/infraestructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitEsp32(r *gin.Engine){
	ps, _ := adapters.NewMySQL()
	err := godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading.env file")
    }
	key := os.Getenv("SECRET_KEY")
	change_status_use_case := application.NewChangeStatusUseCase(ps)
	change_status_controller := controllers.NewChangeStatusController(change_status_use_case)
	routes.Esp32Router(r, key, change_status_controller)
}