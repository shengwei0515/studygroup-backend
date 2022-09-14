package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	//hashPwd, err := model.HashPassword(payload.Password)
	//if err != nil {
	//	ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf(" password hash failed!"))
	//	return
	//}

	//if hashPwd != account.Password {
	//	ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf(" wrong password!"))
	//	return
	//}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(payload.Password))
	if err != nil {
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
