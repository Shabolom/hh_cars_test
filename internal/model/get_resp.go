package model

import "hh_test_autho/internal/domain"

type GetResp struct {
	Cars        []domain.Car
	TotalAmount int
}
