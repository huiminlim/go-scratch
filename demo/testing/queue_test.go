// Use command `go test -v ./demo/testing` to run test and show all tests ran
package queue

import "testing"

func TestAddQueue(t *testing.T) { // Required function signature for test
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Should only add at most 3 items")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Failed to extract items from queue")
		}
		if item != i {
			t.Errorf("FIFO not working")
		}
	}
	_, ok := q.Next()
	if ok {
		t.Errorf("Should not have anymore items")
	}
}
