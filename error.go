package conf

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// ErrFileNotFound is returned as a wrapped error by `Load` when the config file is
// not found in the given search dirs.
var ErrFileNotFound = fmt.Errorf("file not found")

// fieldErrors collects errors for fields of config struct.
type fieldErrors map[string]error

// Error formats all fields errors into a single string.
func (fe fieldErrors) Error() string {
	keys := make([]string, 0, len(fe))
	for key := range fe {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.Grow(len(keys) * 10)

	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString(": ")
		sb.WriteString(fe[key].Error())
		sb.WriteString(", ")
	}

	return strings.TrimSuffix(sb.String(), ", ")
}

// Error implements the error interface and can represents multiple
// errors that occur in the course of a single decode.
type Error struct {
	Errors []string
}

func (e *Error) Error() string {
	points := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		points[i] = fmt.Sprintf("* %s", err)
	}

	sort.Strings(points)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(e.Errors), strings.Join(points, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (e *Error) WrappedErrors() []error {
	if e == nil {
		return nil
	}

	result := make([]error, len(e.Errors))
	for i, e := range e.Errors {
		result[i] = errors.New(e)
	}

	return result
}

func appendErrors(errors []string, err error) []string {
	switch e := err.(type) {
	case *Error:
		return append(errors, e.Errors...)
	default:
		return append(errors, e.Error())
	}
}
