package handler

import (
	"github.com/josephschec/go-htmx/model"
	"github.com/josephschec/go-htmx/view/note"
	"github.com/labstack/echo/v4"
)

func HandlClickShow(c echo.Context) error {

	value := c.FormValue("note")
	noteVal := model.Note{
		Message: value,
	}
	if len(value) == 0 {
		return nil
	}
	return render(c, note.Show(noteVal))
}
