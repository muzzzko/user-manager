package user

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	errorpkg "github/user-manager/internal/error"
	"github/user-manager/internal/generated/server/models"
	usrpkg "github/user-manager/internal/generated/server/restapi/operations/user"
	"github/user-manager/tools/logger"
)

type Handler struct {
	userSrv userService
}

func NewHandler(userSrv userService) *Handler {
	return &Handler{
		userSrv: userSrv,
	}
}

func (h *Handler) CreateUser(params usrpkg.CreateUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	ctxLogger := logger.GetFromContext(ctx)

	userID, err := h.userSrv.CreateUser(ctx, *params.Body)
	if err != nil {
		if errors.Is(err, errorpkg.DomainError) {
			ctxLogger.
				WithError(err).
				WithNickname(params.Body.Nickname).
				Warn("invalid data while creating user")

			return usrpkg.NewCreateUserUnprocessableEntity().WithPayload(&models.Error{
				Code:    errorpkg.GetDomainErrCode(ctx, err),
				Message: err.Error(),
			})
		}

		ctxLogger.
			WithError(err).
			WithNickname(params.Body.Nickname).
			Error("fail to create user")

		return usrpkg.NewCreateUserInternalServerError().WithPayload(&models.Error{
			Code:    errorpkg.GetInternalServiceErrCode(),
			Message: errorpkg.InternalServiceError.Error(),
		})
	}

	return usrpkg.NewCreateUserCreated().WithPayload(&usrpkg.CreateUserCreatedBody{ID: userID})
}

func (h *Handler) DeleteUser(params usrpkg.DeleteUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	ctxLogger := logger.GetFromContext(ctx)

	if err := h.userSrv.DeleteUser(ctx, params.UserID); err != nil {
		ctxLogger.
			WithError(err).
			WithUserID(params.UserID).
			Warn("fail to delete user")

		return usrpkg.NewDeleteUserInternalServerError().WithPayload(&models.Error{
			Code:    errorpkg.GetInternalServiceErrCode(),
			Message: errorpkg.InternalServiceError.Error(),
		})
	}

	return usrpkg.NewDeleteUserNoContent()

}

func (h *Handler) UpdateUser(params usrpkg.UpdateUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	ctxLogger := logger.GetFromContext(ctx)

	newUser, err := h.userSrv.UpdateUser(ctx, *params.Body)
	if err != nil {
		if errors.Is(err, errorpkg.DomainError) {
			ctxLogger.
				WithError(err).
				WithUserID(params.Body.ID).
				Warn("invalid data while updating user")

			return usrpkg.NewUpdateUserUnprocessableEntity().WithPayload(&models.Error{
				Code:    errorpkg.GetDomainErrCode(ctx, err),
				Message: err.Error(),
			})
		}

		ctxLogger.
			WithError(err).
			WithUserID(params.Body.ID).
			Error("fail to update user")

		return usrpkg.NewUpdateUserInternalServerError().WithPayload(&models.Error{
			Code:    errorpkg.GetInternalServiceErrCode(),
			Message: errorpkg.InternalServiceError.Error(),
		})
	}

	return usrpkg.NewUpdateUserOK().WithPayload(&newUser)
}
