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

func TestRand(t *testing.T) {
	input := make([]int, 100, 100)
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
