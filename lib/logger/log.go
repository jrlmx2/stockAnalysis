package log

import(
  "os"
  "github.com/op/go-logging"
)

//Wrapper function for creating different package logs
func NewLogger(name, format, location, level string) (*logging.Logger, error) {
  var log = logging.MustGetLogger(name)

  if format == "" {
    format = `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`
  }
  var logFormatter = logging.MustStringFormatter(format)

  //regular log file
  var stdlog *os.File
  if _, err := os.Stat(location+".log"); err != nil {
    stdlog, err = os.Create(location+".log")
    if err != nil {
      return nil, err
    }
  } else {
    stdlog, err = os.Open(location+".log")
    if err != nil {
        return nil, err
    }
  }

  var stderr *os.File
  if _, err := os.Stat(location+".err"); err != nil {
    stderr, err = os.Create(location+".err")
    if err != nil {
      return nil, err
    }
  } else {
    stderr, err = os.Open(location+".err")
    if err != nil {
        return nil, err
    }
  }



  backend1 := logging.NewLogBackend(stderr, "", 0)
  backend2 := logging.NewLogBackend(stdlog, "", 0)

  backend2Formatter := logging.NewBackendFormatter(backend2, logFormatter)
  backend1Leveled := logging.AddModuleLevel(backend1)

  switch (level){
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
