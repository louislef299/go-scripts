package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Package level variables, which are pointer to log.Logger.
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type Logger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

// initLog initializes log.Logger objects
func initLog(traceHandle, infoHandle, warningHandle, errorHandle io.Writer, isFlag bool) *Logger {
	// Flags for defines the logging properties, to log.New
	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}

	l := &Logger{}
	// Create log.Logger objects.
	l.Trace = log.New(traceHandle, "TRACE: ", flag)
	l.Info = log.New(infoHandle, "INFO: ", flag)
	l.Warning = log.New(warningHandle, "WARNING: ", flag)
	l.Error = log.New(errorHandle, "ERROR: ", flag)

	return l
}

// SetLogLevel sets the logging level preference
func SetLogLevel(level, file string) *Logger {
	var out io.Writer
	// find if outputting to logfile and generate the logger
	if file != "" {
		// Creates os.*File, which has implemented io.Writer intreface
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %s", err.Error())
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	// Calls function initLog by specifying log level preference.
	switch level {
	case "TRACE":
		return initLog(out, out, out, out, true)
	case "INFO":
		return initLog(ioutil.Discard, out, out, out, true)
	case "WARNING":
		return initLog(ioutil.Discard, ioutil.Discard, out, out, true)
	case "ERROR":
		return initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, out, true)
	default:
		log.Printf("log level %s not found", level)
		return initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, false)
	}
}
