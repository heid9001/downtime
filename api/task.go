package api

import "time"

type Task interface {
	Domain() string
	Dur() time.Duration
	String() string
}
