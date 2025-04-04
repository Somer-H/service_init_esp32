package repositories

import "status_esp32_service/src/esp32/domain/entities"

type IEsp32Repository interface {
	ChangeStatus(name string, status string) (*entities.Esp32, error)
}