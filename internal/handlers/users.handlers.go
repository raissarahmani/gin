package handlers

import (
	"errors"
	"log"
	"main/internal/models"
	"main/internal/repositories"
	"main/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	UserRepo *repositories.UserRepositories
}

func NewUserHandler(ur *repositories.UserRepositories) *UserHandler {
	return &UserHandler{
		UserRepo: ur,
	}
}

func (u *UserHandler) RegistUser(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBind(&user); err != nil {
		log.Println(err.Error())
		if status, msg := passValidationHandling(err); status != 0 {
			ctx.JSON(status, models.Response{Msg: msg})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	// check if user exist
	exist, err := u.UserRepo.IsUserExist(ctx.Request.Context(), user.Email)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}
	if exist {
		ctx.JSON(http.StatusConflict, models.Response{Msg: "User already registered, please login"})
		return
	}

	// hash password
	hash := pkg.InitHashConfig()
	hash.UseConfigDefault()
	hashedPass, err := hash.GenHashedPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Hash failed"})
		return
	}

	cmd, err := u.UserRepo.AddNewUser(ctx.Request.Context(), user.Email, hashedPass)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	if cmd.RowsAffected() == 0 {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Registration failed"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "User registered successfully. Please login"})
}

// @summary 	Login User
// @router 		/users [post]
// @accept 		json
// @param 		body body models.Users true "login information"
// @produce 	json
// @success 	200 {object} models.Response
// @failure 	500 {object} models.Response
// @failure		401 {object} models.Response
func (u *UserHandler) Login(ctx *gin.Context) {
	var input models.Users
	if err := ctx.ShouldBind(&input); err != nil {
		log.Println(err)
		if status, msg := passValidationHandling(err); status != 0 {
			ctx.JSON(status, models.Response{Msg: msg})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	// check if user registed
	user, err := u.UserRepo.FindUserByEmail(ctx.Request.Context(), input.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, models.Response{Msg: "User not found, please register"})
			return
		}
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	// compare hashedPass and inputPass
	hash := pkg.InitHashConfig()
	valid, err := hash.CompareHashAndPass(user.Password, input.Password) // hashedPass first, newPass later
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}
	if !valid {
		ctx.JSON(http.StatusUnauthorized, models.Response{Msg: "Email or password is wrong"})
		return
	}

	// login success, give jwt
	claims := pkg.NewClaims(user.Id, user.Role)
	token, err := claims.GenerateToken()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

func passValidationHandling(err error) (status int, msg string) {
	if strings.Contains(err.Error(), "Field validation") {
		if strings.Contains(err.Error(), "min") {
			return http.StatusBadRequest, "Password should have at least 8 characters"
		}
		return http.StatusBadRequest, "Email and password should be filled"
	}
	return 0, ""
}
