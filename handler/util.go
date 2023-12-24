package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, templfile templ.Component) error {
	return templfile.Render(ctx.Request().Context(), ctx.Response())
}
