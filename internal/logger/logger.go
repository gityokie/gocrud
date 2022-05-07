package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	env := os.Getenv("PULSE_ID_ENV")

	Log.SetReportCaller(true)

	if env == "dev" {
		Log.SetLevel(logrus.DebugLevel)
	}

	// TODO: Add hooks to save the logs somewhere
}
