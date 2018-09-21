package core

import (
	"fmt"
	"os"

	logging "github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("fusaker")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} %{level:.4s} %{id:03x}%{color:reset} # %{message}`,
)

func SetupLogger() {
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	logfile, err := os.OpenFile(Config.LogFile, os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err {
		fmt.Println(err.Error())
	} else {
		backend1 = logging.NewLogBackend(logfile, "", 0)
	}

	backendFormatter := logging.NewBackendFormatter(backend2, format)
	backendLeveled := logging.AddModuleLevel(backend1)
	loglevel, err := logging.LogLevel(Config.LogLevel)
	if nil != err {
		backendLeveled.SetLevel(logging.INFO, "")
	} else {
		backendLeveled.SetLevel(loglevel, "")
	}
	logging.SetBackend(backendLeveled, backendFormatter)
}
