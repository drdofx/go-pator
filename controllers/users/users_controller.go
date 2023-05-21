package users

import (
	"net/http"
	"pator-api/models"
	"pator-api/utils/handlers"
	"pator-api/utils/token"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func Register(c *gin.Context) {
	var input UserRegisterValidation

	if err := c.ShouldBindJSON(&input); err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		NIM:      input.NIM,
		Major:    input.Major,
		Year:     input.Year,
	}

	_, err := user.SaveUser()

	if err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	handlers.HandleSuccess(c, http.StatusCreated, nil)
}

func Login(c *gin.Context) {
	var input UserLoginValidation

	if err := c.ShouldBindJSON(&input); err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	token, err := user.Login(user.Email, user.Password)

	if err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	handlers.HandleSuccess(c, http.StatusOK, gin.H{"token": token})
}

func Profile(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		handlers.HandleError(c, http.StatusUnauthorized, err)
		return
	}

	user, err := models.FindUserByID(user_id)

	if err != nil {
		handlers.HandleError(c, http.StatusNotFound, err)
		return
	}

	handlers.HandleSuccess(c, http.StatusOK, user)
}
