package endpoints

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	template *template.Template
}

func NewTemplateRenderer(tmpl *template.Template) TemplateRenderer {
	return TemplateRenderer{template: tmpl}
}

func (t TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}
