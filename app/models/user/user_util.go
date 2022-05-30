package user

import (
	"gohub/pkg/database"
)

// IsEmailExist 判断 email 已被注册
func IsEmailExist(email string) bool  {
	var count int64
	database.DB.Model(User{}).Where("email = ?",email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool  {
	var count int64
	database.DB.Model(User{}).Where("phone = ?",phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string)(userModel User)  {
	database.DB.Where("phone = ?",phone).First(&userModel)
	return
}

//  通过 手机号/email/用户名来获取用户
func GetByMulti(loginID string)(userModel User)  {
	database.DB.Where("phone = ?",loginID).
		Or("emial = ?",loginID).
		Or("name = ? ",loginID).
		First(&userModel)
	return
}