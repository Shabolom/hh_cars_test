package tools

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// LimitSlip возвращает сколько элементов нужно вывести и сколько пропустить
func LimitSlip(c *gin.Context) (uint64, uint64, error) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return 0, 0, err
	}

	skip := page*limit - limit

	uLimit := uint64(limit)
	uSkip := uint64(skip)

	return uLimit, uSkip, nil
}
