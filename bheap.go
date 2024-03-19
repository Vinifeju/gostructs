package gostructs

import "fmt"

type Bheap struct {
	heaplen   int
	heapitems []int
}

func (b *Bheap) HeapLen() int {
	return b.heaplen
}

func (b *Bheap) HeapItems() []int {
	return b.heapitems
}

func InitHeap(nums []int) Bheap {
	return Bheap{heaplen: len(nums) - 1, heapitems: nums}
}

func (b *Bheap) percDown(i int) {
	for i*2 < b.heaplen {
		tmp := b.heapitems[i]
		mc := b.heapitems[b.minChild(i)]

		if b.heapitems[i] > mc {
			b.heapitems[i] = mc
			b.heapitems[b.minChild(i)] = tmp
		}
		i = b.minChild(i)

	}
}

func (b *Bheap) percUp(i int) {
	for i != 0 {
		if b.heapitems[i] < b.heapitems[i/2] {
			b.heapitems[i], b.heapitems[i/2] = b.heapitems[i/2], b.heapitems[i]
		}
		i /= 2
	}
}

func (b *Bheap) minChild(i int) int {
	lc, rc := b.leftChild(i), b.rightChild(i)
	if b.heapitems[lc] < b.heapitems[rc] {
		return lc
	}
	return rc
}

func (b *Bheap) leftChild(i int) int {
	lc := (2 * i) + 1
	if lc > b.heaplen {
		return b.heaplen
	}
	return lc
}

func (b *Bheap) rightChild(i int) int {
	rc := (2 * i) + 2
	if rc > b.heaplen {
		return b.heaplen
	}
	return rc
}

func (b *Bheap) BuildHeap() {
	i := (b.heaplen) / 2

	for i >= 0 {
		b.percDown(i)
		i--
	}
}

func (b *Bheap) DelMin() (int, error) {
	if len(b.heapitems) == 0 {
		return -1, fmt.Errorf("del from empty heap")
	}
	firstElem := b.heapitems[0]
	b.heapitems[0], b.heapitems[b.heaplen] = b.heapitems[b.heaplen], b.heapitems[0]
	b.heapitems = b.heapitems[:b.heaplen]
	b.heaplen--
	b.percDown(0)
	return firstElem, nil
}

func (b *Bheap) Insert(i int) {
	b.heapitems = append(b.heapitems, i)
	b.heaplen++
	b.percUp(b.heaplen)
}
