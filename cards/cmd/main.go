package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/looksaw2/ai-agent-with-go/cards/config"
	
	"github.com/looksaw2/ai-agent-with-go/cards/internal/repository"
	"github.com/looksaw2/ai-agent-with-go/cards/internal/service"
)

type application struct {
	config *config.Config
	handler *TodoController
	service *service.TodoService
	repo repository.Repository
}

func NewApplication() *application {
	cfg := config.Load()
	repo,err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		slog.Error("init repo error","error",err)
		os.Exit(1)
	}
	svc := service.NewTodoSevice(repo)
	handler := NewTodoController(svc)
	return &application{
		config: cfg,
		handler: handler,
		service: svc,
		repo: repo,
	}
}

func Run() error {
	app := NewApplication()
	srv := http.Server{
		Addr: app.config.Port,
		Handler: app.NewRoute(),
		IdleTimeout: 1 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	slog.Info("Run http server on","Port",app.config.Port)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main(){
	_ =Run()
}
