package model

type User struct {
	ID  	int64  `json:"id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID  	int64  `json:"id"`
	Email 	string `json:"email"`
}