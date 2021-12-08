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
	DEFAULT = "\033[0m"
	BLUE    = "\033[34m"
	YELLOW  = "\033[33m"
	RED     = "\033[31m"

	TRACE = "Trace"
	INFO  = "Info "
	WARN  = "Warn "
	ERROR = "Error"
)

//
// Actions
//

func Trace(message string, opts ...interface{}) {
	formatted := formatMessage(TRACE, message, nil, opts...)
	log.Print(BLUE, formatted, DEFAULT)
}

func Info(message string, opts ...interface{}) {
	formatted := formatMessage(INFO, message, nil, opts...)
	log.Print(formatted)
}

func Warn(message string, err error, opts ...interface{}) {
	formatted := formatMessage(WARN, message, err, opts...)
	log.Print(YELLOW, formatted, DEFAULT)
}

func Error(message string, err error, opts ...interface{}) {
	if err != nil {
		formatted := formatMessage(ERROR, message, err, opts...)
		log.Fatal(RED, formatted, DEFAULT)
	}
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
