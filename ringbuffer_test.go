package gotimeseries

func (rb *ringBuffer) bucketsSlice() []uint64 {
	return append(rb.buckets[rb.firstIndex:], rb.buckets[:rb.firstIndex]...)
}
