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
)

//
// Actions
//

func Trace(message string, opts ...interface{}) {
	formatted := fmt.Sprintf(base("Trace", message), opts...)
	log.Print(blue, formatted, reset)
}

func Info(message string, opts ...interface{}) {
	formatted := fmt.Sprintf(base("Info ", message), opts...)
	log.Print(formatted)
}

func Warn(message string, err error, opts ...interface{}) {
	formatted := strings.Builder{}

	formattedMessage := fmt.Sprintf(base("Error", message), opts...)
	formatted.WriteString(formattedMessage)

	if err != nil {
		errorMessage := formatError(err)
		formatted.WriteString(errorMessage.Error())
	}

	log.Print(yellow, formatted.String(), reset)
}

func Error(message string, err error, opts ...interface{}) {
	formatted := strings.Builder{}

	formattedMessage := fmt.Sprintf(base("Error", message), opts...)
	formatted.WriteString(formattedMessage)

	if err != nil {
		errorMessage := formatError(err)
		formatted.WriteString(errorMessage.Error())
	}

	log.Panic(red, formatted.String(), reset)
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
		"%d/%d/%d %d:%d:%d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second(),
	)
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
