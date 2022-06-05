package logging

import (
	"log"
	"fmt"
)

type DefaultLogger struct {
	minLevel LogLevel
	loggers map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if (l.minLevel <= level) {
		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

func (l *DefaultLogger) Tracef(tpl string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(tpl, vals...))
}

func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

func (l *DefaultLogger) Debugf(tpl string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(tpl, vals...))
}

func (l *DefaultLogger) Info(msg string) {
	l.write(Info, msg)
}

func (l *DefaultLogger) Infof(tpl string, vals ...interface{}) {
	l.write(Info, fmt.Sprintf(tpl, vals...))
}

func (l *DefaultLogger) Warn(msg string) {
	l.write(Warn, msg)
}

func (l *DefaultLogger) Warnf(tpl string, vals ...interface{}) {
	l.write(Warn, fmt.Sprintf(tpl, vals...))
}

func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if (l.triggerPanic) {
		panic(msg)
	}
}

func (l *DefaultLogger) Panicf(tpl string, vals ...interface{}) {
	formattedMsg := fmt.Sprintf(tpl, vals...)
	l.write(Fatal, formattedMsg)
	if (l.triggerPanic) {
		panic(formattedMsg)
	}
}
