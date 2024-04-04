package tools

import "hh_test_autho/internal/model"

func TestPlug() []model.Car {
	var cars []model.Car

	carModelEntity := model.Car{
		Mark:    "Lada",
		OwnerID: "2bd08077-cc68-4b9c-8196-50281736c8f5",
		Model:   "Vesta",
		RegNum:  "X123XX150",
		Year:    2002,
	}
	cars = append(cars, carModelEntity)

	return cars
}
