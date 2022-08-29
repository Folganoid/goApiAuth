package models

type Token struct {
	ID 					int 		`json:"id"`
	User				User 		`json:"user"`
	Token 				string 		`json:"token"`
	CreatedAt			int			`json:"created_at"`
	ExpiredAt			int			`json:"expired_at"`
}
