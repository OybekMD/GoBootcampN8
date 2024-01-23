package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
}

var users = []User {
	{
		ID: 1,
		Name: "Anvar",
		LastName: "Nazarov",
		Age: 22,
	},
	{
		ID: 2,
		Name: "Urali",
		LastName: "Mirzahmedov",
		Age: 29,
	},
	{
		ID: 3,
		Name: "Tony",
		LastName: "Anderson",
		Age: 20,
	},
}

func main()  {
	router := gin.Default()

	router.POST("/user", CreatUser)
	router.GET("/users", GetUsers)

	err := router.Run("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
}

func CreatUser(c *gin.Context)  {
	var newUser User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	newUser.ID = len(users)+1
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, users)
}

func GetUsers(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, users)
}