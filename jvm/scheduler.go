package jvm

import (
	"github.com/AmazingRise/rise-jvm-core/logger"
	"strings"
)

type ThreadPool struct {
	head      *Thread
	tail      *Thread
	ctr       int // Self increment ctr
	threadMap map[int]*Thread
	vm        *VM
	// TODO: Add a map of thread
}

const (
	ThreadReady   = 0
	ThreadBlocked = 1
	ThreadExit    = 2
)

type Thread struct {
	Id         int // reserved
	State      int
	FrameStack []*Frame
	Prev       *Thread
	Next       *Thread
}

func (v *VM) CreateThreadPool() *ThreadPool {
	pool := &ThreadPool{
		head:      &Thread{},
		tail:      &Thread{},
		ctr:       0,
		threadMap: make(map[int]*Thread),
		vm:        v,
	}
	pool.head.Next = pool.tail
	pool.tail.Prev = pool.head
	return pool
}

func (p *ThreadPool) CreateThread(frame *Frame) *Thread {
	p.ctr++
	thread := &Thread{
		Id:         p.ctr,
		FrameStack: []*Frame{frame},
	}
	p.addThreadAfter(p.head, thread)
	p.threadMap[p.ctr] = thread
	logger.Infof("Thread %d is created.", p.ctr)
	return thread
}

func (p *ThreadPool) GetThread(idx int) *Thread {
	return p.threadMap[idx]
}

func (p *ThreadPool) moveToTail(thread *Thread) {
	p.removeThread(thread)
	p.addThreadAfter(p.tail.Prev, thread)
}

// removeThread Remove the thread from linked list
func (p *ThreadPool) removeThread(thread *Thread) {
	prev := thread.Prev
	next := thread.Next
	prev.Next = next
	next.Prev = prev
}

func (p *ThreadPool) addThreadAfter(after *Thread, thread *Thread) {
	next := after.Next
	after.Next = thread
	next.Prev = thread
	thread.Prev = after
	thread.Next = next
}

// DeleteThread delete a thread permanently
func (p *ThreadPool) DeleteThread(idx int) {
	// TODO: Exception process
	thread := p.GetThread(idx)
	p.removeThread(thread)
	delete(p.threadMap, thread.Id)
}

// Schedule to schedule the threads in the pool
func (p *ThreadPool) Schedule() {
	for len(p.threadMap) != 0 { // if threads are still running
		thread := p.head.Next
		//logger.Infof("Thread %d is being executed.", thread.Id)
		//frameStack := thread.FrameStack
		frame := thread.FrameStack[len(thread.FrameStack)-1]
		//frame := frameStack[len(frameStack)-1]
		// Execute the last frameStack
		result := p.vm.Exec(frame)

		if thread.State == ThreadBlocked {
			p.moveToTail(thread)
			continue
		}

		switch frame.State {
		case FrameExit:
			if len(thread.FrameStack) > 1 {
				// Delete the empty frame
				thread.FrameStack = thread.FrameStack[:len(thread.FrameStack)-1]
				// Transfer the stack
				last := thread.FrameStack[len(thread.FrameStack)-1]
				last.Stack = append(last.Stack, result...)
				// Move to tail
				p.moveToTail(thread)
				logger.Infoln("Frame of thread", thread.Id, "exits, returns", result)
				//thread.FrameStack[len(thread.FrameStack)-1].PC++
			} else {
				// Exit this thread
				logger.Infof("Thread %d exits normally.", thread.Id)
				p.DeleteThread(thread.Id)
			}
		case FramePush: // if current frame pushed a new frame
			// Get the old stack
			//dataStack := frame.Stack
			// Restore the frame state
			frame.State = FrameReady
			// The result is method ref
			idx := result[0].(uint16)
			// Locate method
			class, name, desc := frame.This.Constants.GetMethodRef(idx)
			method := p.vm.LocateMethod(class, name, desc)

			paramCount := GetParamCount(desc)
			var params []interface{}
			if !method.IsStatic() {
				paramCount++
			}
			params = frame.Stack[len(frame.Stack)-paramCount:]
			frame.Stack = frame.Stack[:len(frame.Stack)-paramCount]
			var newFrame *Frame
			// If method's attr is nil, then it is a runtime method.
			if method.Attrs == nil {
				newFrame = p.vm.InvokeRuntimeMethod(method, params...)
			} else {
				newFrame = p.vm.InvokeMethod(method, params...)
			}
			logger.Infof("Frame of thread %d pushed a new frame named %s::%s.", thread.Id, class, name)

			// Push the frame
			thread.FrameStack = append(thread.FrameStack, newFrame)
		}

		p.moveToTail(thread)
	}
}

func GetParamCount(desc string) int {
	/*
		B	byte	signed byte
		C	char	Unicode character code point in the Basic Multilingual Plane, encoded with UTF-16
		D	double	double-precision floating-point value
		F	float	single-precision floating-point value
		I	int	integer
		J	long	long integer
		L ClassName ;	reference	an instance of class ClassName
		S	short	signed short
		Z	boolean	true or false
		[	reference	one array dimension
	*/
	desc = desc[1:strings.Index(desc, ")")]
	flag := false
	ctr := 0
	for _, c := range desc {
		if c == ';' {
			flag = false
			ctr++
			continue
		}
		if flag {
			continue
		}
		switch c {
		case 'B':
			fallthrough
		case 'C':
			fallthrough
		case 'D':
			fallthrough
		case 'F':
			fallthrough
		case 'I':
			fallthrough
		case 'J':
			fallthrough
		case 'S':
			fallthrough
		case 'Z':
			ctr++
		case 'L':
			flag = true
		}
	}
	return ctr
}
