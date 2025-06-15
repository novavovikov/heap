# Generic Heap Library for Go

## Features

- **Generics:** Works with any type T
- **Type Safety:** No interface{} or type assertions
- **Custom Comparator:** Provide your own cmp(a, b T) bool to define ordering
- **Core Operations:** Push, Pop, PeekAt, Len in O(log n) or O(1)

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourname/heap"
)

func main() {
    h := heap.New(func(a, b int) bool { return a < b })
    for _, v := range []int{5, 2, 9, 1, 7} {
        h.Push(v)
    }

    for h.Len() > 0 {
        v, _ := h.Pop()
        fmt.Println(v)
    }
}
```

## API

```go
type Heap[T any] struct { ... }

func New[T any](cmp func(a, b T) bool) *Heap[T]

func (h *Heap[T]) Len() int
func (h *Heap[T]) PeekAt(idx int) (T, bool)
func (h *Heap[T]) Push(v T)
func (h *Heap[T]) Pop() (T, bool)
```
