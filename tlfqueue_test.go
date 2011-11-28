package lfqueue

import (
	"math/rand"
	"testing"
)

func TestEnqDeq(t *testing.T) {
	q := New()
	eval := "value"
	q.enq(eval)
	dval, _ := q.deq()
	if eval != dval {
		t.Errorf("Failed to enq and deq symmetrically expecting \"%s\" dequeued \"%s\"", eval, dval)
	}
}

func TestEnqDeqDeq(t *testing.T) {
	q := New()
	eval := "value"
	q.enq(eval)
	q.deq()
	dval, _ := q.deq()
	if dval != nil {
		t.Errorf("Failed, expecting \"nil\" dequeued \"%s\"", dval)
	}
}

func TestEnqEnqDeqDeq(t *testing.T) {
	q := New()
	eval1 := "valueI"
	eval2 := "valueII"
	q.enq(eval1)
	q.enq(eval2)
	dval1, _ := q.deq()
	dval2, _ := q.deq()
	if eval1 != dval1 || eval2 != dval2 {
		t.Errorf("Failed, expecting \"%s\" dequeued \"%s\", expecting \"%s\" dequeued \"%s\"", eval1, dval1, eval2, dval2)
	}
}

func testRand(t *testing.T) {
	input := make([]int, 1000, 1000)
	for i := range input {
		input[i] = rand.Int()
	}
	q := New()
	for i := range input {
		q.enq(input[i])
	}
	for i := range input {
		r, _ := q.deq()
		if r != input[i] {
			t.Errorf("Failed, expecting \"%v\" dequeued \"%v\"", input[i], r)
		}
		q.enq(input[i])
	}
}

func TestConc(t *testing.T) {
	inputs := 1000000
	grts := 8
	input := make([]int, inputs)
	termChans := make([]chan bool, grts)
	for i := range termChans {
		termChans[i] = make(chan bool)
	}
	for i := range input {
		input[i] = rand.Int()
	}
	q := New()
	for i := range termChans {
		go enqAndTerminate(q, input, termChans[i])
	}
	for i := range termChans {
		<-termChans[i]
	}
	results := make(map[int]int)
	for {
		if val, ok := q.deq(); ok {
			intVal := val.(int)
			if count, ok := results[intVal]; ok {
				results[intVal] = count + 1
			} else {
				results[intVal] = 1
			}
		} else {
			break
		}
	}
	for i := range input {
		val := input[i]
		if results[val] % grts != 0 {
			t.Error("Failed")
		}
	}
}

func enqAndTerminate(q *Q, input []int, termChan chan bool) {
	for _, v := range input {
		q.enq(v)
	}
	termChan<-true
}
