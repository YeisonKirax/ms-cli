package example

import (
	"example/src/config"
	"example/src/modules/example/application/get_examples"
	"example/src/modules/example/domain/services"
	"example/src/modules/example/infrastructure/db/repositories"
	"example/src/modules/example/infrastructure/handlers"
	"example/src/modules/example/infrastructure/routes"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ExampleModule struct {
	environment config.Environment
}

var instance *ExampleModule
var once sync.Once

func NewExampleModule(environment config.Environment) *ExampleModule {
	logrus.Info("Starting module ExampleModule")
	once.Do(
		func() {
			instance = &ExampleModule{
				environment: environment,
			}
		},
	)
	return instance
}

func (m *ExampleModule) Init(app fiber.Router) {
	examplesRepository := repositories.NewExamplesRepositoryImpl()
	examplesService := services.NewExamplesServiceImpl(examplesRepository)
	getExamplesUseCase := get_examples.NewGetExamplesUseCase(examplesService)
	examplesHandlers := handlers.NewExamplesHandler(getExamplesUseCase)
	examplesRouter := routes.NewExamplesRouter(examplesHandlers)
	examplesRouter.LoadRoutes(app)
	logrus.Info("Module ExampleModule started successfully")
}
