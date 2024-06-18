package userservice

import (
	"errors"
	"go-backend-clean-arch/internal/domain/dto/userdto"
	"go-backend-clean-arch/internal/domain/entity"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/internal/infrastructure/security/bcrypt"
	"go-backend-clean-arch/pkg/message"
)

func (u *UserInteractor) Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	const op = message.OpUserUsecaseRegister

	isUnique, err := u.repository.IsMobileUnique(req.Mobile)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	if !isUnique {
		return userdto.RegisterResponse{},
			richerror.New(op).
				WithErr(errors.New(message.ErrorMsgMobileIsNotUnique)). //nolint:err113
				WithMessage(message.ErrorMsgInvalidInput).
				WithKind(kind.KindStatusBadRequest)
	}

	var user = entity.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}

	encryptedPass, _ := bcrypt.Generate(req.Password, bcrypt.DefaultCost) // Check err
	user.Password = encryptedPass

	createdUser, err := u.repository.Create(user)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	return userdto.RegisterResponse{
		User: userdto.UserInfo{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Mobile:    createdUser.Mobile,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		},
	}, nil
}