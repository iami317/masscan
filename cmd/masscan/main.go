package main

import (
	"github.com/iami317/logx"
	"github.com/iami317/masscan/pkg/runner"
)

func main() {
	opts, err := runner.NewOptions()
	if err != nil {
		logx.Error(err.Error())
		return
	}

	runner, err := runner.NewRunner(opts)
	if err != nil {
		logx.Error(err.Error())
		return
	}

	ts, err := runner.GetTargets()
	if err != nil {
		logx.Error(err.Error())
		return
	}

	for t := range ts {
		logx.Debug(t)
	}
}
