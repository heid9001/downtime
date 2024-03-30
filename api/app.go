package api

type App interface {
	Console() Console
	Context() Context
	Parser() Parser
	Scheduler() Scheduler
}
