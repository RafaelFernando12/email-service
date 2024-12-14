package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Level uint

const (
	FATAL Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

type logger struct {
	applicationName string
	level           Level
}

type printLogger struct {
	log   *log.Logger
	level Level
}

type MultiLogger interface {
	Debug() PrintLogger
	Info() PrintLogger
	Warning() PrintLogger
	Error() PrintLogger
	Fatal() PrintLogger
}

type PrintLogger interface {
	Printf(format string, a ...any)
	Println(a ...any)
	Print(a ...any)
}

func NewLogger(applicationName, level string) *logger {
	return &logger{
		applicationName: applicationName,
		level:           getLevelByString(strings.ToUpper(level)),
	}
}

func (l *Level) String() string {
	switch *l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}

	return ""
}

func getLevelByString(level string) Level {
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return DEBUG
	}
}

func (l *logger) getPrintLoggerByLevel(out io.Writer, level Level, flag int) *printLogger {
	if level > l.level {
		return &printLogger{level: level}
	}
	log := log.New(out, fmt.Sprintf("%s %s:", l.applicationName, level.String()), flag)

	return &printLogger{
		log:   log,
		level: level,
	}
}

func (l *logger) Debug() PrintLogger {
	return l.getPrintLoggerByLevel(os.Stdout, DEBUG, log.Ldate|log.Ltime)
}

func (l *logger) Info() PrintLogger {
	return l.getPrintLoggerByLevel(os.Stdout, INFO, log.Ldate|log.Ltime)
}

func (l *logger) Warning() PrintLogger {
	return l.getPrintLoggerByLevel(os.Stdout, WARNING, log.Ldate|log.Ltime)
}

func (l *logger) Error() PrintLogger {
	return l.getPrintLoggerByLevel(os.Stderr, ERROR, log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *logger) Fatal() PrintLogger {
	return l.getPrintLoggerByLevel(os.Stderr, FATAL, log.Ldate|log.Ltime|log.Lshortfile)
}

func (p *printLogger) Print(v ...any) {
	if p.log == nil {
		return
	}
	switch p.level {
	case FATAL:
		p.log.Fatal(v...)
	default:
		p.log.Print(v...)
	}
}

func (p *printLogger) Println(v ...any) {
	if p.log == nil {
		return
	}
	switch p.level {
	case FATAL:
		p.log.Fatalln(v...)
	default:
		p.log.Println(v...)
	}
}

func (p *printLogger) Printf(format string, a ...any) {
	if p.log == nil {
		return
	}
	switch p.level {
	case FATAL:
		p.log.Fatalf(format, a...)
	default:
		p.log.Printf(format, a...)
	}
}
