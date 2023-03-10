package users

type UserRegisterValidation struct {
	Email                string `json:"email" binding:"required,email"`
	Name                 string `json:"name" binding:"required"`
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
	NIM                  string `json:"nim" binding:"required"`
	Major                string `json:"major" binding:"required"`
	Year                 int    `json:"year" binding:"required"`
}

type UserLoginValidation struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
