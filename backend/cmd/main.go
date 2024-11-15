package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//initialize db
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB(db)

	// Create a new ServeMux
	mux := http.NewServeMux()

	
	//define routes
	mux.HandleFunc("/api/register", handlers.RegisterUserHandler(db))
	mux.HandleFunc("/api/login", handlers.LoginUserHandler(db))
	mux.Handle("/api/gps", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetGPSDataHandler(db))))
	mux.Handle("/api/getpreferences", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetPreferencesHandler(db))))
	mux.Handle("/api/setpreferences", middleware.AuthMiddleware(http.HandlerFunc(handlers.SetPreferencesHandler(db))))
	
	//create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.EnableCORS(mux),
	}

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
	<-sigChan //block until sig received

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
