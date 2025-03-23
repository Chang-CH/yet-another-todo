package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"yata-api/internal/repository"
	"yata-api/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	repo := repository.NewTaskRepository(dbUser, dbPassword, dbHost, dbPort, dbName)
	defer repo.Close()
	uc := usecase.NewTaskUseCase(repo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "CRUD Go Server with pgx is running!")
	})
	//create router
	router := mux.NewRouter()
	router.HandleFunc("/users", uc.CreateTask).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
