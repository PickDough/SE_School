package main

import (
	"fmt"
	"github.com/Pick-Down/BTC_API/controllers"
	"github.com/Pick-Down/BTC_API/dal"
	"github.com/Pick-Down/BTC_API/middleware"
	"github.com/Pick-Down/BTC_API/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Dependency Injection
	userServ := services.UserService{Repo: &dal.FileRepository{}}
	controllers.UserServ = &userServ
	controllers.BtcServ = &services.BtcService{}
	controllers.AuthServ = &services.AuthService{}

	router := mux.NewRouter()
	router.HandleFunc("/user/create", controllers.Create).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/btcRate", controllers.Rate).Methods("GET")
	http.Handle("/", router)

	router.Use(middleware.JwtMiddleware)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Print(err)
	}
}
