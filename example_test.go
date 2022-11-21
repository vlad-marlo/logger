package logger_test

import (
	"github.com/sirupsen/logrus"
	log "github.com/vlad-marlo/logger"
	"os"
)

func Example() {
	// initializing logger
	// same is
	/*
		logger := logrus.New()
		log.AddOpts(
			logger,
			log.WithLevel(logrus.DebugLevel),
			log.WithOutput(os.Stdout),
			// add opts whatever you want
		)
	*/
	logger := log.WithOpts(
		log.WithLevel(logrus.DebugLevel),
		log.WithOutput(os.Stdout),
		// add opts whatever you want
	)
	logger.Trace("this is trace msg; you'll not see it in logs")
	logger.Debug("this is debug msg; you'll see it in logs")

	// changing logger options
	log.AddOpts(
		logger,
		log.WithLevel(logrus.TraceLevel),
	)
	logger.Trace("this is trace msg; now you'll see it in logs")
	logger.Debug("this is debug msg; you'll see it in logs as before")
}
