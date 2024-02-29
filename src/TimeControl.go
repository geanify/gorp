package main

import "time"

type TimeControl struct {
	start time.Time
}

func (timeControl *TimeControl) getElapsed() time.Duration {
	now := time.Now()
	elapsed := now.Sub(timeControl.start)

	return elapsed
}

func (timeControl *TimeControl) shouldExecute() bool {
	elapsed := timeControl.getElapsed()

	if elapsed.Milliseconds() >= tickRateMS {
		timeControl.start = time.Now()
		return true
	}
	return false
}

func createTimeControl() *TimeControl {
	return &TimeControl{start: time.Now()}
}
