package stack

import (
	"testing"
)

func TestDynamicSliceStack(t *testing.T) {
	st := NewDynamicSliceStack()

	if !st.IsEmpty() {
		t.Errorf("expected IsEmpty to be true, but got=%t", st.IsEmpty())
	}
	_, err := st.Peek()
	if err == nil {
		t.Errorf("expected error, but got=%v", err)
	}
	st.Push(1)

	if st.Size() != 1 {
		t.Errorf("expected Size=1, but got=%d", st.Size())
	}

	st.Push(2)

	if st.Size() != 2 {
		t.Errorf("expected Size=2, but got=%d", st.Size())
	}

	value, err := st.Peek()

	if err != nil {
		t.Errorf("expected error to be nil, but got %v", err)
	}

	if value != 2 {
		t.Errorf("expected Peek=2, but got=%d", value)
	}

	value, _ = st.Pop()

	if value != 2 {
		t.Errorf("expected Pop=2, but got=%d", value)
	}

	value, _ = st.Pop()

	if value != 1 {
		t.Errorf("expected Pop=1, but got=%d", value)
	}

	value, err = st.Pop()

	if err == nil {
		t.Errorf("expected error, but got=%v", err)
	}

	if value != -1 {
		t.Errorf("expected Pop=-1, but got=%d", value)
	}

	if st.Size() != 0 {
		t.Errorf("expected Size=0, but got=%d", st.Size())
	}
}
