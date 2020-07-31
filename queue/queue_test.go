package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSyncQueue(t *testing.T) {
	q := NewSyncQueue()
	if !assert.NotNil(t, q) {
		t.FailNow()
	}
	if !assert.False(t, q.closed) {
		t.FailNow()
	}
}

func TestSyncQueue(t *testing.T) {
	q := NewSyncQueue()
	go func() {
		for i := 0; i <= 10; i++ {
			q.Push(i)
		}
	}()
	for {
		v, ok := q.TryPop()
		if !ok {
			continue
		}
		if v.(int) > 10 || v.(int) < 0 {
			t.Fatal(v)
		}
		if v.(int) == 10 {
			return
		}
	}
}
