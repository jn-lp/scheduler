package scheduler

type Task struct {
	Start, Wcet, Priority, Deadline, Progress int
}

func (t *Task) Done() bool {
	return t.Wcet <= t.Progress
}
