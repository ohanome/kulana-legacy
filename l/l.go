package l

import (
	"github.com/op/go-logging"
	"kulana/options"
	"os"
)

var log = logging.MustGetLogger("kulana")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:04x}%{color:reset} %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.ERROR, "")
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendLeveled, backendFormatter)
}

func Debug(level int, args string) {
	if options.VerbosityLevel() >= level {
		log.Debug(args)
	}
}

func Info(args string) {
	log.Info(args)
}

func Notice(args string) {
	log.Notice(args)
}

func Warning(args string) {
	log.Warning(args)
}

func Error(args string) {
	log.Error(args)
}

func Critical(args string) {
	log.Critical(args)
}
