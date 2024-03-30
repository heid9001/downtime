package api

import (
	"errors"
)

var MatchError = errors.New("failed to match line")

type Console interface {
	ParseTask(line string) (Task, error)
	Start() <-chan Task
}
