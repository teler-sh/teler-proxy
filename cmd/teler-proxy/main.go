package main

import "github.com/kitabisa/teler-proxy/internal/runner"

func main() {
	opt := runner.ParseOptions()

	if err := opt.Validate(); err != nil {
		opt.Logger.Fatal("Cannot validate options", "err", err)
	}

	if err := runner.New(opt); err != nil {
		opt.Logger.Fatal("Something went wrong", "err", err)
	}
}
