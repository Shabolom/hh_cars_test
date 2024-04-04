package api

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"hh_test_autho/config"
	"hh_test_autho/internal/model"
	"hh_test_autho/internal/service"
	"hh_test_autho/internal/tools"
	"net/http"
)

type CarApi struct {
}

func NewCarApi() *CarApi {
	return &CarApi{}
}

var carService = service.NEwCarsService()

func (ca *CarApi) Post(c *gin.Context) {
	var regNums model.RegNums
	var cars []model.Car

	err := tools.ShortUnmarshal(c.Request.Body, &regNums)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	if config.Env.Production == true {

		for _, regNum := range regNums.RegNums {
			reader := bytes.NewReader([]byte(regNum))

			resp, err2 := tools.RequestCreator("GET", config.Env.ConnectionGet, reader)
			if err2 != nil {
				tools.CreateError(http.StatusBadRequest, err2, c)
				log.WithField("component", "rest").Debug(err2)
				return
			}

			err2 = tools.ShortUnmarshal(resp.Body, &cars)
			if err2 != nil {
				tools.CreateError(http.StatusBadRequest, err2, c)
				log.WithField("component", "rest").Debug(err2)
				return
			}
		}

		return
	}

	err = carService.Post(tools.TestPlug())
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.String(http.StatusCreated, "данные занесены")
}

func (ca *CarApi) Update(c *gin.Context) {
	var car model.Car
	carID := c.Param("id")

	err := tools.ShortUnmarshal(c.Request.Body, &car)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	err = carService.Update(car, carID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.String(http.StatusOK, "данные упешно изменены")
}

func (ca *CarApi) Delete(c *gin.Context) {
	carID := c.Param("id")

	err := carService.Delete(carID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.String(http.StatusOK, "запись удалена")
}

func (ca *CarApi) Get(c *gin.Context) {
	limit, skip, err := tools.LimitSlip(c)
	car, err := tools.GetParams(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	result, err := carService.Get(car, limit, skip)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ca *CarApi) GetID(c *gin.Context) {
	carID := c.Param("id")

	result, err := carService.GetID(carID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
