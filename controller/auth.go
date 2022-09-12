package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studygroup/form"
	"studygroup/model"
	"studygroup/util"
)

type AuthController struct{}

type AuthResponse struct {
	Message string `json:"message"`
}

// @Summary Login
// @Schemes
// @Description Login
// @Tags Auth
// @Accept json
// @produce json
// @Param account body form.AccountSignup true "login account"
// @Success 200 {object} AuthResponse
// @Router /auth/login [post]
func (controller AuthController) Login(ctx *gin.Context) {
	if hasSession := util.HasSession(ctx); hasSession == true {
		response := AuthResponse{
			Message: fmt.Sprintf("already login"),
		}
		ctx.JSON(http.StatusOK, response)
		return
	}

	var payload form.AccountSignup
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(payload)

	accountModel := model.Account{}
	account, err := accountModel.Get(payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf(" %s is not exist!", payload.Name))
		return
	}

	if payload.Password != account.Password {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf(" wrong password!"))
		return
	}

	err = util.SaveAuthSession(ctx, account.Id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf(" not able to store session"))
		return
	}

	response := AuthResponse{
		Message: fmt.Sprintf("login success"),
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Logout
// @Schemes
// @Description Logout
// @Tags Auth
// @produce json
// @Success 200 {object} AuthResponse
// @Router /auth/logout [get]
func (controller AuthController) Logout(ctx *gin.Context) {
	util.ClearAuthSession(ctx)

	response := AuthResponse{
		Message: fmt.Sprintf("logout success"),
	}
	ctx.JSON(http.StatusOK, response)
}
