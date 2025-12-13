package repository

import (
	"context"

	"github.com/looksaw2/ai-agent-with-go/cards/model"
)


type Repository interface {
	CreateTodo(ctx context.Context , todo *model.Todo ) error 
	GetTodoByID(ctx context.Context , id int)(*model.Todo ,error)
	GetAllTodos(ctx context.Context)([]*model.Todo ,error)
	UpdateTodo(ctx context.Context ,id int, updates map[string]any)error 
	DeleteTodo(ctx context.Context,id int) error
}