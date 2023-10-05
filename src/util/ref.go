package util

func ToRef[T any](input T) *T {
	return &input
}
