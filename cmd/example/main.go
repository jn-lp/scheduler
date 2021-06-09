package main

import (
	"log"
	"math/rand"

	"github.com/Arafatk/glot"
	"github.com/jn-lp/scheduler"
)

const (
	tasksCount = 10
	coresCount = 2
)

func main() {
	tasks := prepareTasks(tasksCount)
	constructors := map[string]scheduler.Constructor{
		"EDF":  scheduler.NewEDF,
		"FIFO": scheduler.NewFIFO,
		"RM":   scheduler.NewRM,
	}

	dimensions := 2

	QueueSizes, err := glot.NewPlot(dimensions, false, false)
	if err != nil {
		log.Fatal(err)
	}

	if err = QueueSizes.SetTitle("QueueSizes"); err != nil {
		log.Fatal(err)
	}

	WaitTimes, err := glot.NewPlot(dimensions, false, false)
	if err != nil {
		log.Fatal(err)
	}

	if err = WaitTimes.SetTitle("WaitTimes"); err != nil {
		log.Fatal(err)
	}

	for name, constructor := range constructors {
		tmp := make([]*scheduler.Task, len(tasks))

		for i, p := range tasks {
			v := p
			tmp[i] = &v
		}

		report := constructor(tmp).Start(coresCount)

		if err = QueueSizes.AddPointGroup(name, "lines", report.QueueSizes); err != nil {
			log.Fatal(err)
		}

		if err = WaitTimes.AddPointGroup(name, "lines", report.WaitTimes); err != nil {
			log.Fatal(err)
		}

		log.Printf("%s: Queue Avg Size = %f, Avg Wait Time = %f", name, report.QueueAvgSize, report.AvgWaitTime)
	}

	if err = QueueSizes.SavePlot("QueueSizes.png"); err != nil {
		log.Fatal(err)
	}

	if err = WaitTimes.SavePlot("WaitTimes.png"); err != nil {
		log.Fatal(err)
	}
}

func prepareTasks(n int) []scheduler.Task {
	tasks := make([]scheduler.Task, 0, n)

	for i := 1; i <= n; i++ {
		start := rand.Intn(i)
		maxPriority := 5
		maxDeadline := 10

		tasks = append(tasks, scheduler.Task{
			Start:     start,
			Wcet:      1 + rand.Intn(1),
			Priority:  rand.Intn(maxPriority),
			Deadline:  start + 1 + rand.Intn(maxDeadline),
			Protected: rand.Float32() > 0.5,
		})
	}

	return tasks
}
