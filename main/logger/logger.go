package logger

import (
	"github.com/go-kit/kit/log"
	"os"
)

var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.NewSyncLogger(Logger)
	Logger = log.With(Logger,
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
}
