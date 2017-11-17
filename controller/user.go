package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"godemo/service"
)
func Save(c echo.Context) error{
	return c.JSON(http.StatusCreated,service.Save())
}


func  GetAll(c echo.Context) error {
	//service.QuerybySql()
	return c.JSON(http.StatusCreated,service.Acclist())
}

func Query(c echo.Context) error{
	sql := "select id,name,username from usertable"
	return c.JSON(http.StatusOK,service.QuerybySql(sql))
}
