package gotimeseries

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestTimeSeries(t *testing.T) {
	events := []struct {
		at          time.Duration
		increase    bool
		wantBuckets []uint64
	}{
		{0 * time.Second, true, []uint64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{0 * time.Second, true, []uint64{2, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{5 * time.Second, true, []uint64{1, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{5 * time.Second, true, []uint64{2, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{59 * time.Second, true, []uint64{3, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{59 * time.Second, true, []uint64{4, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{60 * time.Second, true, []uint64{5, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{60 * time.Second, true, []uint64{6, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{61 * time.Second, true, []uint64{1, 6, 2, 0, 0, 0, 0, 0, 0, 0}},
		{61 * time.Second, true, []uint64{2, 6, 2, 0, 0, 0, 0, 0, 0, 0}},
		{320 * time.Second, true, []uint64{1, 0, 0, 0, 2, 6, 2, 0, 0, 0}},
		{320 * time.Second, true, []uint64{2, 0, 0, 0, 2, 6, 2, 0, 0, 0}},
		{1000 * time.Second, true, []uint64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{1000 * time.Second, true, []uint64{2, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}

	is := is.New(t)

	startDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	series := New(time.Minute, 10, startDate)

	for _, event := range events {
		series.Update(startDate.Add(event.at))

		if event.increase {
			series.Increase()
		}

		is.Equal(series.buckets.bucketsSlice(), event.wantBuckets)

		total := uint64(0)
		for _, b := range event.wantBuckets {
			total += b
		}

		is.Equal(series.Total(), total)
	}
}
