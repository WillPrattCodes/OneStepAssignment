package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//initialize db
	database.InitDB()
	defer database.CloseDB()

	//create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	//define api
	http.HandleFunc("/api/register", handlers.RegisterUserHandler)
	http.HandleFunc("/api/login", handlers.LoginUserHandler)

	//run server in goroutine
	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v\n", err)
		}
	}()

	//shutdown on interrupt sig
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan // Wait for interrupt signal

	log.Println("Shutting down server")
	//create context and cancel func with 5 sec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//shutdown server using context
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
