package controller

import (
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"studygroup/form"
	"studygroup/model"
	"studygroup/util"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

type CreateAccountResponse struct {
	Message string `json:"message"`
}

var accountModel = new(model.Account)

// Create Account
// @Summary Create Account
// @Schemes
// @Description Create Account
// @Tags Register
// @Accept json
// @Produce json
// useing param with [param_name param_type para_data_type is_required comment]
// @Param account body form.AccountSignup true "account info to create"
// @Success 200 {object} CreateAccountResponse
// @Router /register [post]
func (controller AccountController) CreateAccount(ctx *gin.Context) {
	var payload form.AccountSignup
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(payload)

	result, err := accountModel.Create(payload)

	if err != nil {
		if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
			err = fmt.Errorf("account is already exist")
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := CreateAccountResponse{
		Message: fmt.Sprintf("create account success, effect row: %d", result.RowsAffected),
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get Account Info
// @Schemes
// @Description Get Account Info
// @Tags Account
// @produce json
// @Success 200 {object} model.Account
// @Router /authed/account [get]
func (controller AccountController) GetAccountInfo(ctx *gin.Context) {
	id := util.GetSessionUserId(ctx)
	account, err := accountModel.GetFromId(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// @Summary Delete Account
// @Schemes
// @Description Delete Account
// @Tags Account
// @produce json
// @Success 200
// @Router /authed/account [delete]
func (controller AccountController) DeleteAccount(ctx *gin.Context) {
	id := util.GetSessionUserId(ctx)
	_, err := accountModel.GetFromId(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err = accountModel.Delete(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	util.ClearAuthSession(ctx)

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("delete account success")})
}

// @Summary Update Password
// @Schemes
// @Description Update Password
// @Tags Account
// @Accept json
// @produce json
// @Param updatePassword body form.AccountResetPassword true "new password info to update"
//// @Success 200 {object} CreateAccountResponse
// @Success 200
// @Router /authed/account/resetpwd [put]
func (controller AccountController) ResetPassword(ctx *gin.Context) {
	id := util.GetSessionUserId(ctx)
	account, err := accountModel.GetFromId(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var payload form.AccountResetPassword
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(account)

	result, err := accountModel.UpdatePasswd(id, payload.NewPassword)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("reset account success, effect row: %d", result.RowsAffected)})
}

// @Summary Get All Accounts
// @Schemes
// @Description Get All Accounts
// @Tags Account
// @produce application/json
// @Success 200 {array} model.Account
// @Router /authed/account/all [get]
func (controller AccountController) GetAllAccounts(ctx *gin.Context) {
	accounts, err := accountModel.GetAll()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
