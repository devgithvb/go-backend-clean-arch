package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/bind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/sanitize"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"go.uber.org/zap"
)

func (u *UserHandler) CreateTask(c echo.Context) error {
	// Bind
	var req = userdto.CreateTaskRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			echo.Map{
				"status":  false,
				"message": message.ErrorMsg400BadRequest,
				"errors":  bind.CheckErrFromBind(err).Error(),
			},
		)
	}

	// Validation
	if fieldsErrs, err := u.userValidator.ValidateCreateTaskRequest(req); err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  fieldsErrs,
			},
		)
	}

	// Sanitize
	sanitize.New().
		SetPolicy(sanitize.StrictPolicy).
		Struct(&req) // Check err

	// Usage Use-case
	resp, err := u.userInteractor.CreateTask(req)
	if err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		u.app.Logger.Set().Named("users").Error("create-task", zap.Any("error", err.Error()))
		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})
	}
	return c.JSON(http.StatusOK,
		echo.Map{
			"status":  true,
			"message": message.MsgUserCreateTaskSuccessfully,
			"data":    resp,
		},
	)
}
