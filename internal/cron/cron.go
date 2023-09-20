package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Cron struct {
	*gocron.Scheduler
	*gocron.Job
}

func New() (*Cron, error) {
	c := new(Cron)

	tz, err := time.LoadLocation("Local")
	if err != nil {
		return c, err
	}

	c.Scheduler = gocron.NewScheduler(tz)
	c.Job, err = c.Scheduler.Every(1).Day().At("00:00").WaitForSchedule().Do(task)

	return c, err
}
