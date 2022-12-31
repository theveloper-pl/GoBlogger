package main

import (
	"fmt"
	"main/Blogger/data"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	r.GET("/post/:id", post)
	r.GET("/posts/page/:page", posts)

	r.Run()
}

func post(c *gin.Context) {
	postID := c.Param("id")
	id := 0
	if postID != ""{
		id, _ = strconv.Atoi(postID)		
	}

	rows, err := DB.Table("posts").Select("posts.id, users.first_name, users.last_name, posts.title, posts.description, posts.short, posts.created_at").Joins("left join users on users.id = user_id").Where("posts.id = ?", id).Rows()
	  if err != nil {
		fmt.Println(err)
	} 

	post := data.UserPost{}

	for rows.Next() {
	
		err := rows.Scan(&post.ID, &post.UserFirstName, &post.UserLastName, &post.Title, &post.Description, &post.Short, &post.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}

	  }

	fmt.Printf("%#v",post)
	c.HTML(http.StatusOK, "post.html", gin.H{"post": post, "title": "IT", "short": "Why its so cool ?", })

}

func home(c *gin.Context) {
	rows, err := DB.Table("posts").Select("posts.id, users.first_name, users.last_name, posts.title, posts.description, posts.short, posts.created_at").Joins("left join users on users.id = user_id").Limit(2).Rows()
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


	c.HTML(http.StatusOK, "home.html", gin.H{"posts": allposts, "title": "IT", "short": "Why its so cool ?"})

}

func posts(c *gin.Context) {
	perPage := 10
	page := 1
	pageStr := c.Param("page")

	if pageStr != ""{
		page, _ = strconv.Atoi(pageStr)		
	}


	var totalRows int64
	DB.Table("posts").Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(perPage)))
	offset := (page - 1) * perPage
	rows, err := DB.Table("posts").Select("posts.id, users.first_name, users.last_name, posts.title, posts.description, posts.short, posts.created_at").Joins("left join users on users.id = user_id").Limit(perPage).Offset(offset).Rows()
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


	c.HTML(http.StatusOK, "home.html", gin.H{"posts": allposts, "title": "Programming", "short": "Learn IT with us !", "pagination":data.PaginationData{NextPage:page+1, PreviousPage: page-1, CurrentPage:page, TotalPages:int(totalPages)+1}})

}
