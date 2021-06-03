package scheduler

type EDF struct {
	*scheduler
	current *Task
}

func NewEDF(tasks []*Task) Scheduler {
	return &EDF{scheduler: newScheduler(tasks)}
}

func (e *EDF) Start() *Report {
	return e.scheduler.Start(func() *Task {
		if _, ok := e.Queue[e.current]; !ok {
			var max int
			for task := range e.Queue {
				if task.Deadline > max {
					e.current = task
				}
			}
		}

		return e.current
	})
}
