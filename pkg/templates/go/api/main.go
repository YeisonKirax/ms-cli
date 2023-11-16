package main

import (
	"example/src/config"
	"example/src/modules/example"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(
		logger.New(
			logger.Config{
				TimeFormat: time.RFC822,
				TimeZone:   "America/Santiago",
			},
		),
	)
	env := config.Load()
	config.InitLogrus(env.Env)
	// Rutas de usuarios agregada
	examplesModule := example.NewExampleModule(env)
	examplesModule.Init(app)

	port := env.Port
	log.Fatalln(app.Listen(port))
}
