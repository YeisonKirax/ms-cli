package interfaces

type ExampleService interface {
	GetExamples() ([]string, error)
}
