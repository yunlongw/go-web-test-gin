package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (Auth, bool) {
	var auth Auth
	db.Select("id,username").Where(Auth{Username: username, Password: Sh1Md5(password)}).First(&auth)
	if auth.ID > 0 {
		return auth, true
	}
	return auth, false
}

func AddMember(maps map[string]interface{}) error {
	auth := Auth{
		Username: maps["username"].(string),
		Password: Sh1Md5(maps["password"].(string)),
	}

	if err := db.Create(&auth).Error; err != nil {
		return err
	}
	return nil
}

func ExistMemberByUserName(username string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("username=?", username).Find(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return false, nil
	}
	return true, nil
}

func Sh1Md5(pwd string) string {
	hs := sha1.New()
	hs.Write([]byte(pwd))
	return fmt.Sprintf("%x", hs.Sum(nil))
}
