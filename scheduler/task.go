package scheduler

import (
	"fmt"
	"heid9/downtime/api"
	"time"
)

type Task struct {
	domain string
	dur    time.Duration
}

func NewTask(
	domain string,
	dur time.Duration,
) api.Task {
	return &Task{
		domain: domain,
		dur:    dur,
	}
}

func (t *Task) Domain() string {
	return t.domain
}

func (t *Task) Dur() time.Duration {
	return t.dur
}

func (t *Task) String() string {
	return fmt.Sprintf("Task{%s %s %s}\n", t.domain, t.dur)
}
