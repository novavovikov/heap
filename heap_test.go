package heap

import (
	"slices"
	"testing"
)

func TestHeap(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 9, 1, 7}, []int{1, 2, 5, 7, 9}},
		{[]int{}, []int{}},
		{[]int{42}, []int{42}},
	}

	for _, c := range cases {
		h := New(func(a, b int) bool { return a < b })

		for _, v := range c.input {
			h.Push(v)
		}

		var got []int
		for h.Len() > 0 {
			v, _ := h.Pop()
			got = append(got, v)
		}

		if !slices.Equal(got, c.want) {
			t.Errorf("for input %v: got %v, want %v", c.input, got, c.want)
		}
	}
}
