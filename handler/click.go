package handler

import (
	"github.com/josephschec/go-htmx/view/clicked"
	"github.com/labstack/echo/v4"
)

type ClickHandler struct {}

func (handler ClickHandler) HandlClickShow(c echo.Context) error {
	return render(c, clicked.Show())
}
