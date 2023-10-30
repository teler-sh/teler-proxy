package cron

import (
	"os"

	"github.com/kitabisa/teler-waf/threat"
)

var task = func() error {
	updated, err := threat.IsUpdated()
	if err != nil {
		return err
	}

	if !updated {
		path, err := threat.Location()
		if err != nil {
			return err
		}

		if err = os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}
