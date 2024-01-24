package utils

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type (
	ParamQuery struct {
		Page    int
		Limit   int
		Offset  int
		Keyword string
	}
)

func SetParamQuery(c echo.Context) *ParamQuery {
	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limit := c.QueryParam("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	return &ParamQuery{
		Page:   pageInt,
		Limit:  limitInt,
		Offset: (pageInt - 1) * limitInt,
	}
}
