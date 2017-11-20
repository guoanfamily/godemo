package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"godemo/service"
	// "github.com/mitchellh/mapstructure"
)
func Save(c echo.Context) error{
	return c.JSON(http.StatusCreated,service.Save())
}


func  GetAll(c echo.Context) error {
	//service.QuerybySql()
	return c.JSON(http.StatusCreated,service.Acclist())
}
type myuser struct {
	Id string
	Name string
}
func Query(c echo.Context) error{
	sql := "select id,name from usertable"
	return c.JSON(http.StatusOK,service.QuerybySql(sql))
}

func Find(c echo.Context) error{
	return c.JSON(http.StatusOK,service.Select())
}
