package scheduler

type Scheduler interface {
	Start(cores int) *Report
}

type Constructor func([]*Task) Scheduler

type scheduler struct {
	Time  int
	Tasks []*Task
	Queue map[*Task]*struct{}
}

func newScheduler(tasks []*Task) *scheduler {
	return &scheduler{
		Tasks: tasks,
		Queue: make(map[*Task]*struct{}, len(tasks)),
	}
}

func (s *scheduler) Start(nextTask func() *Task, cores int) *Report {
	var (
		qs, wt []int
		task   *Task
	)

	for dones := 0; dones < len(s.Tasks); s.Time++ {
		for _, t := range s.Tasks {
			if t.Start == s.Time {
				s.Queue[t] = nil
			}
		}

		if l := len(s.Queue); l <= 0 {
			continue
		} else {
			qs = append(qs, l)
		}

		if task == nil || !task.Protected {
			task = nextTask()
		}
		task.Progress += cores

		if task.Done() || task.Deadline == s.Time {
			wt = append(wt, s.Time-task.Start)

			delete(s.Queue, task)

			dones++
		}
	}

	return &Report{
		QueueSizes:   qs,
		WaitTimes:    wt,
		QueueAvgSize: sum(qs) / float32(len(qs)),
		AvgWaitTime:  sum(wt) / float32(len(wt)),
	}
}

func sum(xs []int) (res float32) {
	for _, x := range xs {
		res += float32(x)
	}

	return
}
