package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"godemo/service"
)



func  GetAll(c echo.Context) error {

	return c.JSON(http.StatusCreated,service.Frist())
}
