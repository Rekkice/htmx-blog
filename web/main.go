package main

import (
	_ "html/template"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func addRoutes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Renderer = Renderer
		e.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Set("dao", app.Dao())
				c.Set("app", app)
				return next(c)
			}
		})

		e.Router.Static("/static", "ui/static")

		e.Router.GET("/", HomeHandler, apis.ActivityLogger(app))

		return nil
	})
}

func main() {
	app := pocketbase.New()

	addRoutes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.HandleFunc("/", home)
}
