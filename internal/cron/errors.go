package cron

import "errors"

var (
	errCallbackNotFunc = errors.New("callback parameter is not a function")
)
