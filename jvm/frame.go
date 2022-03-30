package jvm

import (
	"rise-jvm-core/entity"
)

type Frame struct {
	entity.ByteCode
	PC     uint32
	Stack  []interface{}
	Locals []interface{}
	//Ins    [][]byte

	MethodRef *entity.Method
	State     int

	This *entity.Class
}

const (
	FrameReady = 0
	FrameExit  = 1
	FramePush  = 2
)
