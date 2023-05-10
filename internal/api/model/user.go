package model

type User struct {
	ID  	uint  `json:"id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID  	uint  `json:"id"`
	Email 	string `json:"email"`
}