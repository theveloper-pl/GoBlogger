package main

import (
  "net/http"
  "database/sql"
  "github.com/gin-gonic/gin"
  "log"
  "fmt"
  "time"
  "main/Blogger/data"
  _ "github.com/lib/pq"
)

const POSTGRES_USER = "postgres"
const POSTGRES_PASSWORD = "password"
const POSTGRES_DB = "blogger"

var db *sql.DB

func main() {
  db = initDB()	
  err := db.Ping()
  if err != nil {
	fmt.Println(err)
  }

  r := gin.Default()
  r.LoadHTMLGlob("templates/*")
  r.Static("/static", "./static/")

  r.GET("", home)

  r.Run() 
}

func home(c *gin.Context){
    // c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	//   })

	models := data.New(db)
	posts, err := models.Post.GetAll()
	// posts, err := model.GetOne(1)
	if err != nil {
		fmt.Println(err)
	  }

	  fmt.Print(posts)
	//   Data: dataMap,
	// c.HTML(http.StatusOK,"home.html",gin.H{"title": "Home Page",},)

	dataMap := make(map[string]any)
	dataMap["posts"] = posts

	c.HTML(http.StatusOK,"home.html",gin.H{"Data": dataMap,},)

}


func initDB() *sql.DB{
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cant connect to database")
	}
	return conn
}

func connectToDB() *sql.DB{
	counts :=0

	for {

		if counts > 5 {
			return nil
		}

		connection, err := openDB()
		if err != nil {
			log.Println("Postgres not yet ready...")
		} else {
			log.Println("Connected to database !")
			return connection
		}
		counts += 1
		time.Sleep(time.Second * 2)
	}

}

func openDB() (*sql.DB, error){
	connStr := fmt.Sprintf("postgresql://%s:%s@0.0.0.0:5432/%s?sslmode=disable",POSTGRES_USER,POSTGRES_PASSWORD, POSTGRES_DB)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}