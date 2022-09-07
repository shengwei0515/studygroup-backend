package controller

import (
	"fmt"
	"net/http"
	"studygroup/form"
	"studygroup/model"

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
// @Tags Account
// @Accept json
// @Produce json
// useing param with [param_name param_type para_data_type is_required comment]
// @Param account body form.AccountSignup true "account info to create"
// @Success 200 {object} CreateAccountResponse
// @Router /account [post]
func (controller AccountController) CreateAccount(ctx *gin.Context) {
	var payload form.AccountSignup
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	fmt.Println(payload)

	result, err := accountModel.Create(payload)
	response := CreateAccountResponse{
		Message: fmt.Sprintf("create account success, effect row: %d", result.RowsAffected),
	}

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, response)
}
