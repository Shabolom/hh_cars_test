package tools

import (
	"github.com/gin-gonic/gin"
	"hh_test_autho/internal/model"
	"strconv"
)

func GetParams(c *gin.Context) (model.Car, error) {

	year := 0
	mark := c.Query("mark")
	owner_id := c.Query("owner_id")
	carModel := c.Query("model")
	reg_num := c.Query("reg_num")

	if c.Query("year") != "" {
		year2, err := strconv.Atoi(c.Query("year"))
		if err != nil {
			return model.Car{}, err
		}
		year = year2
	}

	carFilter := model.Car{
		Mark:    mark,
		OwnerID: owner_id,
		Model:   carModel,
		RegNum:  reg_num,
		Year:    year,
	}

	return carFilter, nil
}
