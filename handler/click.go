package handler

import (
	"github.com/josephschec/go-htmx/model"
	"github.com/josephschec/go-htmx/view/note"
	"github.com/labstack/echo/v4"
)

type ClickHandler struct{}

func (handler ClickHandler) HandlClickShow(c echo.Context) error {

	value := c.FormValue("note")
	noteVal := model.Note{
		Message: value,
	}
	return render(c, note.Show(noteVal))
}
