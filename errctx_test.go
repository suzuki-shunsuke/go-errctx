package errctx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_initFields(t *testing.T) {
	require.NotNil(t, initFields(nil))
	fields := Fields{"foo": "bar"}
	require.Equal(t, fields, initFields(fields))
}

func TestWrap(t *testing.T) {
	data := []struct {
		err       error
		fields    Fields
		msgs      []string
		expFields Fields
		expMsgs   []string
	}{{
		fmt.Errorf("foo"), nil, nil, nil, nil,
	}, {
		Error{err: fmt.Errorf("foo")}, Fields{"foo": "bar"}, nil, Fields{"foo": "bar"}, nil,
	}}
	for _, d := range data {
		e := Wrap(d.err, d.fields, d.msgs...)
		require.NotNil(t, e)
		require.Equal(t, d.expFields, e.Fields())
		require.Equal(t, d.expMsgs, e.Msgs())
	}
}

func TestWrapf(t *testing.T) {
	e := Wrapf(fmt.Errorf("foo"), nil, "failed to %s", "get users")
	e2 := Wrap(fmt.Errorf("foo"), nil, fmt.Sprintf("failed to %s", "get users"))
	require.Equal(t, e, e2)
}

func TestErrorCause(t *testing.T) {
	msg := "foo"
	err := Error{err: fmt.Errorf(msg)}
	e := err.Cause()
	require.NotNil(t, e)
	require.Equal(t, msg, e.Error())
}

func TestErrorError(t *testing.T) {
	err := Wrap(fmt.Errorf("foo"), nil)
	require.Equal(t, "foo", err.Error())
	e := Wrap(err, nil, "bar")
	require.Equal(t, "foo", err.Error())
	require.Equal(t, "foo : bar", e.Error())
}

func TestErrorFields(t *testing.T) {
	data := []struct {
		fields Fields
		exp    Fields
	}{{
		nil, nil,
	}, {
		Fields{"foo": "bar"}, Fields{"foo": "bar"},
	}}
	for _, d := range data {
		err := Error{fields: d.fields}
		require.Equal(t, d.exp, err.Fields())
	}
}

func TestErrorMsgs(t *testing.T) {
	msgs := []string{"foo", "bar"}
	err := Error{err: fmt.Errorf("hello"), msgs: msgs}
	require.Equal(t, msgs, err.Msgs())
}
