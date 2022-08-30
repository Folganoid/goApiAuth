package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type User struct {
	ID 					int 	`json:"id"`
	Username			string 	`json:"username"`
	Email 				string 	`json:"email"`
	Password			string	`json:"password,omitempty"`
	HashPassword		string	`json:"hash_password,omitempty"`
	RegisterAt			time.Time`json:"register_at"`
	Role				Role 	`json:"role"`
	Notice				string	`json:"notice"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		u.HashPassword = GetMD5Hash(u.Password)
	}
	u.Password = ""
	u.RegisterAt = time.Now()
	defaultRole := Role{ID: 5}
	u.Role = defaultRole
	u.Notice = "new user"

	return nil
}

func (u *User) BeforeUpdate() error {
	if len(u.Password) > 0 {
		u.HashPassword = GetMD5Hash(u.Password)
	}
	u.Password = ""
	u.Notice = "new user"

	return nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}