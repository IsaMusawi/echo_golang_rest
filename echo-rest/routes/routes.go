package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"restgolang.com/restGolang/echo-rest/controllers"
)

func Init() *echo.Echo {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

	e.GET("/pegawai", controllers.FetchPegawai)

	e.POST("/pegawai", controllers.StorePegawai)

	e.PUT("/pegawai", controllers.UpdatePegawai)
    
	return e
}

