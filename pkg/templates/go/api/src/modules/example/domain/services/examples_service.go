package services

import (
	"example/src/modules/example/domain/interfaces"
	"sync"

	"github.com/sirupsen/logrus"
)

type ExamplesServiceImpl struct {
	examplesRepository interfaces.ExampleRepository
}

var once sync.Once
var instance *ExamplesServiceImpl

func NewExamplesServiceImpl(
	examplesRepository interfaces.ExampleRepository,
) *ExamplesServiceImpl {
	logrus.Info("Starting service ExamplesServiceImpl")
	once.Do(
		func() {
			instance = &ExamplesServiceImpl{
				examplesRepository: examplesRepository,
			}
		},
	)
	return instance
}
func (ex *ExamplesServiceImpl) GetExamples() ([]string, error) {
	return ex.examplesRepository.Get(), nil
}
