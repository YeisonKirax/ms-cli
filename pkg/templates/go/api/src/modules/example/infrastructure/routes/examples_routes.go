package routes

import (
	"example/src/modules/example/infrastructure/handlers"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type ExamplesRouter struct {
	examplesHandlers *handlers.ExamplesHandler
}

var instance *ExamplesRouter
var once sync.Once

func NewExamplesRouter(examplesHandlers *handlers.ExamplesHandler) *ExamplesRouter {
	once.Do(
		func() {
			instance = &ExamplesRouter{
				examplesHandlers: examplesHandlers,
			}
		},
	)
	return instance
}

func (er *ExamplesRouter) LoadRoutes(app fiber.Router) {
	examplesRoute := app.Group("/examples")
	examplesRoute.Get("", er.examplesHandlers.GetExamples)
}
