package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/vikas/config"
	"github.com/vikas/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckByEmail(e string) (*models.Admin, error) {
	db := config.DB
	var user models.Admin
	err := db.Where(&models.Admin{Email: e}).Find(&user).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func CheckEmployeePassword(psd, hsh string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hsh), []byte(psd))
	return err == nil
}
