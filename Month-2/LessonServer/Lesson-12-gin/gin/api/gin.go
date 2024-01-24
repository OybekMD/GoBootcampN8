package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"nt_bootcamp/bootcamp_go/gin/models"
	"nt_bootcamp/bootcamp_go/gin/storage"
	"strconv"
)

func main() {

	router := gin.Default()

	router.POST("/user/create", CreateUser)
	router.GET("/user/getAll", GetAllUser)
	router.GET("/user/getUser/:id", GetUserByID)
	router.POST("/user/update/:id", UpdateUserByID)
	router.GET("/user/delete/:id", DeleteUserByID)

	err := router.Run("localhost:8090")
	if err != nil {
		panic(err)
	}
}

func CreateUser(g *gin.Context) {
	var newUser *models.UserGin

	err := g.BindJSON(&newUser)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	//var respUser models.User

	id := uuid.New().String()

	newUser.ID = id

	respUser, err := storage.CreateUser(newUser)
	if err != nil {
		log.Println("Error while creating user", err)
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}
	g.JSON(http.StatusCreated, respUser)

	//users = append(users, newUser)

}

func GetAllUser(g *gin.Context) {
	page := g.Request.URL.Query().Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error while concert page to int")
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := g.Request.URL.Query().Get("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Error while concert limit to int")
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respUsers, err := storage.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("Error while getting all users")
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}
	g.JSON(http.StatusCreated, respUsers)
}

func GetUserByID(g *gin.Context) {
	userId := g.Param("id")

	user, err := storage.GetUser(userId)
	if err != nil {
		log.Println("Error while getting user by id")
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusCreated, user)
}

func UpdateUserByID(g *gin.Context) {
	userId := g.Param("id")

	var user *models.UserGin
	err := g.BindJSON(&user)
	if err != nil {
		log.Println("Error while bindjson in update user")
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user.ID = userId

	respUser, err := storage.UpdateUserById(user)
	if err != nil {
		log.Println("Error while updating user by id")
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	g.JSON(http.StatusCreated, respUser)
}

func DeleteUserByID(g *gin.Context) {
	userId := g.Param("id")

	err := storage.DeleteUserByID(userId)
	if err != nil {
		log.Println("Error while deleting user by id")
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	g.JSON(http.StatusCreated, "User deleted")
}
