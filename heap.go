package heap

type Heap[T any] struct {
	items []T
	cmp   func(a, b T) bool
}

func New[T any](cmp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		items: make([]T, 0),
		cmp:   cmp,
	}
}

func (h *Heap[T]) Len() int {
	return len(h.items)
}

func (h *Heap[T]) PeekAt(idx int) (T, bool) {
	if idx < 0 || idx >= h.Len() {
		var zero T
		return zero, false
	}

	return h.items[idx], true
}

func (h *Heap[T]) Push(v T) {
	h.items = append(h.items, v)
	h.siftUp(h.Len() - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	top, ok := h.PeekAt(0)
	if !ok {
		return top, false
	}

	h.swap(0, h.Len()-1)
	h.items = h.items[:h.Len()-1]
	h.siftDown(0)

	return top, true
}

func (h *Heap[T]) siftUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2

		if h.cmp(h.items[idx], h.items[parent]) {
			h.swap(idx, parent)
			idx = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) siftDown(idx int) {
	n := h.Len()
	for {
		left := idx*2 + 1
		right := idx*2 + 2
		smallest := idx

		if left < n && h.cmp(h.items[left], h.items[smallest]) {
			smallest = left
		}

		if right < n && h.cmp(h.items[right], h.items[smallest]) {
			smallest = right
		}

		if smallest == idx {
			return
		}

		h.swap(idx, smallest)
		idx = smallest
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.items[j], h.items[i] = h.items[i], h.items[j]
}
