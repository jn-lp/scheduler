package scheduler

type FIFO struct {
	*scheduler
	current *Task
}

func NewFIFO(tasks []*Task) Scheduler {
	return &FIFO{scheduler: newScheduler(tasks)}
}

func (f *FIFO) Start(cores int) *Report {
	return f.scheduler.Start(func() *Task {
		if _, ok := f.Queue[f.current]; !ok {
			var cur *Task
			for task := range f.Queue {
				cur = task

				break
			}
			f.current = cur
		}

		return f.current
	}, cores)
}
