package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikas/config"
	"github.com/vikas/models"
	"github.com/vikas/utils"
)

// GetAllUsers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /admin/allUser [get]
func GetAllUsers(c *gin.Context) {
	var user []models.User

	db := config.DB
	// Get all users.
	db.Find(&user)
	// all users from id 3
	// db.Offset(3).Find(&user)

	// db.Limit(10).Offset(5).Find(&user)

	// db.Offset(10).Find(&user).Offset(-1).Find(&user)

	// limit one
	// err1 := db.Limit(1).Find(&user)

	// slice of primary_key
	// db.Where([]int{1, 2}).Find(&user)

	//map find by name
	// db.Where(map[string]interface{}{"name": "sid"}).Find(&user)

	// struct find

	// db.Where(&models.User{Email: "vs@gmail.com"}).Find(&user)

	// Return status 200 OK.
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "successfull",
		"data":    user,
	})
}

// Getuser func gets user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /admin/get/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	db := config.DB
	var user models.User
	// user.Name = userId
	// find user by ID
	db.Find(&user, id)
	// error if user is not found
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "unable to find user by id",
		})
		return
	}
	// return 200 status ok
	c.JSON(http.StatusOK, gin.H{
		// "userid": userId,
		"data":   user,
		"status": true,
	})
}

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /admin/createUser [post]
func CreateUser(c *gin.Context) {
	db := config.DB
	user := new(models.User)
	data := new(models.User)
	// creating a struct to hold the user data
	type UserData struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		DOB     string `json:"dob"`
		Gender  string `json:"gender"`
		Address string `json:"address"`
		Country string `json:"country"`
		State   string `json:"state"`
	}

	var input UserData

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "provide valid json details",
		})
		return
	}
	db.Find(&data, "email = ?", user.Email)
	if data.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email is already exist",
		})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "password is required",
		})
		return
	} else {
		HashPass, _ := utils.HashPassword(user.Password)
		user.Password = HashPass
	}
	// fmt.Println("whehjdsahjdshjd")
	// fmt.Println(userId)
	// required fields
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name field is required",
		})
		return
	} else if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email field is required",
		})
		return
	} else if user.Country == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "country field is required",
		})
		return
	} else if user.Gender == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gender field is required",
		})
		return
	} else if user.State == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "state field is required",
		})
		return
	} else if user.DOB == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "dob field is required",
		})
		return
	}

	input.Name = user.Name
	input.Email = user.Email
	input.DOB = user.DOB
	input.Gender = user.Gender
	input.Address = user.Address
	input.Country = user.Country
	input.State = user.State
	// creating user
	db.Create(&user)
	// return 200 status ok
	c.JSON(http.StatusOK, gin.H{
		"createdUser": input,
		"status":      true,
		"message":     "User Created successfully",
	})
}

// DeleteUser func for deletes user by given ID or error.
// @Description Delete user by given ID.
// @Summary delete user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /admin/delete{id} [delete]
func (h *profileHandler) DeleteUser(c *gin.Context) {
	var user models.User

	db := config.DB
	id := c.Param("id")
	c.ShouldBindJSON(&id)

	db.Find(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Unable to find  id",
		})
		return
	}
	// fmt.Println(id, "id kya hia")

	db.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "successfully deleted",
	})
}

// UpdateUser func for updateuser by given ID or error.
// @Description Update user by given ID.
// @Summary Update user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Update body string true "update"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /admin/update{id} [put]
func UpdateUser(c *gin.Context) {
	var s = new(models.User)

	var data = new(models.User)

	c.ShouldBindJSON(&s)

	id := c.Param("id")
	db := config.DB
	// fmt.Println("userId", userId)
	db.Find(&data, id)
	if data.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Id not found",
		})
		return
	}
	data.Name = s.Name
	data.Email = s.Email
	// data.Gender = s.Gender
	data.Address = s.Address
	data.Password = s.Password
	// data.Country = s.Country
	// data.DOB = s.DOB
	// data.State = s.State
	db.Update(&data, id)
	db.Save(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Updated user",
		"user":    s,
	})
}

// "email": "vs1868860@gmail.com", "password": "Vikas@12345"
