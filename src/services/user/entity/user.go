package entity

import "time"

//UserEntity for orm
type UserEntity struct {
	ID             uint64     `gorm:"column:ID"`
	Username       string     `gorm:"column:username"`
	Password       string     `gorm:"column:password"`
	Phone          string     `gorm:"column:phone"`
	Email          string     `gorm:"column:email"`
	Profile        string     `gorm:"column:profile"`
	EmailValidated bool       `gorm:"column:email_validated"`
	PhoneValidated bool       `gorm:"column:phone_validated"`
	LastActive     *time.Time `gorm:"column:last_active"`
	SignupAt       *time.Time `gorm:"column:signup_at"`
	Status         int        `gorm:"column:status"`
}

//TableName for gorm specify table name
func (UserEntity) TableName() string {
	return "user"
}
