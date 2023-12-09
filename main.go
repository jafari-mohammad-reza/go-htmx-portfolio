package main

import (
	"fmt"
	"jafari-mohammad-reza/portfolio/pkg"
	"jafari-mohammad-reza/portfolio/pkg/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	
	e.Use(middleware.Recover())
	pkg.NewTemplateRenderer(e, "views/*.html")
	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index.html", map[string]string{
			"Title": "Htmx Echo practice.",
		})
	})
	config.SetupConfigs()
	privateConfig, _ := config.GetConfigs()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", privateConfig.Port)))
}
