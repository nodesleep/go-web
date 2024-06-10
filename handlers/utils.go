package handlers

import (
	"bytes"
	"html/template"
)

func loadTemplate(filename string) string {
    tmpl, err := template.ParseFiles(filename)
    if err != nil {
        return ""
    }
    var buf bytes.Buffer
    err = tmpl.ExecuteTemplate(&buf, "content", nil)
    if err != nil {
        return ""
    }
    return buf.String()
}
