package form

type AccountSignup struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AccountResetPassword struct {
	NewPassword string `json:"newPassword"`
}
