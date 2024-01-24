package main

import (
	"encoding/json"
	"fmt"
	"handler_test/models"
	"handler_test/storage"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	// Requests
	http.HandleFunc("/user/create", CreateUser)
	http.HandleFunc("/user/all", GetAllUsers)
	http.HandleFunc("/user/delete", DeleteUser)
	http.HandleFunc("/user/update", UpdateUser)

	log.Println("Server is running...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error while running server", err)
	}

}

// Update user With Patch and Put
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *models.User

	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("error while unmasheling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respUser, err := storage.UpdateUser(user)
	if err != nil {
		log.Println("error while updating user ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

// Create user Method Post
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *models.User

	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("error while unmasheling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Println("error while getting uuid", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.ID = id.String()

	respUser, err := storage.CreateUser(user)
	if err != nil {
		log.Println("error while creating user ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

// Get users mthod Post
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error while converting page, is not intenger", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := r.URL.Query().Get("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("error while converting page, is not intenger", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := storage.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("error while getting all users smth wrong", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	respBody, err := json.Marshal(user)
	if err != nil {
		log.Println("error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

// Delete user method Delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := storage.DeleteUser(id)
	if err != nil {
		log.Println("error while getting all users smth wrong", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted "))
}