// main entry point
// run >>> go run ./cmd/api
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"task-service/internal/api"
	"task-service/internal/repository"
	"task-service/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// INIT DATABASE
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "tasks.db"
	}	

	log.Printf("starting task-service")
	log.Printf("database path: %s", dbPath)

	repo, err := repository.NewSQLiteRepository(dbPath)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("database initialized")


	taskService := service.NewTaskService(repo)
	taskService.StartWorkers(3)

	log.Println("worker pool started (workers=3)")

	handler := api.NewHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handler.CreateTask)
	mux.HandleFunc("/task", handler.GetTask)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("API running on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Shutdown(shutdownCtx)
	taskService.Shutdown()

	log.Println("graceful shutdown complete")
}
