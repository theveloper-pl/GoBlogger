package data

import (
	"context"
	"log"
	"time"
)

type Post struct {
	ID        	int
	UserID		int
	Title     	string
	Description string
	Short string
	CreatedAt time.Time
}

func (u *Post) GetAll() ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	select 
    	id, 
		user_id,
		title,
		description,
		short,
		created_at
	from 
	    posts 
	order by 
		created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,			
			&post.Description,
			&post.Short,
			&post.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func (u *Post) GetOne(id int) (*Post, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, user_id, title, description, short, created_at
				from posts 
				where id = $1`

	var post Post
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Description,
		&post.Short,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}


	return &post, nil
}




// // Update updates one user in the database, using the information
// // stored in the receiver u
// func (u *User) Update() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	stmt := `update users set
// 		email = $1,
// 		first_name = $2,
// 		last_name = $3,
// 		user_active = $4,
// 		updated_at = $5
// 		where id = $6`

// 	_, err := db.ExecContext(ctx, stmt,
// 		u.Email,
// 		u.FirstName,
// 		u.LastName,
// 		u.Active,
// 		time.Now(),
// 		u.ID,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Delete deletes one user from the database, by User.ID
// func (u *User) Delete() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	stmt := `delete from users where id = $1`

// 	_, err := db.ExecContext(ctx, stmt, u.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // DeleteByID deletes one user from the database, by ID
// func (u *User) DeleteByID(id int) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	stmt := `delete from users where id = $1`

// 	_, err := db.ExecContext(ctx, stmt, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Insert inserts a new user into the database, and returns the ID of the newly inserted row
// func (u *User) Insert(user User) (int, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var newID int
// 	stmt := `insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
// 		values ($1, $2, $3, $4, $5, $6, $7) returning id`

// 	err = db.QueryRowContext(ctx, stmt,
// 		user.Email,
// 		user.FirstName,
// 		user.LastName,
// 		hashedPassword,
// 		user.Active,
// 		time.Now(),
// 		time.Now(),
// 	).Scan(&newID)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return newID, nil
// }
