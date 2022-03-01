package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikas/config"
	"github.com/vikas/models"
	gomail "gopkg.in/mail.v2"
)

func (u *profileHandler) UpdateUserPass(email string, password string) (err error) {
	db := config.DB
	user := new(models.Admin)
	ad := new(models.Admin)

	user.Password = ad.Password
	fmt.Println(password, "qwertyuioiuydfghjmdfgh")
	err1 := db.Debug().Model(&ad).Where("email = ?", email).Update("password", password)
	if err1 != nil {
		fmt.Println("err1 is thghng an error", err1)
		return
	}
	db.Update(&user)
	db.Save(&user)

	return err
}

// ResetPass func for resetpass as a admin.
// @Description resetpass as admin.
// @Summary resetpass as admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param password body string true "password"
// @Success 200 {object} models.Admin
// @Security ApiKeyAuth
// @Router /admin/resetPass [post]
func (u *profileHandler) ResetPass(c *gin.Context) {
	var data models.PasswordResetCommand
	// var user = new(models.Admin)
	metadata, err := u.tk.ExtractTokenMetadata(c.Request)
	fmt.Println(err, "error in all")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Error while fetching token",
		})
		return
	}
	userId, err := u.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error while validating token",
		})
		return
	}
	if c.ShouldBindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}
	pass := data.Password

	if !isValid(pass) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Password should have a valid password Like [Atleast 1 Lower, 1 Upper, 1 SpecialChar, 1 number]",
		})
		return
	}

	if data.Password != data.Confirm {
		c.JSON(400, gin.H{"message": "Passwords do not match"})
		c.Abort()
		return
	}

	// HashPass, _ := utils.HashPassword(user.Password)
	// user.Password = HashPass
	fmt.Println(userId)
	fmt.Println(data.Password, "userpasasssssssssss")

	// Update user account
	_err := u.UpdateUserPass(userId, data.Password)
	if _err != nil {
		// Return response if we are not able to update user password
		c.JSON(500, gin.H{"message": "Somehting happened while updating your password try again"})
		c.Abort()
		return
	}
	c.JSON(201, gin.H{"message": "Password has been updated, log in"})
	c.Abort()
	return

}

// ForgotPass func for forgotpass as a admin.
// @Description forgotpass as admin.
// @Summary forgotpass as admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} models.Admin
// @Router /forgotPass [post]
func (u *profileHandler) ForgotPass(c *gin.Context) {
	var data models.ResendCommand
	user := models.Admin{}
	db := config.DB

	if (c.BindJSON(&data)) != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON"})
		c.Abort()
		return
	}

	err := db.Debug().Model(&user).Where("email = ?", data.Email).Take(&user).Error
	fmt.Println(err, "email finind")
	fmt.Println(data.Email, "found")

	if user.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}

	resetToken, _ := u.tk.CreateToken(user.Email)
	fmt.Println(resetToken)

	saveErr := u.rd.CreateAuth(user.Email, resetToken)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}

	x1 := gomail.NewMessage()
	x1.SetHeader("From", "devofgolang@gmail.com")
	x1.SetHeader("To", data.Email)
	x1.SetHeader("Subject", " reset Password")
	x1.SetBody("text/plain", "you can reset your pasword by using this  link : \n http://127.0.0.1:8080/xx/resetpaswd.html?="+resetToken.AccessToken)

	y1 := gomail.NewDialer("smtp.gmail.com", 587, "devofgolang@gmail.com", "Abhishek@123")

	err = y1.DialAndSend(x1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email Sent Successfully!")

	tokens := map[string]string{
		"access_token": resetToken.AccessToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"emai":  "mail sent succrssfully",
		"link":  "you can reset your pasword by using this  link : \n http://127.0.0.1:8080/xx/resetpaswd.html?=" + resetToken.AccessToken,
		"token": tokens,
	})

}
