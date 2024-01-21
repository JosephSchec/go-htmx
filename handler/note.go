package handler

import (
	"github.com/josephschec/go-htmx/model"
	"github.com/josephschec/go-htmx/view/note"
	"github.com/labstack/echo/v4"
)

type NoteHandler struct{}

func (handler NoteHandler) HandlNoteShow(c echo.Context) error {
	noteMessage := model.Note{
		Message: "",
	}
	return render(c, note.Show(noteMessage))
}
