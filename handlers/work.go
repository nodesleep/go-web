package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func workPage(c echo.Context) error {
    content := loadTemplate("templates/work.html")
    return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
        "title": "Work",
        "content": template.HTML(content),
    })
}
