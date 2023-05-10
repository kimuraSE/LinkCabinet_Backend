package model

import "time"

type Link struct {
	ID          uint  `json:"id"`
	UserID      uint  `json:"user_id"`
	Title       string `json:"title"`
	Url 	   string `json:"url"`
	Created_at time.Time `json:"created_at"`
}

type LinkResponse struct {
	ID          uint  `json:"id"`
	Title	   string `json:"title"`
	Url 	   string `json:"url"`
}
