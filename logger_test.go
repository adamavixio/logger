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
	Warn("%s", nil, "yellow")
}

func TestWarnWithErr(t *testing.T) {
	err := errors.New("error message")
	Warn("%s", err, "yellow")
}

func TestWarnWithLongErr(t *testing.T) {
	err := errors.New(randomSentence(20, 10))
	Warn("%s", err, "yellow")
}

func TestErrorWithoutErr(t *testing.T) {
	defer mockRecover()

	Error("%s", nil, "red")
}

func TestErrorWithErr(t *testing.T) {
	defer mockRecover()

	err := errors.New("error message")
	Error("%s", err, "red")
}

func TestErrWithLongErr(t *testing.T) {
	defer mockRecover()

	err := errors.New(randomSentence(20, 10))
	Error("%s", err, "red")
}

//
// Tools
//

func mockRecover() {
	recover()
}

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
