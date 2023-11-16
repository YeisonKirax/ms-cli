package repositories

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type ExamplesRepositoryImpl struct {
}

var once sync.Once
var instance *ExamplesRepositoryImpl

func NewExamplesRepositoryImpl() *ExamplesRepositoryImpl {
	logrus.Info("Starting repósitory ExamplesRepositoryImpl")
	once.Do(
		func() {
			instance = &ExamplesRepositoryImpl{}
		},
	)
	return instance
}
func (ex *ExamplesRepositoryImpl) Get() []string {
	return []string{"example1", "example2"}
}
