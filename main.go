package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed public
var frontend embed.FS

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Use("/normal", filesystem.New(filesystem.Config{
		Root:  http.Dir("./public"),
		Index: "index.html",
	}))

	app.Use("/embed", filesystem.New(filesystem.Config{
		Root:       http.FS(frontend),
		Index:      "index.html",
		PathPrefix: "/public",
	}))

	log.Fatal(app.Listen("localhost:3000"))
}
