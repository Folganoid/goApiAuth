package models

type Role struct {
	ID 					int 	`json:"id"`
	Name				string	`json:"name"`
	Level				int 	`json:"level"`
	Notice				string	`json:"notice"`
}
