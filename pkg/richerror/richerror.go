package richerror

import (
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
)

type Op string

type RichError struct {
	op           Op
	wrappedError error
	message      string
	kind         kind.Kind
	meta         map[string]interface{}
}

func (e RichError) Op() Op {
	return e.op
}

func (e RichError) WrappedError() error {
	return e.wrappedError
}

func (e RichError) Error() string {
	if e.message == "" && e.wrappedError != nil {
		return e.wrappedError.Error()
	}

	return e.message
}

func (e RichError) Message() string {
	return e.message
}

func (e RichError) Kind() kind.Kind {
	return e.kind
}

func (e RichError) Meta() map[string]interface{} {
	return e.meta
}

func (e RichError) Get() map[string]interface{} {
	return map[string]interface{}{
		"op":      e.Op(),
		"error":   e.Error(),
		"message": e.Message(),
		"kind":    e.Kind(),
	}
}

type BuilderError struct {
	RichError
}

func New(op Op) BuilderError {
	return BuilderError{RichError{op: op}}
}

func (r BuilderError) WithOp(op Op) BuilderError {
	r.op = op

	return r
}

func (r BuilderError) WithErr(err error) BuilderError {
	r.wrappedError = err

	return r
}

func (r BuilderError) WithMessage(message string) BuilderError {
	r.message = message

	return r
}

func (r BuilderError) WithKind(kind kind.Kind) BuilderError {
	r.kind = kind

	return r
}

func (r BuilderError) WithMeta(meta map[string]interface{}) BuilderError {
	r.meta = meta

	return r
}

func (r BuilderError) Error() string {
	if r.message == "" && r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r BuilderError) Op() Op {
	if r.op != "" {
		return r.op
	}

	var re BuilderError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return ""
	}

	return re.Op()
}

func (r BuilderError) Kind() kind.Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re BuilderError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r BuilderError) WrappedError() error {
	if r.wrappedError != nil {
		return r.wrappedError
	}

	var re BuilderError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return nil
	}

	return re.WrappedError()
}

func (r BuilderError) Message() string {
	if r.message != "" {
		return r.message
	}

	var re BuilderError
	ok := errors.As(r.wrappedError, &re)
	if ok {
		return re.Message()
	}

	if r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return ""
}

func (r BuilderError) Meta() map[string]interface{} {
	if len(r.meta) != 0 {
		return r.meta
	}

	var re BuilderError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return make(map[string]interface{})
	}

	return re.Meta()
}

func (r BuilderError) Build() RichError {
	return r.RichError
}

func Analysis(err error) RichError {
	var richErr BuilderError

	switch {
	case errors.As(err, &richErr):
		return RichError{
			op:           richErr.Op(),
			wrappedError: richErr.WrappedError(),
			message:      richErr.Message(),
			kind:         richErr.Kind(),
			meta:         richErr.Meta(),
		}

	default:
		return RichError{}
	}
}
