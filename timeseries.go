package gotimeseries

import "time"

// TimeSeries tracks the number of events that occur in a given time period.
type TimeSeries struct {
	// granularity is the time period that each bucket represents.
	granularity time.Duration

	// firstBucketTime is the time that the first bucket represents.
	// Specifically, the first bucket represents the time period firstBucketTime-granularity < x <= firstBucketTime.
	firstBucketTime time.Time

	// buckets is a ring buffer that stores the number of events that occurred in time periods of granularity.
	// The first bucket represents the latest of those periods, while the last bucket represents the oldest.
	buckets *ringBuffer
}

// New creates a new TimeSeries with the given granularity and number of buckets.
func New(granularity time.Duration, buckets int, now time.Time) *TimeSeries {
	return &TimeSeries{
		granularity:     granularity,
		firstBucketTime: now,
		buckets:         newRingBuffer(buckets),
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

	sliceShifts := min(numShifts, t.buckets.size())
	for i := 0; i < sliceShifts; i++ {
		t.buckets.turnBack()
		t.buckets.putFirst(0)
	}
}

// Increase increments the count of events in the current time period.
func (t *TimeSeries) Increase() {
	t.buckets.putFirst(t.buckets.first() + 1)
}

// Total returns the total number of events that have occurred in the time series.
func (t *TimeSeries) Total() uint64 {
	return t.buckets.total()
}
