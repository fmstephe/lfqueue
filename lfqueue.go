package lfqueue

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	val interface{}
	nxt unsafe.Pointer
}

type Q struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func New() (q *Q) {
	q = new(Q)
	n := unsafe.Pointer(new(node))
	q.head = n
	q.tail = n
	return
}

func (q *Q) enq(val interface{}) {
	var t, n unsafe.Pointer
	n = unsafe.Pointer(&node{val: val, nxt: nil})
	for {
		t = q.tail
		nxt := ((*node)(t)).nxt
		if nxt != nil {
			atomic.CompareAndSwapPointer(&q.tail, t, nxt)
		} else if atomic.CompareAndSwapPointer(&((*node)(t)).nxt, nil, n) {
			break
		}
	}
	atomic.CompareAndSwapPointer(&q.tail, t, n)
}

func (q *Q) deq() (val interface{}, success bool) {
	var h, t, n unsafe.Pointer
	for {
		h = q.head
		t = q.tail
		n = ((*node)(h)).nxt
		if h == t {
			if n == nil {
				return nil, false
			} else {
				atomic.CompareAndSwapPointer(&q.tail, t, n)
			}
		} else {
			val = ((*node)(n)).val // Enq(...) write to val may not be visible
			if atomic.CompareAndSwapPointer(&q.head, h, n) {
				return val, true
			}
		}
	}
	panic("Unreachable")
}
