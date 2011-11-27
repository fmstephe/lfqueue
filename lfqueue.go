package lfqueue

import (
	"unsafe"
	"sync/atomic"
)

type node struct {
	val interface{}
	nxt *node
}

type Q struct {
	head *node
	tail *node
}

func New() (q *Q) {
	q = new(Q)
	n := new(node)
	q.head = n
	q.tail = n
	return
}

func (q *Q) enq(val interface{}) {
	var t, n *node
	n = &node{val: val, nxt: nil}
	for {
		t = q.tail
		if t.nxt != nil {
			tp := unsafe.Pointer(q.tail)
			atomic.CompareAndSwapPointer(&tp, unsafe.Pointer(t), unsafe.Pointer(t.nxt))
			continue
		}
		tp := unsafe.Pointer(q.tail)
		if atomic.CompareAndSwapPointer(&tp, nil, unsafe.Pointer(n)) {
			break
		}
	}
	tp := unsafe.Pointer(q.tail)
	atomic.CompareAndSwapPointer(&tp, unsafe.Pointer(t), unsafe.Pointer(n))
}

func (q *Q) deq() (val interface{}, success bool) {
	var h, t, n *node
	for {
		h = q.head
		t = q.tail
		n = h.nxt
		if h == t {
			if n == nil {
				return nil, false
			} else {
				tp := unsafe.Pointer(q.tail)
				atomic.CompareAndSwapPointer(&tp, unsafe.Pointer(t), unsafe.Pointer(n))
			}
		} else {
			val = n.val
			tp := unsafe.Pointer(q.head)
			if atomic.CompareAndSwapPointer(&tp, unsafe.Pointer(h), unsafe.Pointer(n)) {
				return val, true
			}
		}
	}
	panic("Unreachable")
}
