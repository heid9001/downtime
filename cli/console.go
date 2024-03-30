package cli

import (
	"bufio"
	"fmt"
	"heid9/downtime/api"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"heid9/downtime/scheduler"
)

type Console struct {
}

func NewConsole() api.Console {
	return &Console{}
}

func (c *Console) Start() <-chan api.Task {
	ch := make(chan api.Task)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			task, err := c.ParseTask(line)
			if err != nil {
				fmt.Println(err)
			}
			ch <- task
		}
	}()
	return ch
}

func (c *Console) ParseTask(line string) (api.Task, error) {
	var (
		domain string
		dur    time.Duration
	)
	pattern := regexp.MustCompile(`^(.+)\s-i\s+([0-9]+)$`)
	res := pattern.FindAllStringSubmatch(line, -1)
	if res == nil {
		return nil, api.MatchError
	}
	if _, err := url.ParseRequestURI(res[0][1]); err != nil {
		return nil, api.MatchError
	} else {
		domain = res[0][1]
	}
	durRaw, err := strconv.Atoi(res[0][2])
	if err != nil {
		return nil, api.MatchError
	} else {
		dur = time.Duration(durRaw) * time.Second
	}
	return scheduler.NewTask(domain, dur), nil
}
