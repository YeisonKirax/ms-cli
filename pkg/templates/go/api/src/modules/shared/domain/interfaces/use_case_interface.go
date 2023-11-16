package interfaces

type UseCase[T any, U any] interface {
	Execute(data T) (U, error)
}
