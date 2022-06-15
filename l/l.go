package l

import (
	"github.com/lajosbencz/glo"
	"kulana/options"
	"kulana/setup"
	"os"
)

var log = glo.NewFacility()

func init() {
	f, err := os.OpenFile(setup.GetLogFile(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	handlerBfr := glo.NewHandler(f)
	log.PushHandler(handlerBfr)

	handlerStd := glo.NewHandler(os.Stdout)
	filter := glo.NewFilterLevel(0)
	handlerStd.PushFilter(filter)
	log.PushHandler(handlerStd)
}

func Debug(level int, args string) {
	if options.VerbosityLevel() >= level {
		err := log.Debug(args)
		if err != nil {
			Emergency(err.Error())
		}
	}
}

func Info(args string) {
	err := log.Info(args)
	if err != nil {
		Emergency(err.Error())
	}
}

func Notice(args string) {
	err := log.Notice(args)
	if err != nil {
		Emergency(err.Error())
	}
}

func Warning(args string) {
	err := log.Warning(args)
	if err != nil {
		Emergency(err.Error())
	}
}

func Error(args string) {
	err := log.Error(args)
	if err != nil {
		Emergency(err.Error())
	}
}

func Critical(args string) {
	err := log.Critical(args)
	if err != nil {
		Emergency(err.Error())
	}
}

func Emergency(args string) {
	err := log.Emergency(args)
	if err != nil {
		panic(err)
	}
	os.Exit(1)
}
