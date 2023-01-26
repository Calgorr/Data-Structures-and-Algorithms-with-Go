package main

import "fmt"

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) push(value T) {
	stack.values = append(stack.values, value)
}
func (stack *Stack[T]) Top() (T, error) {
	var value T
	if stack.IsEmpty() {
		return value, fmt.Errorf("Stack is empty")
	}
	return stack.values[len(stack.values)-1], nil
}

func (stack *Stack[T]) IsEmpty() bool {
	if len(stack.values) == 0 {
		return true
	}
	return false
}

func (stack *Stack[T]) pop() (T, error) {
	var value T
	if stack.IsEmpty() {
		return value, fmt.Errorf("Stack is empty")
	}
	value, stack.values = stack.values[len(stack.values)-1], stack.values[:len(stack.values)-1]
	return value, nil
}
