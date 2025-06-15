package handlers

import (
	"fmt"
	"log"
	"main/internal/models"
	"main/internal/repositories"
	"mime/multipart"

	"main/pkg"
	"net/http"
	fp "path/filepath"

	"time"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	ProfileRepo *repositories.ProfileRepositories
}

func NewProfileHandler(pr *repositories.ProfileRepositories) *ProfileHandler {
	return &ProfileHandler{
		ProfileRepo: pr,
	}
}

func (p *ProfileHandler) GetProfile(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	profile, err := p.ProfileRepo.GetProfileByUserID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to get profile"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: profile})
}

func (p *ProfileHandler) EditProfile(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	var formBody models.ProfileForm
	if err := ctx.ShouldBind(&formBody); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid form data"})
		return
	}

	// handle profile update
	profile := models.Profile{
		User:       userID,
		First_name: formBody.First_name,
		Last_name:  formBody.Last_name,
		Phone:      formBody.Phone,
		Email:      formBody.Email,
	}

	// Handle image upload
	var filename, filepath string
	if formBody.Image != nil {
		var err error
		filename, filepath, err = fileHandling(ctx, formBody.Image, userID)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Error uploading image"})
			return
		}
		profile.Image = filename
	} else {
		oldProfile, err := p.ProfileRepo.GetProfileByUserID(ctx.Request.Context(), userID)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch existing profile"})
			return
		}
		profile.Image = oldProfile.Image
	}

	err := p.ProfileRepo.UpdateProfile(ctx.Request.Context(), profile)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to update profile"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Profile updated", Data: gin.H{
		"filename": filename,
		"path":     filepath,
	}})
}

func (p *ProfileHandler) ChangePassword(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid input"})
		return
	}

	// find user data
	user, err := p.ProfileRepo.GetProfileByUserID(ctx.Request.Context(), userID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to retrieve user"})
		return
	}

	hash := pkg.InitHashConfig()
	valid, err := hash.CompareHashAndPass(user.Password, input.OldPassword)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}
	if !valid {
		ctx.JSON(http.StatusUnauthorized, models.Response{Msg: "Old password is incorrect"})
		return
	}

	// hash new password and update
	hash.UseConfigDefault()
	newHashed, err := hash.GenHashedPassword(input.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to hash new password"})
		return
	}

	err = p.ProfileRepo.UpdatePassword(ctx.Request.Context(), userID, newHashed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to update password"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Password updated"})
}

func fileHandling(ctx *gin.Context, file *multipart.FileHeader, userID int) (filename, filepath string, err error) {
	ext := fp.Ext(file.Filename)
	filename = fmt.Sprintf("%d_%d_profile_image%s", time.Now().UnixNano(), userID, ext)
	filepath = fp.Join("public", "img", filename)
	if err = ctx.SaveUploadedFile(file, filepath); err != nil {
		return "", "", err
	}
	return filename, filepath, nil
}
