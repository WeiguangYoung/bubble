package models

import (
	"bubble/dao"
)

// User Model
type User struct {
	ID int `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

/*
	User这个Model的增删改查操作都放在这里
 */
// CreateAUser 创建user
func CreateAUser(user *User) (err error){
	err = dao.DB.Create(&user).Error
	return
}

func GetAllUser() (userList []*User, err error){
	if err = dao.DB.Find(&userList).Error; err != nil{
		return nil, err
	}
	return
}

func GetAUser(id string)(user *User, err error){
	user = new(User)
	if err = dao.DB.Debug().Where("id=?", id).First(user).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateAUser(user *User)(err error){
	err = dao.DB.Save(user).Error
	return
}

func DeleteAUser(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return
}

