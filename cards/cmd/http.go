package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/looksaw2/ai-agent-with-go/cards/internal/service"
	"github.com/looksaw2/ai-agent-with-go/cards/model"
)

type TodoController struct {
	service *service.TodoService
}

func NewTodoController(s *service.TodoService) *TodoController {
	return &TodoController{
		service: s,
	}
}

//
func(h *TodoController)CreateTodo(c *gin.Context){
	var req model.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	m , err := h.service.CreateTodo(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusCreated,m)
}

//
func(h *TodoController)GetAllTodos(c *gin.Context){
	todos ,err := h.service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusCreated,todos)
}

//
func(h *TodoController)GetTodoByID(c *gin.Context){
	id := c.Param("id")
	idn,err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	m,err := h.service.GetTodoByID(idn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusCreated,m)
}

//
func(h *TodoController)UpdateTodo(c *gin.Context){
	id := c.Param("id")
	idn,err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	var req model.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	m,err := h.service.UpdateTodo(idn,&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusCreated,m)
}

//
func(h *TodoController)DeleteTodo(c *gin.Context){
	id := c.Param("id")
	idn,err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	err = h.service.DeleteTodo(idn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusNoContent,"already deleted............")
}
