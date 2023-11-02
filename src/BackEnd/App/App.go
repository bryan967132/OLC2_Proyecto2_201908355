package main

import (
	controller "TSwift/Controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	files := controller.Files{}
	interpreter := controller.Controller{}

	app.Get("/", interpreter.Running)
	app.Post("/interpreter/parser", interpreter.Parser)
	app.Post("/interpreter/getAST", interpreter.GetAST)
	app.Get("/interpreter/getSymbolsTable", interpreter.GetSymbolsTable)
	app.Get("/interpreter/getErrors", interpreter.GetErrors)
	app.Post("/files/openFile", files.OpenFile)
	app.Get("/files/openedFiles", files.OpenedFiles)
	app.Post("/files/setCurrentN", files.SetCurrentN)
	app.Get("/files/getCurrentN", files.GetCurrentN)
	app.Post("/files/close", files.Close)
	app.Post("/files/save", files.Save)
	app.Listen(":9000")
}
