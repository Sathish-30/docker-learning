package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type response struct {
	Message string `json:"message"`
}

type User struct {
	UserName string `json:"userName"`
	UserAge  int    `json:"userAge"`
}

func main() {
	for {
		if err := connect(); err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Error in migrating")
	}
	fmt.Println("Database connected")
	const PORT int = 8000
	server := http.NewServeMux()
	server.HandleFunc("GET /", handleHomeRoute)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), middleware(server))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application.json")
		next.ServeHTTP(w, r)
	})
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&response{
		Message: "In Home route",
	})
}

func connect() error {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}
