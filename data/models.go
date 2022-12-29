package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type Models struct {
	Post Post
}


func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Post: Post{},
	}
}

