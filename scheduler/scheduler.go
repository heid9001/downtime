package scheduler

import (
	"fmt"
	"heid9/downtime/api"
	storage "heid9/downtime/storage/inmem"
	"os"
	"os/signal"
	"time"
)

type Scheduler struct {
	app   api.App
	tasks []api.Task
}

func NewScheduler(app api.App) api.Scheduler {
	return &Scheduler{
		app: app,
	}
}

func (s *Scheduler) Run(task api.Task) {
	go func() {
		for range time.Tick(task.Dur()) {
			var state bool
			r, err := s.app.Parser().LoadFromUrl(task.Domain())
			if err == nil {
				state = s.app.Parser().Match(r)
			}
			fmt.Println(s.app.Context().Add(storage.NewResult(task.Domain(), state)))
		}
	}()
}

func (s *Scheduler) Start() {
	var (
		sig = make(chan os.Signal)
	)
	signal.Notify(sig, os.Interrupt)
	tasks := s.app.Console().Start()
	for {
		select {
		case task := <-tasks:
			s.Run(task)
		case s := <-sig:
			if s == os.Interrupt {
				fmt.Print("bye")
				os.Exit(0)
			}
		}
	}
}

func (s *Scheduler) Add(task api.Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) App() api.App {
	return s.app
}
