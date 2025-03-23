package tracerr

import (
	"fmt"
	"runtime"
	"time"
)

// DefaultCap is a default cap for frames array.
// It can be changed to number of expected frames
// for purpose of performance optimisation.
var DefaultCap = 20

// Error is an error with stack trace.
type Error interface {
	Error() string
	StackTrace() []Frame
	Unwrap() error
}

type errorData struct {
	// error category|kind
	kind string
	// timestamp format string
	time string
	// err contains original error.
	err error
	// frames contains stack trace of an error.
	frames []Frame
}

// NewError creates an error with provided frames.
func NewError(kind string, err error, frames []Frame) Error {
	t := time.Now().UTC().Format(time.RFC3339Nano)
	return &errorData{
		kind:   kind,
		time:   t,
		err:    err,
		frames: frames,
	}
}

// Errorf creates new error with stacktrace and formatted message.
// Formatting works the same way as in fmt.Errorf.
func Errorf(kind string, message string, args ...interface{}) Error {
	return trace(kind, fmt.Errorf(message, args...), 2)
}

// New creates new error with stacktrace.
func New(kind string, message string) Error {
	return trace(kind, fmt.Errorf("%s", message), 2)
}

// Wrap adds stacktrace to existing error.
func Wrap(kind string, err error) Error {
	if err == nil {
		return nil
	}
	e, ok := err.(Error)
	if ok {
		return e
	}
	return trace(kind, err, 2)
}

// Unwrap returns the original error.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	e, ok := err.(Error)
	if !ok {
		return err
	}
	return e.Unwrap()
}

// Error returns error message.
func (e *errorData) Error() string {
	return e.err.Error()
}

// StackTrace returns stack trace of an error.
func (e *errorData) StackTrace() []Frame {
	return e.frames
}

// Unwrap returns the original error.
func (e *errorData) Unwrap() error {
	return e.err
}

// Frame is a single step in stack trace.
type Frame struct {
	// Func contains a function name.
	Func string
	// Line contains a line number.
	Line int
	// Path contains a file path.
	Path string
}

// StackTrace returns stack trace of an error.
// It will be empty if err is not of type Error.
func StackTrace(err error) []Frame {
	e, ok := err.(Error)
	if !ok {
		return nil
	}
	return e.StackTrace()
}

// String formats Frame to string.
func (f Frame) String() string {
	return fmt.Sprintf("%s:%d %s()", f.Path, f.Line, f.Func)
}

func trace(kind string, err error, skip int) Error {
	t := time.Now().UTC().Format(time.RFC3339Nano)
	frames := make([]Frame, 0, DefaultCap)
	for {
		pc, path, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		frame := Frame{
			Func: fn.Name(),
			Line: line,
			Path: path,
		}
		frames = append(frames, frame)
		skip++
	}
	return &errorData{
		kind:   kind,
		time:   t,
		err:    err,
		frames: frames,
	}
}

// AS

// IS
