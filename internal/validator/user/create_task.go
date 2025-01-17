package user

import (
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateCreateTaskRequest(req task.CreateRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Title,
			validation.Required,
			validation.Length(_titleMinLen, _titleMaxLen)),

		validation.Field(&req.Description,
			validation.Required,
			validation.Length(_descMinLen, _descMaxLen)),
	); err != nil {
		var fieldErrors = make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(_opTaskValidatorValidateCreateTaskRequest).WithErr(err).
			WithMessage(_errMsgInvalidInput).
			WithKind(kind.KindStatusUnprocessableEntity)
	}

	return nil, nil //nolint:nilnil // return both the `nil` error and invalid value: use a sentinel error instead (nilnil)
}
