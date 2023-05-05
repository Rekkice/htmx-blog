package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v5"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found: " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base", data)
}

var templates = map[string]*template.Template{
	"home.go.html": template.Must(template.ParseFiles("ui/html/pages/home.go.html", "ui/html/base.go.html")),
}

var Renderer = &TemplateRegistry{
	templates: templates,
}

func HomeHandler(c echo.Context) error {
	println("HomeHandler called!")

	return c.Render(200, "home.go.html", nil)
}
