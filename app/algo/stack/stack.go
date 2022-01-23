package main

import (
	"errors"
	"fmt"
)

type Element interface{}

type Stack struct {
	top  int
	size int
	Data []Element
}

func (s *Stack) New(size int) *Stack {
	s.top = -1
	s.size = size
	s.Data = make([]Element, size)
	return s
}

func (s *Stack) Push(e Element) error {

	if s.IsFull() {
		return errors.New("Stack overflow")
	}

	s.top++
	s.Data[s.top] = e

	return nil
}

func (s *Stack) Pop() (Element, error) {

	if s.IsEmpty() {
		return nil, errors.New("Stack underflow")
	}

	lastElem := s.Data[s.top]
	s.top--

	return lastElem, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack) Peek() Element {
	return s.Data[s.top]
}

func doStack() {
	var data []Element = []Element{"Aqua", "Beast", "Dusk"}
	var data2 []Element = []Element{"Reptile", "Bug", "Dawn"}

	var stack Stack

	stack.New(3)

	checkError(stack.Push(data[0]))
	checkError(stack.Push(data[1]))
	checkError(stack.Push(data[2]))

	element, err := stack.Pop()
	checkError(err)
	fmt.Println("Popped Element:", element)

	fmt.Println("Top Element:", stack.Peek())

	checkError(stack.Push(data2[0]))
	checkError(stack.Push(data2[1]))

	fmt.Println(stack.Peek())
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	doStack()
}
