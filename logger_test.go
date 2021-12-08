package logger

import (
	"errors"
	"math/rand"
	"strings"
	"testing"
)

//
// Tests
//

func TestTrace(t *testing.T) {
	Trace("%s", "blue")
}

func TestInfo(t *testing.T) {
	Info("%s", "default")
}

func TestWarnWithoutErr(t *testing.T) {
	Warn(nil, "%s", "yellow")
}

func TestWarnWithErr(t *testing.T) {
	err := errors.New("error message")
	Warn(err, "%s", "yellow")
}

func TestWarnWithLongErr(t *testing.T) {
	err := errors.New(randomSentence(20, 10))
	Warn(err, "%s", "yellow")
}

//
// Error
// In order to test, changed the error method to panic
// and uncomment out the code below
//

// func mockRecover() {
// 	recover()
// }

// func TestErrorWithoutErr(t *testing.T) {
// 	defer mockRecover()

// 	Error("%s", nil, "red", testFlag{})
// }

// func TestErrorWithErr(t *testing.T) {
// 	defer mockRecover()

// 	err := errors.New("error message")
// 	Error("%s", err, "red", testFlag{})
// }

// func TestErrWithLongErr(t *testing.T) {
// 	defer mockRecover()

// 	err := errors.New(randomSentence(20, 10))
// 	Error("%s", err, "red", testFlag{})
// }

//
// Tools
//

func randomLetter() byte {
	letters := "abcdefghijklmnopqrstuvwxyz"
	return letters[rand.Intn(len(letters))]
}

func randomWord(maxWordLength int) string {
	word := strings.Builder{}

	length := 1 + rand.Intn(maxWordLength)
	for i := 0; i < length; i++ {
		word.WriteByte(randomLetter())
	}

	return word.String()
}

func randomSentence(wordCount, maxWordLength int) string {
	sentence := strings.Builder{}

	for i := 0; i < wordCount; i++ {
		word := randomWord(maxWordLength)
		sentence.WriteString(word)
		sentence.WriteString(" ")
	}

	return sentence.String()
}
