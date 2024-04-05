package service

import (
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
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
		log.WithField("component", "service").Debug(err)
		return err
	}

	for _, car := range cars {
		ownerUuid, err2 := uuid.FromString(car.OwnerID)
		if err2 != nil {
			log.WithField("component", "service").Debug(err2)
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

	carID, err := uuid.FromString(carStrID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	ownerUuid, err := uuid.FromString(car.OwnerID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
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

	err = carRepo.Update(carEntity, carID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CarsService) Delete(carStrID string) error {

	carID, err := uuid.FromString(carStrID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	err = carRepo.Delete(carID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CarsService) Get(car model.Car, limit, skip uint64) (model.GetResp, error) {
	var ownerID uuid.UUID

	if car.OwnerID != "" {
		ownerID2, err := uuid.FromString(car.OwnerID)
		if err != nil {
			log.WithField("component", "service").Debug(err)
			return model.GetResp{}, err
		}
		ownerID = ownerID2
	}

	result, err := carRepo.Get(car, ownerID, limit, skip)
	if err != nil {
		return model.GetResp{}, err
	}

	return result, nil
}

func (cs *CarsService) GetID(strCarID string) (domain.Car, error) {
	carID, err := uuid.FromString(strCarID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return domain.Car{}, err
	}

	result, err := carRepo.GetID(carID)
	if err != nil {
		return domain.Car{}, err
	}

	return result, nil
}
