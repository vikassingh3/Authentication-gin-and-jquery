package models

type UserModel struct{}

// User represents the model for an user
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	DOB      string `json:"dob"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Password string `json:"password"`
}

// Admin represents the model for an admin

type Admin struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Names    string `json:"names"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Token struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// ResendCommand defines resend email
type ResendCommand struct {
	// We only need the email to initialize an email sendout
	Email string `json:"email" binding:"required"`
}

// PasswordResetCommand defines user password reset form struct
type PasswordResetCommand struct {
	Password string `json:"password" binding:"required"`
	Confirm  string `json:"confirm" binding:"required"`
}
