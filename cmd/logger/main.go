package main

import (
	"errors"
	"github.com/jemgunay/logger"
	"os"
	"strings"
)

var (
	Info      = logger.NewLogger(os.Stdout, "INFO", false)
	Error     = logger.NewLogger(os.Stderr, "ERROR", true)
	Incoming  = logger.NewLogger(os.Stdout, "INCOMING", true)
	Outgoing  = logger.NewLogger(os.Stdout, "OUTGOING", true)
	File      = logger.NewLogger(nil, "FILE", true)
	Formatted = logger.NewLogger(os.Stdout, "FORMATTED", true)
)

func main() {
	// enable a previously disabled logger
	Info.Enable()
	Info.Log("this logger has been enabled")

	// an error logger (outputs to Stderr)
	err := errors.New("this is an error")
	Error.Logf("example formatted error message, err:[%v]", err.Error())

	// an incoming request logger
	Incoming.SetMessagePrefix("< ")
	Incoming.Logf("this is a formatted incoming request: %v", "localhost:8080/test")

	// an outgoing request logger
	Outgoing.Enable()
	Outgoing.SetMessagePrefix("> ")
	Outgoing.Logln("this is an outgoing request followed by a new line: ", "http://google.co.uk")

	Outgoing.Disable()
	Outgoing.Log("this outgoing request will not be written to Stdout: ", "http://google.co.uk")

	// a logger which writes to a file
	fileWriter, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	// close the file when finished logging
	defer fileWriter.Close()

	File.SetOutput(fileWriter)
	File.Log("this message has been written to a file")

	// provide a function to format all logged messages, e.g.
	// ...to capitalise whole message:
	Formatted.SetMsgFormatter(strings.ToUpper)
	Formatted.Logf("this message used to be lower case")

	// ...to alternate letter case:
	alternateCase := func(msg string) (result string) {
		i := 0
		for _, char := range msg {
			i++
			if i%2 == 0 {
				result += strings.ToUpper(string(char))
				continue
			}
			result += string(char)
		}
		return
	}
	Formatted.SetMsgFormatter(alternateCase)
	Formatted.Logf("this message also used to be lower case")
}
