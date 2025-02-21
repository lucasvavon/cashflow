package views

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() (*Templates, error) {
	tmpl, err := template.ParseGlob("views/*.html")
	if err != nil {
		return nil, fmt.Errorf("error loading templates: %w", err)
	}
	return &Templates{templates: tmpl}, nil
}
