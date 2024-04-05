package tools

import (
	"github.com/gofrs/uuid"
	"hh_test_autho/internal/domain"
	"time"
)

func TestPlug() []domain.Car {
	var cars []domain.Car

	id, _ := uuid.NewV4()
	ownerID, _ := uuid.FromString("2bd08077-cc68-4b9c-8196-50281736c8f5")

	carModelEntity := domain.Car{
		ID:        id,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		OwnerID:   ownerID,
		Mark:      "Lada",
		Model:     "Vesta",
		RegNum:    "X123XX150",
		Year:      2002,
	}
	cars = append(cars, carModelEntity)

	return cars
}
