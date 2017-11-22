package router
import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"godemo/controller"
)


func Router(){
	e:=echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST("/users",controller.GetAll)

	e.GET("/save",controller.Save)
	e.GET("/ws", controller.Hello)
	e.GET("/select",controller.SelectPersion)
	e.GET("/query",controller.Query)
	e.GET("/find",controller.Find)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

