package scheduler

type Task struct {
	Start, Wcet, Priority, Deadline, Progress int
	Protected                                 bool
}

func (t *Task) Done() bool {
	return t.Wcet <= t.Progress
}
