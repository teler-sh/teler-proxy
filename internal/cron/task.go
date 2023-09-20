package cron

import "github.com/kitabisa/teler-waf/threat"

var task = func() error {
	updated, err := threat.IsUpdated()
	if err != nil {
		return err
	}

	if !updated {
		err = threat.Get()
		if err != nil {
			return err
		}
	}

	return nil
}
