package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func aboutPage(c echo.Context) error {
    content := loadTemplate("templates/about.html")
    return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
        "title": "About Me",
        "content": template.HTML(content),
    })
}
