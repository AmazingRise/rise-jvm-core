package jvm

import (
	"wasm-jvm/logger"
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
	ThreadDead    = 2
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
		curr := p.head.Next
		//logger.Infof("Thread %d is being executed.", curr.Id)
		if len(curr.FrameStack) == 0 { // empty thread, kill it
			// But here should not be reached
			logger.Infof("Thread %d exits, because nothing to do.", curr.Id)
			p.DeleteThread(curr.Id)
			continue
		}
		frameStack := curr.FrameStack
		// Execute the last frameStack
		result := frameStack[len(frameStack)-1].Exec()
		//logger.Infoln("Frame:", len(frameStack))
		switch frameStack[len(frameStack)-1].State {
		case FrameExit:
			if len(frameStack) > 1 {
				// Delete the empty frame
				curr.FrameStack = curr.FrameStack[:len(curr.FrameStack)-1]
				// Transfer the stack
				curr.FrameStack[len(curr.FrameStack)-1].Stack = result
				// Move to tail
				p.moveToTail(curr)
				logger.Infoln("Frame of thread", curr.Id, "exits, returns", curr.FrameStack[len(curr.FrameStack)-1].Stack)
				//curr.FrameStack[len(curr.FrameStack)-1].PC++
			} else {
				// Exit this thread
				logger.Infof("Thread %d exits normally.", curr.Id)
				p.DeleteThread(curr.Id)
			}
		case FramePush: // if current frame pushed a new frame
			// Get the old stack
			dataStack := frameStack[len(frameStack)-1].Stack
			// Restore the frame state
			frameStack[len(frameStack)-1].State = FrameReady
			// The result is method ref
			idx := result[0].(uint16)
			// Locate method
			class, name, _ := frameStack[len(frameStack)-1].This.Constants.GetMethodRef(idx)
			method := p.vm.LocateMethod(class, name)
			frame := p.vm.InvokeStaticMethod(method, dataStack...)
			logger.Infof("Frame of thread %d pushed a new frame named %s::%s.", curr.Id, class, name)
			// Empty the stack
			frameStack[len(frameStack)-1].Stack = []interface{}{}
			// Push the frame
			curr.FrameStack = append(curr.FrameStack, frame)
		case FramePatch:

		}

	}
}
