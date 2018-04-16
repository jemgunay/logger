// Package logger is a log.Logger wrapper.
package logger

import (
	"fmt"
	"io"
	"log"
	"strings"
)

var (
	prefixPaddingEnabled = true
	loggers              = make(map[*Logger]bool)
	prefixMaxPadding     = 0
)

// MsgFormatter is used to pass a function to a logger for formatting messages before they are output.
type MsgFormatter func(string) string

// Logger is log.Logger wrapped with extra features.
type Logger struct {
	logger       *log.Logger
	enabled      bool
	msgPrefix    string
	msgFormatter MsgFormatter
}

// NewLogger creates a new logger given an io.Writer to log to, a prefix to display before the datetime and
// a flag to determine whether the logger is enabled by default.
func NewLogger(handle io.Writer, loggerPrefix string, enabled bool) Logger {
	if loggerPrefix != "" {
		loggerPrefix = "[" + loggerPrefix + "] "
	}

	// create new logger
	l := log.New(handle, loggerPrefix, log.Ldate|log.Ltime)
	newLogger := Logger{logger: l, enabled: enabled}

	// reset prefix padding
	loggers[&newLogger] = true
	if len(loggerPrefix) > prefixMaxPadding {
		prefixMaxPadding = len(loggerPrefix)
	}
	SetPrefixPadding(prefixPaddingEnabled)

	return newLogger
}

// SetPrefixPadding is used to enable or disable padding after the prefix to align logged messages.
func SetPrefixPadding(enabled bool) {
	prefixPaddingEnabled = enabled

	for l := range loggers {
		resetPrefix := strings.TrimSpace(l.logger.Prefix())

		if prefixPaddingEnabled {
			paddingCount := prefixMaxPadding - len(resetPrefix)
			resetPrefix += strings.Repeat(" ", paddingCount)

		} else {
			// remove all padding from each logger prefix
			if resetPrefix != "" {
				resetPrefix += " "
			}
		}

		l.logger.SetPrefix(resetPrefix)
	}
}

// performLog formats & writes a log msg to the io.Writer. Each of the Logx functions depend on performLog.
func (l *Logger) performLog(msg string, newline bool) {
	if l.enabled == false {
		return
	}

	// perform formatting on string
	if l.msgFormatter != nil {
		msg = l.msgFormatter(msg)
	}
	msg = l.msgPrefix + msg

	if newline {
		msg += "\n\n"
	}

	l.logger.Print(msg)
}

// Log logs the provided message if the Logger is enabled.
func (l *Logger) Log(msg ...interface{}) {
	l.performLog(fmt.Sprint(msg...), false)
}

// Logf logs the provided message with formatting if the Logger is enabled.
func (l *Logger) Logf(format string, args ...interface{}) {
	l.performLog(fmt.Sprintf(format, args...), false)
}

// Logln logs the provided message followed by a new line if the Logger is enabled.
func (l *Logger) Logln(msg ...interface{}) {
	l.performLog(fmt.Sprint(msg...), true)
}

// Enable enables the logger.
func (l *Logger) Enable() {
	l.enabled = true
}

// Disable disables the logger.
func (l *Logger) Disable() {
	l.enabled = false
}

// MessagePrefix returns the currently set message prefix.
func (l *Logger) MessagePrefix() string {
	return l.msgPrefix
}

// SetMessagePrefix sets the message prefix.
func (l *Logger) SetMessagePrefix(messagePrefix string) {
	l.msgPrefix = messagePrefix
}

// SetMsgFormatter sets the function used to format a string. Set to null to disable formatting.
func (l *Logger) SetMsgFormatter(formatter MsgFormatter) {
	l.msgFormatter = formatter
}

// SetOutput sets the output destination, i.e. Stdout, Stderr, a text file.
func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

// SetFlag sets the output flags.
func (l *Logger) SetFlag(flag int) {
	l.logger.SetFlags(flag)
}
