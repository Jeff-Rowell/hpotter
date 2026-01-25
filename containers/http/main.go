package main

import (
	"html/template"
	"log"

	"github.com/hpotter/containers/http/pkg/endpoints"
	"github.com/hpotter/containers/http/pkg/endpoints/auth"

	"github.com/labstack/echo/v4"
)

func main() {
	var err error
	tmpl := template.New("")

	tmpl, err = tmpl.ParseGlob("./public/views/pages/*.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.ParseGlob("./public/views/pages/auth/*.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.ParseGlob("./public/views/layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Renderer = endpoints.NewTemplateRenderer(tmpl)
	e.Use(endpoints.NewLoggerMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	e.POST("/", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	e.PUT("/", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	e.DELETE("/", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	e.Static("/static", "public/static")

	publicGroup := e.Group("/api/v1")
	publicGroup.GET("/auth/login", auth.HandleLoginForm)
	publicGroup.POST("/auth/login", auth.HandleLocalLogin)

	e.Logger.Fatal(e.Start(":80"))
}
