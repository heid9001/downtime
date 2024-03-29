package cli

import (
	"time"
)

type Command interface {
	DurationArg() (time.Duration, bool)
	Urls() []string
}
