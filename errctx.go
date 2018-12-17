// Package errctx is a library for structured error.
package errctx

import (
	"fmt"
	"strings"
)

type (
	// Error is a structured error.
	Error struct {
		err    error
		msgs   []string
		fields Fields
	}

	// Fields is structured data of error.
	Fields map[string]interface{}
)

func initFields(fields Fields) Fields {
	if fields == nil {
		return Fields{}
	}
	return fields
}

// Wrap returns an error added fields and msgs.
// err should not be nil.
func Wrap(err error, fields Fields, msgs ...string) Error {
	if e, ok := err.(Error); ok {
		e.msgs = append(e.msgs, msgs...)
		e.fields = initFields(e.fields)
		for k, v := range fields {
			e.fields[k] = v
		}
		return e
	}
	return Error{err: err, msgs: msgs, fields: fields}
}

// Wrapf is a shordhand of combination of Wrap and fmt.Sprintf .
func Wrapf(err error, fields Fields, msg string, a ...interface{}) Error {
	return Wrap(err, fields, fmt.Sprintf(msg, a...))
}

// Cause returns a base error.
func (e Error) Cause() error {
	return e.err
}

// Error returns a message represents error.
func (e Error) Error() string {
	msg := strings.Join(e.msgs, " : ")
	if len(e.msgs) == 0 {
		return e.err.Error()
	}
	return fmt.Sprintf("%s : %s", e.err, msg)
}

// Fields returns structured data of error.
func (e Error) Fields() Fields {
	return e.fields
}

// Msgs returns messages.
func (e Error) Msgs() []string {
	return e.msgs
}
