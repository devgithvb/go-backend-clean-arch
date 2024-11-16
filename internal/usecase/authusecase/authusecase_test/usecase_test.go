package authusecase_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

func TestCreateToken(t *testing.T) {
	config := authusecase.Config{
		AccessTokenSecret:      "TOKENSECRET",
		RefreshTokenSecret:     "REFRESHSECRET",
		AccessTokenSubject:     "as",
		RefreshTokenSubject:    "rs",
		AccessTokenExpiryTime:  7 * time.Hour,
		RefreshTokenExpiryTime: 120 * time.Hour,
	}

	u := entity.User{
		ID:        7,
		Name:      "John",
		Mobile:    "09111111111",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	ai := authusecase.New(config)

	at, _ := ai.CreateAccessToken(userauthservicedto.CreateTokenRequest{User: u})
	t.Log(at)

	rt, _ := ai.CreateRefreshToken(userauthservicedto.CreateTokenRequest{User: u})
	t.Log(rt)
}
