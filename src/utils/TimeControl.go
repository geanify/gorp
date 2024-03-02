package utils

import "time"

const TickRateMS = 100 //miliseconds

type TimeControl struct {
	start time.Time
}

func (timeControl *TimeControl) GetElapsed() time.Duration {
	now := time.Now()
	elapsed := now.Sub(timeControl.start)

	return elapsed
}

func (timeControl *TimeControl) ShouldExecute() bool {
	elapsed := timeControl.GetElapsed()

	if elapsed.Milliseconds() >= TickRateMS {
		timeControl.start = time.Now()
		return true
	}
	return false
}

func CreateTimeControl() *TimeControl {
	return &TimeControl{start: time.Now()}
}
