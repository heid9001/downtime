// по большей части это DTO результата проверки домена
package inmem

import (
	"fmt"
	"time"

	api "heid9/downtime/api/storage"
)

const (
	DOWNTIME_MSG = "%s - recovered after %s\n"
	OK_MSG       = "%s - ok\n"
	FAIL_MSG     = "%s - fail\n"
)

type Result struct {
	domain   string
	state    bool
	downTime time.Duration
	time     time.Time
	ctx      api.Context
}

func NewResult(domain string, state bool) api.Result {
	return NewResultAt(domain, state, time.Now())
}

func NewResultAt(domain string, state bool, dt time.Time) api.Result {
	return &Result{
		domain: domain,
		state:  state,
		time:   dt,
	}
}

func (res *Result) SetContext(ctx api.Context) {
	res.ctx = ctx
}

func (res *Result) Domain() string {
	return res.domain
}

func (res *Result) State() bool {
	return res.state
}

func (res *Result) Context() api.Context {
	return res.ctx
}

func (res *Result) String() string {
	if res.state && res.downTime > 0 {
		return fmt.Sprintf(DOWNTIME_MSG, res.domain, res.downTime)
	} else if res.state {
		return fmt.Sprintf(OK_MSG, res.domain)
	} else {
		return fmt.Sprintf(FAIL_MSG, res.domain)
	}
}

func (res *Result) DownTime() time.Duration {
	return res.downTime
}

func (res *Result) SetDowntime(downTime time.Duration) {
	res.downTime = downTime
}

func (res *Result) Created() time.Time {
	return res.time
}
