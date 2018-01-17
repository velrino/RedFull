package models

type User struct  {
	ID        uint   `json:"id"`
	Email string `json:"email"`
	Gravatar  string `json:"gravatar"`	
	Name  string `json:"name "`
	Password  string `json:"password"`
}