package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/josephschec/go-htmx/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	envFile, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal(err)
	}
	port := envFile["PORT"]
	if port == "" {
		log.Fatal("MISSINGPORT", port)
	}

	portString := fmt.Sprintf(":%s", port)

	app := echo.New()
	noteHandle := handler.NoteHandler{}
	// app.Use(withNote)
	app.GET("/", noteHandle.HandlNoteShow)
	clickHandle := handler.ClickHandler{}
	app.POST("/click", clickHandle.HandlClickShow)
	app.Start(portString)
}

// func withNote(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := context.WithValue(c.Request().Context(), "note", "hello world ctx")
// 		c.SetRequest(c.Request().WithContext(ctx))
// 		return next(c)
// 	}
// }
