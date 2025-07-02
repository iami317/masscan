package runner

import (
	"github.com/iami317/logx"
)

var Version = "0.0.1"

func ShowBanner() {
	logx.Debugf("\n\tM A S S C A N\t%s\n\n", Version)
}
