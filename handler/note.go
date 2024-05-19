package handler

import (
	"github.com/josephschec/go-htmx/model"
	"github.com/josephschec/go-htmx/view/note"
	"github.com/labstack/echo/v4"
)

func HandlNoteShow(c echo.Context) error {
	noteMessage := model.Note{
		Message: "",
	}
	return render(c, note.Show(noteMessage))
}
