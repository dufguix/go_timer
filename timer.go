package timer

import "time"

type Timer struct {
	initTime    time.Time
	elapsedTime time.Duration
}

func (t *Timer) Reset() {
	t.initTime = time.Now()
	t.elapsedTime = 0
}

// Refreshes elapsed time and Returns true if the time is elapsed.
func (t *Timer) RefreshAndCheck(duration time.Duration) bool {
	t.Refresh()
	return t.Check(duration)
}

// Refreshes elapsed time
func (t *Timer) Refresh() {
	t.elapsedTime = time.Since(t.initTime)
}

// Returns true if the time is elapsed.
func (t *Timer) Check(duration time.Duration) bool {
	return t.elapsedTime > duration
}

func (t *Timer) GetElapsedTime() time.Duration {
	return t.elapsedTime
}
