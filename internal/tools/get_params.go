package tools

import (
	"github.com/gin-gonic/gin"
	"hh_test_autho/internal/model"
	"strconv"
)

func GetParams(c *gin.Context) (model.Car, error) {

	year := 0
	mark := c.Query("mark")
	ownerId := c.Query("owner_id")
	carModel := c.Query("model")
	regNum := c.Query("reg_num")

	if c.Query("year") != "" {
		year2, err := strconv.Atoi(c.Query("year"))
		if err != nil {
			return model.Car{}, err
		}
		year = year2
	}

	carFilter := model.Car{
		Mark:    mark,
		OwnerID: ownerId,
		Model:   carModel,
		RegNum:  regNum,
		Year:    year,
	}

	return carFilter, nil
}
