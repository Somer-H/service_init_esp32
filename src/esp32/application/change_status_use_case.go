package application

import (
	"status_esp32_service/src/esp32/domain/entities"
	"status_esp32_service/src/esp32/domain/repositories"
)

type ChangeStatusUseCase struct {
	repository repositories.IEsp32Repository
}

func NewChangeStatusUseCase(repository repositories.IEsp32Repository) *ChangeStatusUseCase {
    return &ChangeStatusUseCase{repository: repository}
}

func (uc *ChangeStatusUseCase) Execute(name string, status string) (*entities.Esp32, error) {
    esp32, err := uc.repository.ChangeStatus(name, status)
    if err != nil {
        return nil, err
    }
	return esp32, nil
}
