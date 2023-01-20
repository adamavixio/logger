package logger

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

//
// Definitions
//

const (
	reset  = "\033[0m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	red    = "\033[31m"

	tracePrefix = "Trace"
	infoPrefix  = "Info "
	warnPrefix  = "Warn "
	errorPrefix = "Error"
)

//
// Exports
//

// Trace log messages are meant to track progress
// through a system or process. The opts are variables
// that can be subsituted into the message using format.
func Trace(message string, opts ...interface{}) {
	formatted := formatMessage(tracePrefix, message, nil, opts...)
	log.Print(blue, formatted, reset)
}

// Info log messages are meant to display information
// or data that is unrelated to the progress of a system.
// The opts are variables that can be subsituted into the
// message using format.
func Info(message string, opts ...interface{}) {
	formatted := formatMessage(infoPrefix, message, nil, opts...)
	log.Print(formatted)
}

// Warn log messages are meant to display information for an
// error that does not require the program to exit. If the
// error is nil, there will be no output. The opts are variables
// that can be subsituted into the message using format.
func Warn(err error, message string, opts ...interface{}) {
	formatted := formatMessage(warnPrefix, message, err, opts...)
	log.Print(yellow, formatted, reset)
}

// Error log messages are meant to display information for an
// error that does not require the program to exit. If the
// error is nil, there will be no output. The opts are variables
// that can be subsituted into the message using format.
func Error(err error, message string, opts ...interface{}) {
	formatted := formatMessage(errorPrefix, message, err, opts...)
	log.Print(red, formatted, reset)
}

// Fatal log messages are meant to display information for an
// error that does require the program to exit. If the
// error is nil, there will be no output and the program
// will not exit. The opts are variables that can be subsituted
// into the message using format.
func Fatal(err error, message string, opts ...interface{}) {
	formatted := formatMessage(errorPrefix, message, err, opts...)
	log.Fatal(red, formatted, reset)
}

//
// Tools
//

func base(level string, message string) string {
	log.SetFlags(0)
	return fmt.Sprintf("%s | %s | %s", level, timestamp(), message)
}

func timestamp() string {
	date := time.Now().UTC()
	return fmt.Sprintf(
		"%04d/%02d/%02d %02d:%02d:%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second(),
	)
}

func formatMessage(level string, message string, err error, opts ...interface{}) string {
	formatted := strings.Builder{}

	formattedMessage := fmt.Sprintf(base(level, message), opts...)
	formatted.WriteString(formattedMessage)

	if err != nil {
		errorMessage := formatError(err)
		formatted.WriteString(errorMessage.Error())
	}

	return formatted.String()
}

func formatError(err error) error {
	message := err.Error()

	buff := strings.Builder{}
	buff.WriteString(prefix())

	flag := false
	for i, byte := range message {
		buff.WriteRune(byte)

		if i > 0 && i%50 == 0 {
			flag = true
			continue
		}

		if flag && string(byte) == " " {
			buff.WriteString(prefix())
			flag = false
		}
	}

	return errors.New(buff.String())
}

func prefix() string {
	builder := strings.Builder{}

	builder.WriteString("\n")
	builder.WriteString(spaces(6))
	builder.WriteString("|")
	builder.WriteString(spaces(1))

	return builder.String()
}

func spaces(amount int) string {
	spaces := strings.Builder{}

	for i := 0; i < amount; i++ {
		spaces.WriteString(" ")
	}

	return spaces.String()
}
