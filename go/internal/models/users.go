package models

type User struct {
	ID 					int `json:"id"`
	Username			string `json:"username"`
	Email 				string `json:"email"`
	HashPassword		string	`json:"hash_password,omitempty"`
	RegisterAt			int		`json:"register_at"`
	Role				Role 	`json:"role"`
	Notice				string	`json:"notice"`
}