package taskvalidator

import "github.com/saeedjhn/go-backend-clean-arch/internal/delivery/http/handler/taskhandler"

type Validator struct {
}

var _ taskhandler.Validator = (*Validator)(nil)

func New() Validator {
	return Validator{}
}
