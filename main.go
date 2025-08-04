package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	PORT := ":8085"
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")

	fmt.Println("Server starting on PORT", PORT)
	http.ListenAndServe(PORT, router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/health route called")
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/users route called")
	users := []map[string]string{
		{
			"id":   "1",
			"name": "test user 1",
		},
		{
			"id":   "2",
			"name": "test user 2",
		},
		{
			"id":   "3",
			"name": "test user 3",
		},
		{
			"id":   "4",
			"name": "test user 4",
		},
	}
	userData, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Sending users data")
	w.Write(userData)
}
