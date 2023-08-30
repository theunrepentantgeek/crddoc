package cmd

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func CreateLogger() logr.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05.999",
	}

	zl := zerolog.New(output).
		With().Timestamp().
		Logger()

	zerologr.SetMaxV(0)

	// Use standard interface for logging
	zerologr.VerbosityFieldName = "" // Don't include verbosity in output
	log := zerologr.New(&zl)

	return log
}

func verboseLogging() {
	zerologr.SetMaxV(1)
}
