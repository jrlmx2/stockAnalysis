package logger

import (
	"os"

	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/op/go-logging"
)

// NewLogger wraps the logger creation code
func NewLogger(conf config.LogConfig) (*logging.Logger, error) {
	var log = logging.MustGetLogger(conf.Name)

	if conf.Format == "" {
		conf.Format = `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`
	}
	var logFormatter = logging.MustStringFormatter(conf.Format)

	//regular log file
	var stdlog *os.File
	if _, err := os.Stat(conf.File + ".log"); err != nil {
		stdlog, err = os.Create(conf.File + ".log")
		if err != nil {
			return nil, err
		}
	} else {
		stdlog, err = os.Open(conf.File + ".log")
		if err != nil {
			return nil, err
		}
	}

	var stderr *os.File
	if _, err := os.Stat(conf.File + ".err"); err != nil {
		stderr, err = os.Create(conf.File + ".err")
		if err != nil {
			return nil, err
		}
	} else {
		stderr, err = os.Open(conf.File + ".err")
		if err != nil {
			return nil, err
		}
	}

	backend1 := logging.NewLogBackend(stderr, "", 0)
	backend2 := logging.NewLogBackend(stdlog, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, logFormatter)
	backend1Leveled := logging.AddModuleLevel(backend1)

	switch conf.Level {
	case "DEBUG":
		backend1Leveled.SetLevel(logging.DEBUG, "")
	case "INFO":
		backend1Leveled.SetLevel(logging.INFO, "")
	case "NOTICE":
		backend1Leveled.SetLevel(logging.NOTICE, "")
	case "WARN":
		backend1Leveled.SetLevel(logging.WARNING, "")
	case "ERROR":
		backend1Leveled.SetLevel(logging.ERROR, "")
	case "CRIT":
		backend1Leveled.SetLevel(logging.CRITICAL, "")
	}

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)

	return log, nil
}

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
//type Password string

//func (p Password) Redacted() interface{} {
//    return logging.Redact(string(p))
//}
