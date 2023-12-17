package gotimeseries

import (
	"fmt"
	"time"
)

func Example() {
	const buckets = 10

	now := time.Now()

	// construct a new time series that tracks 10 minutes
	series := New(time.Minute, buckets, now)

	// ... time passes ...
	now = now.Add(20 * time.Second)

	// update the time series to account for the time that has passed
	series.Update(now)

	// track an event
	series.Increase()

	// ... time passes ...
	now = now.Add(5*time.Minute + 40*time.Second)

	// update the time series to account for the time that has passed
	series.Update(now)

	// track an event
	series.Increase()

	// ... time passes ...
	now = now.Add(35 * time.Second)

	// update the time series to account for the time that has passed
	series.Update(now)

	total := series.Total()
	fmt.Printf("events in the last %dm: %d\n", buckets, total)
	fmt.Printf("events/min: %.2f\n", float64(total)/float64(buckets))

	// Output:
	// events in the last 10m: 2
	// events/min: 0.20
}
