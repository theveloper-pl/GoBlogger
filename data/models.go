package data

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    			int	`json:"user_id" gorm:"primary_key"`
	Email     			string `json:"email"`
	FirstName 			string `json:"first_name"`
	LastName  			string `json:"last_name"`
	Password  			string `json:"password"`
	Active    			int	`json:"is_active"`
	IsAdmin   			int	`json:"is_admin"`
	CreatedAt 			time.Time `json:"created_at"`
}

type Post struct {
	gorm.Model
	ID        			int `json:"post_id" gorm:"primary_key"`
	UserID				int `json:"user_id"`
	Title     			string `json:"title"`
	Description 		string `json:"description"`
	Short 				string `json:"short"`
}


type UserPost struct {
	ID        			int
	UserFirstName 		string
	UserLastName 		string
	Title     			string
	Description 		string
	Short 				string
	CreatedAt 			time.Time
}

type PaginationData struct {
	NextPage 		int
	PreviousPage	int
	CurrentPage		int
}