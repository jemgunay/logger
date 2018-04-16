## A log.Logger Wrapper

* Enable/disable specific loggers, providing more control over which log categories to output.
* Add a prefix to each message to clearly categorise message types.
* Choosing the output destination for each logger, i.e. Stdout, Stderr, a text file.
* Provide a logger with a custom function for formatting all output.
* Automatically pad messages in order to make output more readable.

#### main.go output:  

```
$ cd $GOPATH/src/github.com/jemgunay/logger/cmd/logger
$ go build && ./logger
[INFO]      2018/04/16 15:28:46 this logger has been enabled
[ERROR]     2018/04/16 15:28:46 example formatted error message, err:[this is an error]
[INCOMING]  2018/04/16 15:28:46 < this is a formatted incoming request: localhost:8080/test
[OUTGOING]  2018/04/16 15:28:46 > this is an outgoing request followed by a new line: http://google.co.uk

[FORMATTED] 2018/04/16 15:28:46 THIS MESSAGE USED TO BE LOWER CASE
[FORMATTED] 2018/04/16 15:28:46 tHiS MeSsAgE AlSo uSeD To bE LoWeR CaSe
```