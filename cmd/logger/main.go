package main

import (
	"errors"
	"github.com/jemgunay/logger"
	"os"
	"strings"
	"time"
)

var (
	Plain     = &logger.Logger{Writer: os.Stdout, Enabled: true}
	Info      = logger.NewLogger(os.Stdout, "INFO", false)
	Error     = logger.NewLogger(os.Stderr, "ERROR", true)
	Incoming  = logger.NewLogger(os.Stdout, "INCOMING", true)
	Outgoing  = logger.NewLogger(os.Stdout, "OUTGOING", true)
	File      = logger.NewLogger(nil, "FILE", true)
	Formatted = logger.NewLogger(os.Stdout, "FORMATTED", true)
)

func main() {
	example()
	time.Sleep(time.Millisecond)
}

func example() {
	/*
	 * Enable a previously disabled logger.
	 */
	Info.Enable()
	Info.Log("this logger has been enabled")

	/*
	 * An error logger (outputs to Stderr).
	 */
	for i := 1; i <= 4; i++ {
		err := errors.New("this is an error")
		Error.Logf("example error message no. %v, err:[%v]", i, err.Error())
	}

	/*
	 * An incoming request logger.
	 */
	Incoming.Message.Formatter = func(msg string) string { return "< " + msg }
	Incoming.Logf("this is a formatted incoming request: %v", "localhost:8080/retrieve")

	/*
	 * An outgoing request logger.
	 */
	Outgoing.Enable()
	Outgoing.Message.Formatter = func(msg string) string { return "> " + msg }
	Outgoing.Logln("this is an outgoing request followed by a new line: ", "http://google.com")

	/*
	 * A disabled logger.
	 */
	Outgoing.Disable()
	Outgoing.Log("this outgoing request will not be written to Stdout: ", "http://google.com")

	/*
	 * A logger which writes to a file.
	 */
	fileWriter, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	File.Writer = fileWriter
	File.Log("this message has been written to a file")

	/*
	 * Provide a function to format all logged messages, e.g.
	 */
	// ...to capitalise whole message:
	Formatted.Message.Formatter = strings.ToUpper
	Formatted.Logf("this message used to be lower case")

	// ...to alternate letter case:
	alternateCase := func(s string) (result string) {
		i := 0
		for _, char := range s {
			i++
			if i%2 == 0 {
				result += strings.ToUpper(string(char))
				continue
			}
			result += string(char)
		}
		return
	}
	Formatted.Message.Formatter = alternateCase
	Formatted.Logf("this message also used to be lower case")
}
