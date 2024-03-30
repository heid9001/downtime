package main

import (
	"heid9/downtime/api"
	"heid9/downtime/cli"
	"heid9/downtime/parsing"
	"heid9/downtime/scheduler"
	"heid9/downtime/storage/inmem"
)

type App struct {
	parser    api.Parser
	context   api.Context
	console   api.Console
	scheduler api.Scheduler
}

func NewApp() api.App {
	app := &App{
		parser:  parsing.NewParser(),
		console: cli.NewConsole(),
		context: inmem.NewContext(),
	}
	app.scheduler = scheduler.NewScheduler(app)
	return app
}

func (a App) Console() api.Console {
	return a.console
}

func (a App) Context() api.Context {
	return a.context
}

func (a App) Parser() api.Parser {
	return a.parser
}

func (a App) Scheduler() api.Scheduler {
	return a.scheduler
}
