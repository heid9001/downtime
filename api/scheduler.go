package api

type Scheduler interface {
	Run(task Task)
	Start()
}
