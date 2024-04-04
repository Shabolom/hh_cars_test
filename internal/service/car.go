package service

import (
	"github.com/gofrs/uuid"
	"hh_test_autho/internal/domain"
	"hh_test_autho/internal/model"
	"hh_test_autho/internal/repository"
	"time"
)

type CarsService struct {
}

func NEwCarsService() *CarsService {
	return &CarsService{}
}

var carRepo = repository.NewCarRepo()

func (cs *CarsService) Post(cars []model.Car) error {
	var carsDomainMass []domain.Car

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	for _, car := range cars {
		ownerUuid, err2 := uuid.FromString(car.OwnerID)
		if err2 != nil {
			return err2
		}

		carsEntity := domain.Car{
			ID:        id,
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
			OwnerID:   ownerUuid,
			Mark:      car.Mark,
			Model:     car.Model,
			RegNum:    car.RegNum,
			Year:      car.Year,
		}

		carsDomainMass = append(carsDomainMass, carsEntity)
	}

	err = carRepo.Post(carsDomainMass)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CarsService) Update(car model.Car, carStrID string) error {

	carID, err2 := uuid.FromString(carStrID)
	if err2 != nil {
		return err2
	}

	ownerUuid, err2 := uuid.FromString(car.OwnerID)
	if err2 != nil {
		return err2
	}

	// оставлю возможность смены владельца автомабиля.
	carEntity := domain.Car{
		UpdateAt: time.Now(),
		OwnerID:  ownerUuid,
		Mark:     car.Mark,
		Model:    car.Model,
		RegNum:   car.RegNum,
		Year:     car.Year,
	}

	err := carRepo.Update(carEntity, carID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CarsService) Delete(carStrID string) error {

	carID, err := uuid.FromString(carStrID)
	if err != nil {
		return err
	}

	err = carRepo.Delete(carID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CarsService) Get(car model.Car, limit, skip uint64) ([]domain.Car, error) {
	var ownerID uuid.UUID

	if car.OwnerID != "" {
		ownerID2, err := uuid.FromString(car.OwnerID)
		if err != nil {
			return []domain.Car{}, err
		}
		ownerID = ownerID2
	} else {
		ownerID = uuid.Nil
	}

	result, err := carRepo.Get(car, &ownerID, limit, skip)
	if err != nil {
		return []domain.Car{}, err
	}

	return result, nil
}

func (cs *CarsService) GetID(strCarID string) (domain.Car, error) {
	carID, err := uuid.FromString(strCarID)
	if err != nil {
		return domain.Car{}, err
	}

	result, err := carRepo.GetID(carID)
	if err != nil {
		return domain.Car{}, err
	}

	return result, nil
}
