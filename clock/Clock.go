package clock

import (
	"fmt"
	"github.com/alex023/clock"
	"sync"
	"time"
)

var JOBS *Clock

func init() {
	JOBS = &Clock{
		clock: clock.Default(),
	}
}

type Clock struct {
	clock *clock.Clock
	jobs  sync.Map
}

func (c *Clock) Run(name string, actionInterval time.Duration, jobFunc func()) {
	v, ok := c.jobs.Load(name)
	if ok {
		v.(clock.Job).Cancel()
	}
	job, ok := c.clock.AddJobRepeat(actionInterval, 0, func() {
		if err := recover(); err != nil {
			fmt.Println("[ERROR] clock job:", err)
			jobFunc()
		}
		jobFunc()
	})
	if ok {
		c.jobs.Store(name, job)
	}
}

func (c *Clock) Cancel(name string) {
	v, ok := c.jobs.Load(name)
	if ok {
		v.(clock.Job).Cancel()
	}
}
