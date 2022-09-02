package models

import (
	"math/rand"
	"strings"
	"time"
)

type Token struct {
	ID 					int 		`json:"id"`
	User				User 		`json:"user"`
	Token 				string 		`json:"token"`
	CreatedAt			time.Time	`json:"created_at"`
	ExpiredAt			time.Time	`json:"expired_at"`
	IsValid				bool		`json:"is_valid"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
