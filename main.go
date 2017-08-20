// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/TheoRev/go_web/rest/handlers"
// 	"github.com/TheoRev/go_web/rest/models"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	mux := mux.NewRouter()
// 	models.SetDefaultUser()

// 	mux.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("GET")
// 	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
// 	mux.HandleFunc("/api/v1/users", handlers.CreateUser).Methods("POST")
// 	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
// 	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

// 	log.Println("El servidor esta a la escucha en el puerto :3000")
// 	log.Fatal(http.ListenAndServe(":3000", mux))
// }
