package get_examples

import (
	"example/src/modules/example/domain/interfaces"
	"sync"

	"github.com/sirupsen/logrus"
)

type GetExamplesUseCase struct {
	examplesService interfaces.ExampleService
}

var instance *GetExamplesUseCase
var once sync.Once

func NewGetExamplesUseCase(examplesService interfaces.ExampleService) *GetExamplesUseCase {
	logrus.Info("Starting use case GetExamplesUseCase")
	once.Do(
		func() {
			instance = &GetExamplesUseCase{
				examplesService: examplesService,
			}
		},
	)
	return instance
}

func (g *GetExamplesUseCase) Execute(data string) ([]string, error) {
	return g.examplesService.GetExamples()
}
