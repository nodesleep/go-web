package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func homePage(c echo.Context) error {
    content := loadTemplate("templates/home.html")
    return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
        "title": "Home",
        "content": template.HTML(content),
    })
}
