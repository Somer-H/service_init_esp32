package routes

import (
	"status_esp32_service/core"
	"status_esp32_service/src/esp32/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Esp32Router(r *gin.Engine, key string, changeSatatusController *controllers.ChangeStatusController){
	v1 := r.Group("/v1/esp32")
	protectedRoute := v1.Group("/protected")
	protectedRoute.Use(core.RoleMiddleware(key, []string{"controller"}))
    protectedRoute.POST("/change_status", changeSatatusController.Run)
}