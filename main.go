package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/Blogger/data"
	"net/http"
)

const POSTGRES_USER = "postgres"
const POSTGRES_PASSWORD = "password"
const POSTGRES_DB = "blogger"

var DB *gorm.DB

const dsn = "host=localhost user=postgres password=password dbname=blogger port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func setup() {
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	db.AutoMigrate(&data.User{}, &data.Post{})
}

func main() {
	setup()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", home)
	r.GET("/posts", posts)
	r.GET("/posts/page/:page", posts)

	r.Run()
}

func home(c *gin.Context) {


	rows, err := DB.Table("posts").Select("posts.id, users.first_name, users.last_name, posts.title, posts.description, posts.short, posts.created_at").Joins("left join users on users.id = user_id").Rows()
	  if err != nil {
		fmt.Println(err)
	} 

	post := data.UserPost{}
	allposts := []data.UserPost{}

	for rows.Next() {
	
		err := rows.Scan(&post.ID, &post.UserFirstName, &post.UserLastName, &post.Title, &post.Description, &post.Short, &post.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		allposts = append(allposts, post)
		
	  }


	c.HTML(http.StatusOK, "home.html", gin.H{"post": allposts[0], "title": "IT", "short": "Why its so cool ?"})

}

func posts(c *gin.Context) {
	rows, err := DB.Table("posts").Select("posts.id, users.first_name, users.last_name, posts.title, posts.description, posts.short, posts.created_at").Joins("left join users on users.id = user_id").Rows()
	  if err != nil {
		fmt.Println(err)
	} 

	post := data.UserPost{}
	allposts := []data.UserPost{}

	for rows.Next() {
	
		err := rows.Scan(&post.ID, &post.UserFirstName, &post.UserLastName, &post.Title, &post.Description, &post.Short, &post.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		allposts = append(allposts, post)
		
	  }


	c.HTML(http.StatusOK, "posts.html", gin.H{"posts": allposts, "title": "Programming", "short": "Learn IT with us !"})

}
