package gotimeseries

import "time"

// TimeSeries tracks the number of events that occur in a given time period.
type TimeSeries struct {
	// granularity is the time period that each bucket represents.
	granularity time.Duration

	// firstBucketTime is the time that the first bucket represents.
	// Specifically, the first bucket represents the time period firstBucketTime-granularity < x <= firstBucketTime.
	firstBucketTime time.Time

	// buckets is a ring buffer that stores the number of events that occurred in a time period of granularity.
	// The first bucket represents the latest of those periods, while the last bucket represents the oldest.
	buckets []uint64

	// firstBucketIndex is the index of the first bucket in buckets.
	firstBucketIndex int
}

// New creates a new TimeSeries with the given granularity and number of buckets.
func New(granularity time.Duration, buckets int, now time.Time) *TimeSeries {
	return &TimeSeries{
		granularity:     granularity,
		firstBucketTime: now,
		buckets:         make([]uint64, buckets),
	}
}

// Update accounts for the time that has passed since the last update.
// It should always be called before Increase or Total.
func (t *TimeSeries) Update(now time.Time) {
	elapsed := now.Sub(t.firstBucketTime)
	if elapsed <= 0 {
		return
	}

	numShifts := int(elapsed/t.granularity) + 1
	t.firstBucketTime = t.firstBucketTime.Add(time.Duration(numShifts) * t.granularity)

	sliceShifts := min(numShifts, len(t.buckets))
	for i := 0; i < sliceShifts; i++ {
		t.firstBucketIndex--
		if t.firstBucketIndex < 0 {
			t.firstBucketIndex = len(t.buckets) - 1
		}

		t.buckets[t.firstBucketIndex] = 0
	}
}

// Increase increments the count of events in the current time period.
func (t *TimeSeries) Increase() {
	t.buckets[t.firstBucketIndex]++
}

// Total returns the total number of events that have occurred in the time series.
func (t *TimeSeries) Total() uint64 {
	total := uint64(0)
	for _, b := range t.buckets {
		total += b
	}

	return total
}

func (t *TimeSeries) bucketsSlice() []uint64 {
	return append(t.buckets[t.firstBucketIndex:], t.buckets[:t.firstBucketIndex]...)
}
