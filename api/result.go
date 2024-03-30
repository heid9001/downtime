package api

import (
	"time"
)

// Результат проверки страницы
type Result interface {
	SetDowntime(time.Duration)
	SetContext(Context)
	State() bool
	Created() time.Time
	DownTime() time.Duration
	String() string
	Context() Context
	Domain() string
}
