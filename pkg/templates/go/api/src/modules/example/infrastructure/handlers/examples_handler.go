package handlers

import (
	"example/src/modules/shared/domain/interfaces"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type ExamplesHandler struct {
	getExamplesUseCase interfaces.UseCase[string, []string]
}

var instance *ExamplesHandler
var once sync.Once

func NewExamplesHandler(getExamplesUseCase interfaces.UseCase[string, []string]) *ExamplesHandler {
	once.Do(
		func() {
			instance = &ExamplesHandler{
				getExamplesUseCase: getExamplesUseCase,
			}
		},
	)
	return instance
}

func (eh *ExamplesHandler) GetExamples(ctx *fiber.Ctx) error {
	examples, err := eh.getExamplesUseCase.Execute("example")
	if err != nil {
		errorRes := &fiber.Map{"status": false, "error": err.Error()}
		return ctx.JSON(errorRes)
	}
	return ctx.JSON(examples)
}
