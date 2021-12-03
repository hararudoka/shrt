package main

import (
	"html/template"
	"io"
	"shorter/config"
	"shorter/handler"
	"shorter/storage"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	db, err := storage.Open()
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(handler.Handler{DB: db})

	e := echo.New()
	e.Renderer = t

	h.REGISTER(*e.Group(""), &handler.ShortsStorage{})

	c := config.Config{}
	err = c.LoadData()
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":"+c.Port))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
