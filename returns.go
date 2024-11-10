package returns

import "fmt"

// Result type is the combination of value and error
type Result[T any] struct {
	value T
	err   error
}

// Check if the result is an ok result. Ok result means there has no error
func (result Result[T]) IsOk() bool {
	return result.err == nil
}

// Check if the result is an Err result. Err result contains only error
func (result Result[T]) IsErr() bool {
	return result.err != nil
}

// Value return the value when the result is Ok; otherwise, panic
func (result Result[T]) Value() T {
	if result.IsErr() {
		panic(fmt.Sprintf("Err result has not values. err: %s", result.err))
	}
	return result.value
}

// Value returns the value when the result is Ok; otherwise, return the defaultValue
func (result Result[T]) ValueOr(defaultValue T) T {
	if result.IsOk() {
		return result.value
	} else {
		return defaultValue
	}
}

// Create a new Ok result and return
func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

// Create a new Err result and return
func Err(err error) Result[interface{}] {
	return Result[interface{}]{
		value: nil,
		err:   err,
	}
}
