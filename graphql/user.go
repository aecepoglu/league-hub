package graphql

import (
	"crypto/md5"
)

type User struct {
	Email string `gorm:"primary_key;index"`
	Password string `gorm:"not null"`
	Phone string
}

func encryptPass(pass string) string {
	h := md5.New()
	h.Write([]byte(pass))
	return string(h.Sum(nil))
}

func (u *User) EMAIL() string {
	return u.Email
}

func (u *User) PHONE() string {
	return u.Phone
}
