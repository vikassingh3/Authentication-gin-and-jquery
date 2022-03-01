package controllers

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/vikas/auth"
	"github.com/vikas/config"
	"github.com/vikas/models"
	"github.com/vikas/utils"
	"golang.org/x/crypto/bcrypt"
)

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 9 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func CheckPassword(psd, hsh string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hsh), []byte(psd))
	return err == nil
}

// ProfileHandler struct
type profileHandler struct {
	rd auth.AuthInterface
	tk auth.TokenInterface
}

// NewProfileHandler creates a new profile handler
func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface) *profileHandler {
	return &profileHandler{rd, tk}
}

// create user godoc
// @Summary      Register Admin
// @Description  Register Amin
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        RegisterAdmin  body models.Admin  true  "Register Admin"
// @Security 	 BasicAuth
// @Router       /register [post]
func RegisterAdmin(c *gin.Context) {

	db := config.DB
	var data = new(models.Admin)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err, "message": "Error in binding data"})
		return
	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "Email is required"})
		return
	}
	if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "Password is required"})
		return
	}

	email_check, err := utils.CheckByEmail(data.Email)
	if email_check != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err, "message": "Email already exist"})
		return
	}

	hash, err := utils.HashPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err, "message": "Error in hash password"})
		return
	}

	data.Password = hash
	if err := db.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err, "message": "Error in create admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "responseData": data})
}

// Login func for login as a admin.
// @Description Login as admin.
// @Summary login as admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} models.Admin
// @Router /login [post]
func (h *profileHandler) Login(c *gin.Context) {
	var u models.Admin
	db := config.DB
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	// required fields
	if u.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email is required",
		})
		return
	} else if u.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password is required",
		})
		return
	}
	// checking email and password are exists in the database or not
	err := db.Debug().Model(models.Admin{}).Where("email = ? AND password = ?", u.Email, u.Password).Take(&u).Error
	fmt.Println(u.Email, "jfdjfhfjf")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Admin is not valid",
		})
		return
	}
	pass := u.Password
	if !isValid(pass) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
		return
	}
	// creating a token from user.email
	ts, err := h.tk.CreateToken(u.Email)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	// save the token into the redis
	saveErr := h.rd.CreateAuth(u.Email, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}
	// response
	tokens := map[string]string{
		"access_token": ts.AccessToken,
	}
	c.JSON(http.StatusOK, tokens)
}

// Logout func for logout as a admin.
// @Description logout as admin.
// @Summary logout as admin
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} models.Admin
// @Security ApiKeyAuth
// @Router /admin/logout [post]
func (h *profileHandler) Logout(c *gin.Context) {
	//If metadata is passed and the tokens valid, delete them from the redis store
	metadata, _ := h.tk.ExtractTokenMetadata(c.Request)
	if metadata != nil {
		deleteErr := h.rd.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
