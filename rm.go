package scheduler

type RM struct {
	*scheduler
	current *Task
}

func NewRM(tasks []*Task) Scheduler {
	return &RM{scheduler: newScheduler(tasks)}
}

func (rm *RM) Start(cores int) *Report {
	return rm.scheduler.Start(func() *Task {
		if _, ok := rm.Queue[rm.current]; !ok {
			var max int
			for task := range rm.Queue {
				if task.Priority > max {
					rm.current = task
				}
			}
		}

		return rm.current
	}, cores)
}
