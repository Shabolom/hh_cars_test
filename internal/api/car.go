package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

// Post Добавления новых автомобилей в формате принятия массива номеров с последующим обращением на другой сервис и получения данных об этих машинах.
// (сделано так чтобы записи не дублировались, если будет попытка записать повторно ту же запись выдаст ошибку)
//
// @Summary	Добавления новых автомобилей в формате принятия массива номеров с последующим обращением на другой сервис и получения данных об этих машинах.
// @Accept	json
// @Produce	json
// @Tags	car
// @Param	ввод	body		model.RegNums	true	"введите массив номеров необходимых машин"
// @Success	201		{object}	domain.Car
// @Failure	400		{object}	model.Error
// @Router	/api/car [post]
func (ca *CarApi) Post(c *gin.Context) {
	var regNums model.RegNums

	err := tools.ShortUnmarshal(c.Request.Body, &regNums)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "api").Debug(err)
		return
	}

	result, err := carService.Post(regNums)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
	defer c.Request.Body.Close()
}

// Update Изменение одного или нескольких полей по идентификатору.
//
// @Summary	Изменение одного или нескольких полей по идентификатору.
// @Accept	json
// @Produce	json
// @Tags	car
// @Param	ввод	body		model.Car	true	"выберите данные которые хотите изменить"
// @Param	id		path		string	true	"укажите id машины"
// @Success	200		{string}	string 	"данные упешно изменены"
// @Failure	400		{object}	model.Error
// @Router	/api/car/{id} [put]
func (ca *CarApi) Update(c *gin.Context) {
	var car model.Car
	carID := c.Param("id")

	err := tools.ShortUnmarshal(c.Request.Body, &car)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "api").Debug(err)
		return
	}

	err = carService.Update(car, carID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.String(http.StatusOK, "данные упешно изменены")
	defer c.Request.Body.Close()
}

// Delete Удаления по идентификатору.
//
// @Summary	Удаления по идентификатору.
// @Accept	json
// @Produce	json
// @Tags	car
// @Param	id		path		string	true	"укажите id машины"
// @Success	204		{string}	string ""
// @Failure	400		{object}	model.Error
// @Router	/api/car/{id} [delete]
func (ca *CarApi) Delete(c *gin.Context) {
	carID := c.Param("id")

	// обратить внимание, что происходит soft delete.
	err := carService.Delete(carID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Status(http.StatusNoContent)
	defer c.Request.Body.Close()
}

// Get Получение данных с фильтрацией по всем полям и пагинацией, данные передаются в query params.
//
// @Summary	Получение данных с фильтрацией по всем полям и пагинацией, данные передаются в query params.
// @Accept	json
// @Produce	json
// @Tags	car
// @Param	mark	query		string	false	"это поле отвечает за марку машины"
// @Param	owner_id 	query	string	false	"это поле отвечает за марку машины"
// @Param	model	query		string	false	"это поле отвечает за модель машины"
// @Param	reg_num	query		string	false	"это поле отвечает за номер машины"
// @Param	year	query		string	false	"это поле отвечает за год выпуска машины"
// @Param	page	query		string	false	"это поле отвечает за страницу"
// @Param	limit	query		string	false	"это поле отвечает за количество элементов на странице"
// @Success	200		{object}	model.GetResp
// @Success	200		{string}	string
// @Failure	400		{object}	model.Error
// @Router	/api/car [get]
func (ca *CarApi) Get(c *gin.Context) {
	limit, skip, err := tools.LimitSlip(c)

	car, err := tools.GetParams(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "api").Debug(err)
		return
	}

	result, err := carService.Get(car, limit, skip)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
	defer c.Request.Body.Close()
}

// GetID Получение данных получение данных по id машины для дальнейшего заполнения ручки Update для удобства пользователя.
//
// @Summary	Получение данных получение данных по id машины для дальнейшего заполнения ручки Update для удобства пользователя.
// @Accept	json
// @Produce	json
// @Tags	car
// @Param	id	path			string	true	"передайте id машины"
// @Success	200		{object}	domain.Car
// @Failure	400		{object}	model.Error
// @Router	/api/car/{id} [get]
func (ca *CarApi) GetID(c *gin.Context) {
	carID := c.Param("id")

	result, err := carService.GetID(carID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
	defer c.Request.Body.Close()
}
