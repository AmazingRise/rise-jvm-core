package jvm

import "github.com/AmazingRise/rise-jvm-core/logger"

type Stack interface {
	Pop() interface{}
	Push(value interface{})
	Peek() interface{}
	Fill(args ...interface{})
}

type LinkedListStack struct {
	Stack
	curr *Node
}

type Node struct {
	Prev  *Node
	Value interface{}
}

func CreateLinkedListStack(capacity int) *LinkedListStack {
	s := &LinkedListStack{}
	s.curr = &Node{}
	return s
}

func (s *LinkedListStack) Peek() interface{} {
	return s.curr.Value
}

func (s *LinkedListStack) Push(value interface{}) {
	s.curr = &Node{s.curr, value}
}

func (s *LinkedListStack) Pop() interface{} {
	val := s.curr.Value
	s.curr = s.curr.Prev
	return val
}

func (s *LinkedListStack) Fill(args ...interface{}) {
	for _, arg := range args {
		s.Push(arg)
	}
}

type PrimStack struct {
	Stack
	data []interface{}
}

func CreatePrimStack(capacity int) *PrimStack {
	return &PrimStack{data: make([]interface{}, 0)}
}

func (s *PrimStack) Peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *PrimStack) Push(value interface{}) {
	s.data = append(s.data, value)
}

func (s *PrimStack) Pop() interface{} {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *PrimStack) Fill(args ...interface{}) {
	s.data = append(s.data, args...)
}

type ArrStack struct {
	Stack
	data []interface{}
}

func CreateArrStack(capacity int) *ArrStack {
	return &ArrStack{data: make([]interface{}, 0, capacity)}
}

func (s *ArrStack) Peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *ArrStack) Push(value interface{}) {
	s.data = append(s.data, value)
	logger.Infoln("Pushed", value)
}

func (s *ArrStack) Pop() interface{} {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	logger.Infoln("Popped", val)
	return val
}

func (s *ArrStack) Fill(args ...interface{}) {
	s.data = append(s.data, args...)
}

func Serialize(stack Stack) {

}
