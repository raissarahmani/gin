package handlers

import (
	"errors"
	"main/internal/models"
	"main/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) RegistUser(ctx *gin.Context) {
	var user models.Users
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Input invalid"})
		return
	}

	cmd, err := repositories.UserRepo.AddNewUser(ctx.Request.Context(), user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	exist, err := repositories.UserRepo.IsUserExist(ctx.Request.Context(), user.Email)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
        return
    }
    if exist {
        ctx.JSON(http.StatusConflict, models.Response{Msg: "User already registered, please login"})
        return
    }

	if cmd.RowsAffected() == 0 {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Registration failed"})
        return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "User registered successfully"})
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var input models.Users
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Input invalid"})
		return
	}

	user, err := repositories.UserRepo.FindUserByEmail(ctx.Request.Context(), input.Email)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            ctx.JSON(http.StatusNotFound, models.Response{Msg: "User not found, please register"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
        return
	}

    if user.Password != input.Password {
        ctx.JSON(http.StatusUnauthorized, models.Response{Msg: "Email or password is wrong"})
        return
    }

	ctx.JSON(http.StatusOK, models.Response{Msg: "Login success"})
}