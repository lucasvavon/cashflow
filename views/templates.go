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
	userID := c.Get("user_id")
	isAuthenticated := userID != nil

	templateData := map[string]interface{}{
		"IsAuthenticated": isAuthenticated,
		"UserID":          userID,
	}

	if dataMap, ok := data.(map[string]interface{}); ok {
		for k, v := range dataMap {
			templateData[k] = v
		}
	} else if data != nil {
		templateData["Data"] = data
	}

	return t.templates.ExecuteTemplate(w, name, templateData)
}

func NewTemplate() (*Templates, error) {
	tmpl, err := template.ParseGlob("views/*.html")
	if err != nil {
		return nil, fmt.Errorf("error loading templates: %w", err)
	}
	return &Templates{templates: tmpl}, nil
}
