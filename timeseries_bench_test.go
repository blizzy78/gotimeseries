package gotimeseries

import (
	"testing"
	"time"
)

func BenchmarkTimeSeries(b *testing.B) {
	now := time.Now()

	series := New(time.Minute, 120, now)

	for i := 0; i < b.N; i++ {
		now = now.Add(time.Minute)

		series.Update(now)
	}
}
