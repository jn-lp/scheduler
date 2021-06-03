package scheduler

type Report struct {
	QueueSizes, WaitTimes     []int
	QueueAvgSize, AvgWaitTime float32
}
