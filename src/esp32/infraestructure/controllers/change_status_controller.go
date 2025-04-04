package controllers

import (
	"status_esp32_service/src/esp32/application"
	entities "status_esp32_service/src/esp32/application/entites"

	"github.com/gin-gonic/gin"
)

type ChangeStatusController struct {
	useCase *application.ChangeStatusUseCase
}

func NewChangeStatusController(useCase *application.ChangeStatusUseCase) *ChangeStatusController {
    return &ChangeStatusController{useCase: useCase}
}

func (csc *ChangeStatusController) Run(c *gin.Context){
    var status entities.Status
	if err := c.ShouldBindJSON(&status); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	name := c.Params.ByName("name")
	if name == "" {
		c.JSON(404, gin.H{"error": "ID is required"})
        return
    }
	result, err := csc.useCase.Execute(name, status.Status)
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"status": result})
}