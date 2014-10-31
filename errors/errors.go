// This module implements functions which manipulate errors and provide stack
// trace information.
//
// NOTE: This package intentionally mirrors the standard "errors" module.
package errors

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

var GlobalWaysErrors map[int]string

// This interface exposes additional information about the error.
type GlobalWaysError interface {
	// This returns the error message without the stack trace.
	GetMessage() string

	// This returns the stack trace without the error message.
	GetStack() string

	// This returns the stack trace's context.
	GetContext() string

	// This returns the wrapped error.  This returns nil if this does not wrap
	// another error.
	GetInner() error

	// Implements the built-in error interface.
	Error() string

	// This returns the error code. This returns 0 if not error.
	GetCode() int

	// This returns true if is error, false or not.
	IsError() bool

	// This returns errcode & errmessage.
	ErrorMessage() string
}

// This returns the error string without stack trace information.
func GetMessage(err interface{}) string {
	switch e := err.(type) {
	case GlobalWaysError:
		dberr := GlobalWaysError(e)
		if !dberr.IsError() {
			return "ErrorCode == 0, OK"
		}
		ret := []string{}
		for dberr.IsError() {
			ret = append(ret, "ErrCode["+strconv.Itoa(int(dberr.GetCode()))+"]:")
			ret = append(ret, dberr.GetMessage())
			d := dberr.GetInner()
			if d == nil {
				break
			}
			var ok bool
			dberr, ok = d.(GlobalWaysError)
			if !ok {
				ret = append(ret, d.Error())
				break
			}
		}
		return strings.Join(ret, " ")
	case runtime.Error:
		return runtime.Error(e).Error()
	default:
		return "Passed a non-error to GetMessage"
	}
}

// A default implementation of the Error method of the error interface.
func DefaultError(e GlobalWaysError) string {
	// Find the "original" stack trace, which is probably the most helpful for
	// debugging.
	var origStack string
	errLines := make([]string, 1)
	errLines[0] = "ERROR:"
	fillErrorInfo(e, &errLines, &origStack)
	errLines = append(errLines, "")
	errLines = append(errLines, "ORIGINAL STACK TRACE:")
	errLines = append(errLines, origStack)
	return strings.Join(errLines, "\n")
}

// Fills errLines with all error messages, and origStack with the inner-most
// stack.
func fillErrorInfo(err error, errLines *[]string, origStack *string) {
	if err == nil {
		return
	}

	derr, ok := err.(GlobalWaysError)
	if ok {
		*errLines = append(*errLines, "ErrCode["+strconv.Itoa(int(derr.GetCode()))+"]:")
		*errLines = append(*errLines, derr.GetMessage())
		*origStack = derr.GetStack()
		fillErrorInfo(derr.GetInner(), errLines, origStack)
	} else {
		*errLines = append(*errLines, err.Error())
	}
}

// Returns a copy of the error with the stack trace field populated and any
// other shared initialization; skips 'skip' levels of the stack trace.
//
// NOTE: This panics on any error.
func stackTrace(skip int) (current, context string) {
	// grow buf until it's large enough to store entire stack trace
	buf := make([]byte, 128)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, len(buf)*2)
	}

	// Returns the index of the first occurrence of '\n' in the buffer 'b'
	// starting with index 'start'.
	//
	// In case no occurrence of '\n' is found, it returns len(b). This
	// simplifies the logic on the calling sites.
	indexNewline := func(b []byte, start int) int {
		if start >= len(b) {
			return len(b)
		}
		searchBuf := b[start:]
		index := bytes.IndexByte(searchBuf, '\n')
		if index == -1 {
			return len(b)
		} else {
			return (start + index)
		}
	}

	// Strip initial levels of stack trace, but keep header line that
	// identifies the current goroutine.
	var strippedBuf bytes.Buffer
	index := indexNewline(buf, 0)
	if index != -1 {
		strippedBuf.Write(buf[:index])
	}

	// Skip lines.
	for i := 0; i < skip; i++ {
		index = indexNewline(buf, index+1)
		index = indexNewline(buf, index+1)
	}

	isDone := false
	startIndex := index
	lastIndex := index
	for !isDone {
		index = indexNewline(buf, index+1)
		if (index - lastIndex) <= 1 {
			isDone = true
		} else {
			lastIndex = index
		}
	}
	strippedBuf.Write(buf[startIndex:index])
	return strippedBuf.String(), string(buf[index:])
}

// This returns the current stack trace string.  NOTE: the stack creation code
// is excluded from the stack trace.
func StackTrace() (current, context string) {
	return stackTrace(3)
}
