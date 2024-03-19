package tests

import (
	"gostructs"
	"testing"
)

func heapClean(heap *gostructs.Bheap) {
	heaplen := len(heap.HeapItems())
	for i := 0; i <= heaplen; i++ {
		heap.DelMin()
	}
}

func TestDelAllItems(t *testing.T) {
	heap := gostructs.InitHeap([]int{1, 9, 7, 4, 2, 3})

	heapClean(&heap)

	if len(heap.HeapItems()) != 0 {
		t.Errorf("error TestDelAllItems got %v ", heap.HeapItems())
	}

	if heap.HeapLen() != -1 {
		t.Errorf("error TestDelAllItems got %d ", heap.HeapLen())
	}
}

func TestWithEmptyHeap(t *testing.T) {
	heap := gostructs.InitHeap([]int{1, 9, 7, 4, 2, 3})
	heapClean(&heap)
	heap.Insert(10)

	// empty heaplen is  -1
	if heap.HeapLen() != 0 {
		t.Errorf("Error heap insert got %v", heap.HeapItems())
	}

	if heap.HeapItems()[0] != 10 {
		t.Errorf("Error heap insert got %v", heap.HeapItems())
	}
	heap.Insert(2)
	if heap.HeapItems()[0] != 2 {
		t.Errorf("Error heap insert got %v", heap.HeapItems())
	}
}
