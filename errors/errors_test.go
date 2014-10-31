package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestStackTrace(t *testing.T) {
	const testMsg = "test error"
	er := New(0, testMsg)
	e := er.(*GlobalWaysBaseError)

	if e.Msg != testMsg {
		t.Error("error message %s != expected %s", e.Msg, testMsg)
	}

	if strings.Index(e.Stack, "dropbox/util/errors/errors.go") != -1 {
		t.Error("stack trace generation code should not be in the error stack trace")
	}

	if strings.Index(e.Stack, "TestStackTrace") == -1 {
		t.Error("stack trace must have test code in it")
	}

	// compile-time tools.goEnv.test to ensure that DropboxError conforms to error interface
	var err error = e
	_ = err
}

func TestWrappedError(t *testing.T) {
	const (
		innerMsg  = "I am inner error"
		middleMsg = "I am the middle error"
		outerMsg  = "I am the mighty outer error"
	)
	inner := fmt.Errorf(innerMsg)
	middle := Wrap(0, inner, middleMsg)
	outer := Wrap(0, middle, outerMsg)
	errorStr := outer.Error()

	if strings.Index(errorStr, innerMsg) == -1 {
		t.Errorf("couldn't find inner error message in:\n%s", errorStr)
	}

	if strings.Index(errorStr, middleMsg) == -1 {
		t.Errorf("couldn't find middle error message in:\n%s", errorStr)
	}

	if strings.Index(errorStr, outerMsg) == -1 {
		t.Errorf("couldn't find outer error message in:\n%s", errorStr)
	}
}
