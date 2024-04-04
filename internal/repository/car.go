package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"hh_test_autho/config"
	"hh_test_autho/internal/domain"
	"hh_test_autho/internal/model"
	"time"
)

type CarRepo struct {
}

func NewCarRepo() *CarRepo {
	return &CarRepo{}
}

func (cr *CarRepo) Post(cars []domain.Car) error {

	tx, err := config.Pool.Begin(context.TODO())
	if err != nil {
		return err
	}

	for _, car := range cars {
		sql, args, err2 := config.Sq.
			Insert("cars").
			Columns("id", "created_at", "updated_at", "owner_id", "mark", "model", "reg_num", "year").
			Values(car.ID, car.CreatedAt, car.UpdateAt, car.OwnerID, car.Mark, car.Model, car.RegNum, car.Year).
			ToSql()
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec(context.TODO(), sql, args...)
		if err2 != nil {
			tx.Rollback(context.TODO())
			return err2
		}
	}
	err = tx.Commit(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (cr *CarRepo) Update(car domain.Car, carID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("cars").
		Set("updated_at", car.UpdateAt).
		Set("owner_id", car.OwnerID).
		Set("mark", car.Mark).
		Set("model", car.Model).
		Set("reg_num", car.RegNum).
		Set("year", car.Year).
		Where("id = $7", carID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CarRepo) Delete(carID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("cars").
		Set("deleted_at", time.Now()).
		Where("id = $2", carID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CarRepo) Get(filter model.Car, ownerID *uuid.UUID, limit, skip uint64) ([]domain.Car, error) {
	var cars []domain.Car
	var car domain.Car

	fmt.Println("ownerID")
	fmt.Println(ownerID.String())
	sql, args, err := config.Sq.
		Select("c.id", "c.created_at", "c.updated_at", "c.owner_id", "c.mark", "c.model", "reg_num", "c.year").
		From("cars c").
		Where("c.owner_id = ? OR ? = ''", ownerID, ownerID).
		//Where("c.owner_id = $1 OR $2::text = ''", ownerID, ownerID).
		//Where("c.reg_num = $4 OR $5 = ''", filter.RegNum, filter.RegNum).
		//Where("c.model = $6 OR $7 = ''", filter.Model, filter.Model).
		//Where("c.mark = $8 OR $9 = ''", filter.Mark, filter.Mark).
		//Where("c.year = $10 OR $11 = 0", filter.Year, filter.Year).
		//Where("c.deleted_at IS NULL").
		Limit(limit).
		Offset(skip).
		ToSql()

	println(sql)
	fmt.Println(args)

	if err != nil {
		return []domain.Car{}, err
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		return []domain.Car{}, err
	}

	for rows.Next() {
		err2 := rows.Scan(
			&car.ID,
			&car.CreatedAt,
			&car.UpdateAt,
			&car.OwnerID,
			&car.Mark,
			&car.Model,
			&car.RegNum,
			&car.Year)
		if err2 != nil {
			return []domain.Car{}, err2
		}
		cars = append(cars, car)
	}
	defer rows.Close()

	return cars, nil
}

func (cr *CarRepo) GetID(carID uuid.UUID) (domain.Car, error) {
	var car domain.Car

	sql, args, err := config.Sq.
		Select("c.id", "c.created_at", "c.updated_at", "c.owner_id", "c.mark", "c.model", "reg_num", "c.year").
		From("cars c").
		Where("c.id = $1", carID).
		ToSql()
	if err != nil {
		return domain.Car{}, err
	}

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(
		&car.ID,
		&car.CreatedAt,
		&car.UpdateAt,
		&car.OwnerID,
		&car.Mark,
		&car.Model,
		&car.RegNum,
		&car.Year)
	if err != nil {
		return domain.Car{}, err
	}

	return car, nil
}
