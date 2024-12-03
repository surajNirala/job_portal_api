package models

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type PasswordValidation struct {
	MinLength  int
	HasUpper   bool
	HasLower   bool
	HasNumber  bool
	HasSpecial bool
}
