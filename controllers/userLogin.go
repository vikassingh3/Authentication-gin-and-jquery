package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikas/config"
	"github.com/vikas/models"

	"github.com/jinzhu/gorm"
)

/* func CheckPassword(psd, hsh string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hsh), []byte(psd))
	return err == nil
}
*/
func UserbyEmail(e string) (*models.User, error) {
	db := config.DB
	var user models.User
	err := db.Where(&models.User{Email: e}).Find(&user).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UserbyUsername(u string) (*models.User, error) {
	db := config.DB
	var user models.User

	err := db.Where(&models.User{Name: u}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UserbyUsername1(u string) (*models.Token, error) {
	db := config.DB
	var user models.Token

	err := db.Where(&models.Token{Username: u}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// UserLogin func for login as a user.
// @Description login as user.
// @Summary login as user
// @Tags User
// @Accept json
// @Produce json
// @Param Details body string true "Email"
// @Success 200 {object} models.User
// @Router /admin/userlogin [post]
func (h *profileHandler) UserLogin(c *gin.Context) {

	type LoginInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	var ud UserData

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on login request", "data": err})
		return
	}

	e := input.Email
	use := input.Username
	pass := input.Password

	email, err := UserbyEmail(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on email", "data": err})
		return

	}

	user, err := UserbyUsername(use)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on username", "data": err})
		return
	}

	if email == nil || user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User not found", "data": err})
		return
	}

	if email == nil {
		ud = UserData{
			ID:       uint(user.ID),
			Username: user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		ud = UserData{
			ID:       uint(email.ID),
			Username: email.Name,
			Email:    email.Email,
			Password: email.Password,
		}
	}
	db := config.DB

	if !CheckPassword(pass, ud.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid password", "data": nil})
		return
	}
	db.Model(&ud).Where("password = ?", ud.Password)

	ts, err := h.tk.CreateToken(ud.Username)

	saveErr := h.rd.CreateAuth(ud.Username, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Success login",
		"token":   ts,
	})

}

func LogOut(c *gin.Context) {
	db := config.DB
	var tk1 models.Token

	type LogOutInput struct {
		Username string `json:"username"`
	}

	var input LogOutInput

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Error on logout request", "data": err})
		return
	}

	use := input.Username

	user, err := UserbyUsername1(use)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Error on username", "data": err})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User not valid", "data": err})
		return
	}

	db.Where("username = ?", use).Delete(&tk1)

	c.JSON(http.StatusOK, gin.H{
		"user":    use,
		"message": "logout successfull",
	})
}

/*
func UserChangePassword(c *gin.Context) {
	type ChangePassword struct {
		Username        string `json:"username"`
		OldPassword     string `json:"oldpassword"`
		NewPassword     string `json:"newpassword"`
		ConfNewPassword string `json:"confnewpassword"`
	}

	var input ChangePassword

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Error on login request", "data": err})
		return
	}

	np := input.NewPassword
	cp := input.ConfNewPassword

	use := input.Username
	pass := input.OldPassword

	if pass == np {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "your new password is same as old one !!",
			"data":    nil,
		})
		return
	}

	user, err := UserbyUsername(use)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Error on username", "data": err})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User not found", "data": err})
		return
	}

	var ud = new(models.User)

	db := config.DB
	db.Find(&ud, "username = ? ", use)

	if !CheckPassword(pass, ud.Password) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "current password is not mached with existing password",
			"data":    nil,
		})
		return
	}

	if np != cp {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": " password is not mached !!",
			"data":    nil,
		})
		return

	}

	hash, err := utils.HashPassword(np)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Couldn't hash password", "data": err})
		return

	}

	db.Model(&ud).Where("username = ?", use).Update("password", hash)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Password Update successfull !!",
	})

}
*/
