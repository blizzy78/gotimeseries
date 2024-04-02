package gotimeseries

// ringBuffer is a ring buffer of elements.
type ringBuffer struct {
	// buckets stores the elements in the ring buffer.
	buckets []uint64

	// firstIndex is the index of the "first" bucket.
	firstIndex int
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		buckets: make([]uint64, size),
	}
}

func (rb *ringBuffer) size() int {
	return len(rb.buckets)
}

func (rb *ringBuffer) first() uint64 {
	return rb.buckets[rb.firstIndex]
}

func (rb *ringBuffer) putFirst(elem uint64) {
	rb.buckets[rb.firstIndex] = elem
}

func (rb *ringBuffer) turnBack() {
	rb.firstIndex--
	if rb.firstIndex < 0 {
		rb.firstIndex = len(rb.buckets) - 1
	}
}

func (rb *ringBuffer) total() uint64 {
	total := uint64(0)
	for _, b := range rb.buckets {
		total += b
	}

	return total
}
