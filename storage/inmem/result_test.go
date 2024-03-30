package inmem

import (
	"heid9/downtime/api"
	"testing"
	"time"
)

type unit struct {
	time.Time
	ctx      api.Context
	domain   string
	downtime time.Duration
	state    bool
}

// Проверка суммирования интервалов c fail состоянием (напр. -i 5sec => Fail, Ok, Fail, Fail , Ok => Ok with failure 10sec)
func TestAggregate(t *testing.T) {
	makeUnit("domain1", 3*time.Second, true).
		addState(false, time.Second).
		addState(true, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		assert("TestSimple#1", t)
	makeUnit("domain1", 3*time.Second, true).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		assert("TestSimple#2", t)
	makeUnit("domain1", 3*time.Second, true).
		addState(true, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		assert("TestSimple#3", t)
	makeUnit("domain1", 3*time.Second, true).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		assert("TestSimple#4", t)
	makeUnit("domain1", 0, true).
		addState(false, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		addState(true, time.Second).
		assert("TestSimple#5", t)
}

// Проверка с единичным интервалом
func TestSimple(t *testing.T) {
	makeUnit("domain1", time.Second, true).
		addState(false, time.Second).
		addState(true, time.Second).
		assert("TestSimple#1", t)
	makeUnit("domain1", 0, false).
		addState(true, time.Second).
		addState(false, time.Second).
		assert("TestSimple#2", t)
	makeUnit("domain1", time.Second, true).
		addState(true, time.Second).
		addState(false, time.Second).
		addState(true, time.Second).
		assert("TestSimple#3", t)
}

// Базовые тесты.
func TestTrue(t *testing.T) {
	makeUnit("domain1", 0, true).
		addState(true, time.Second).
		assert("TestTrue", t)
}

func TestFalse(t *testing.T) {
	makeUnit("domain1", 0, false).
		addState(false, time.Second).
		assert("TestFalse", t)
}

// Хелперы
func makeUnit(domain string, downtime time.Duration, state bool) *unit {
	return makeUnitWithContext(domain, downtime, state, NewContext())
}

func makeUnitWithContext(domain string, downtime time.Duration, state bool, ctx api.Context) *unit {
	return &unit{
		Time:     time.Now(),
		ctx:      ctx,
		domain:   domain,
		downtime: downtime,
		state:    state,
	}
}

func (u *unit) assert(name string, t *testing.T) {
	state, downtime := u.result()
	if state != u.state || downtime != u.downtime {
		t.Fatalf("%s => expected (state: %t, downtime: %s) != actual (state: %t, downtime: %s)",
			name,
			u.state, u.downtime,
			state, downtime,
		)
	}
}

func (u *unit) result() (bool, time.Duration) {
	res := u.ctx.Results()[u.domain]
	return res.State(), res.DownTime()
}

func (u *unit) addState(state bool, step time.Duration) *unit {
	u.ctx.Add(NewResultAt(u.domain, state, u.Time))
	(*u).Time = (*u).Add(step)
	return u
}
