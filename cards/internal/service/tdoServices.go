package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/looksaw2/ai-agent-with-go/cards/internal/repository"
	"github.com/looksaw2/ai-agent-with-go/cards/model"
)


type TodoService struct {
	repo repository.Repository 	
}

//New Service
func NewTodoSevice(repo repository.Repository) *TodoService{
	return &TodoService{
		repo: repo,
	}
}



//验证CreateRequest
func (s *TodoService) validateCreateRequest(req *model.CreateTodoRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	title := strings.TrimSpace(req.Title)
	if title == "" {
		return fmt.Errorf("title is required")
	}

	if len(title) > 255 {
		return fmt.Errorf("title cannot exceed 255 characters")
	}

	return nil
}

func (s *TodoService) validateUpdateRequest(req *model.UpdateTodoRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	if req.Title == nil && req.Description == nil && req.Completed == nil {
		return fmt.Errorf("at least one field must be provided for update")
	}

	if req.Title != nil {
		title := strings.TrimSpace(*req.Title)
		if len(title) > 255 {
			return fmt.Errorf("title cannot exceed 255 characters")
		}
	}

	return nil
}

//
func (s *TodoService) CreateTodo(req *model.CreateTodoRequest) (*model.Todo, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return nil, err
	}
	ctx := context.Background()

	todo := &model.Todo{
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		Completed:   false,
	}

	if err := s.repo.CreateTodo(ctx,todo); err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return todo, nil
}
func (s *TodoService) GetTodoByID(id int) (*model.Todo, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid todo ID: %d", id)
	}
	ctx := context.Background()

	todo, err := s.repo.GetTodoByID(ctx,id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetAllTodos() ([]*model.Todo, error) {
	todos, err := s.repo.GetAllTodos(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	return todos, nil
}

func (s *TodoService) UpdateTodo(id int, req *model.UpdateTodoRequest) (*model.Todo, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid todo ID: %d", id)
	}

	if err := s.validateUpdateRequest(req); err != nil {
		return nil, err
	}
	ctx := context.Background()
	updates := make(map[string]any)

	if req.Title != nil {
		trimmedTitle := strings.TrimSpace(*req.Title)
		if trimmedTitle == "" {
			return nil, fmt.Errorf("title cannot be empty")
		}
		updates["title"] = trimmedTitle
	}

	if req.Description != nil {
		updates["description"] = strings.TrimSpace(*req.Description)
	}

	if req.Completed != nil {
		updates["completed"] = *req.Completed
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no valid updates provided")
	}

	if err := s.repo.UpdateTodo(ctx,id, updates); err != nil {
		return nil, err
	}

	return s.repo.GetTodoByID(ctx,id)
}

func (s *TodoService) DeleteTodo(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid todo ID: %d", id)
	}
	ctx := context.Background()

	return s.repo.DeleteTodo(ctx,id)
}
