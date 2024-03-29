package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"heid9/downtime/api/cli"
	"heid9/downtime/api/parsing"
	"heid9/downtime/api/storage"
	implcli "heid9/downtime/cli"
	implparsing "heid9/downtime/parsing"
	implstorage "heid9/downtime/storage/inmem"
)

type App struct {
	cmd    cli.Command
	ctx    storage.Context
	parser parsing.Parser
	dur    time.Duration
	urls   []string
}

func NewApp() *App {
	app := &App{
		cmd:    implcli.NewCommand(),
		ctx:    implstorage.NewContext(),
		parser: implparsing.NewParser(),
	}
	dur, ok := app.cmd.DurationArg()
	if !ok {
		log.Fatalln("Укажите параметр -i <секунды>")
	}
	app.dur = dur
	app.urls = app.cmd.Urls()
	if len(app.urls) == 0 {
		log.Fatalln("Укажите список разделенный \\n в STDIN. Например: cat links.txt | go run . -i 5")
	}
	return app
}

func (app *App) Start() {
	for {
		started := time.Now()
		app.Process(app.urls)
		now := time.Now()
		// при ожидании учитываем уже прошедшее время
		d := app.dur - now.Sub(started)
		if d > 0 {
			time.Sleep(d)
		}
	}
}

func (app *App) Process(urls []string) {
	g := &sync.WaitGroup{}
	ch := make(chan storage.Result)
	for _, domain := range urls {
		g.Add(1)
		go app.HandleRequest(domain, g, ch)
	}
	go app.StoreResults(app.ctx, ch)
	g.Wait()
	close(ch)
	fmt.Println("Finished")
	for _, res := range app.ctx.Results() {
		fmt.Println(res)
	}
}

func (app *App) StoreResults(
	ctx storage.Context,
	ch <-chan storage.Result,
) {
	for result := range ch {
		ctx.Add(result)
	}
}

func (app *App) HandleRequest(
	domain string,
	g *sync.WaitGroup,
	ch chan<- storage.Result,
) {
	var state bool
	reader, err := app.parser.LoadFromUrl(domain)
	if err == nil {
		state = app.parser.Match(reader)
	}
	ch <- implstorage.NewResult(domain, state)
	g.Done()
}
