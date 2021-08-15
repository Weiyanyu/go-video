package repository

import (
	"go-video/src/services/user/constant"
	"go-video/src/services/user/db/mysqldb"
	"go-video/src/services/user/entity"
	"time"
)

//ExistUserByUsername determines whether the user(username) is in the database
func ExistUserByUsername(username string) bool {
	var count int64
	mysqldb.GetConnection().
		Table(entity.UserEntity{}.TableName()).
		Where(&entity.UserEntity{Username: username}).
		Count(&count)
	return count > 0
}

//InsertUser insert a user record to database
func InsertUser(username, password string) error {
	now := time.Now()
	user := entity.UserEntity{
		Username:   username,
		Password:   password,
		SignupAt:   &now,
		LastActive: &now,
		Status:     constant.UserStatusAvailable,
	}
	err := mysqldb.GetConnection().
		Table(user.TableName()).Create(&user).Error
	return err
}
