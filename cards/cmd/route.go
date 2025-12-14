package main

import "github.com/gin-gonic/gin"

func(app *application) NewRoute() *gin.Engine{
	route := gin.Default()
	route.GET("/todos",app.handler.GetAllTodos)
	route.POST("/todos",app.handler.CreateTodo)
	route.GET("/todos/:id",app.handler.GetTodoByID)
	route.PUT("/todos/:id",app.handler.UpdateTodo)
	route.DELETE("/todos/:id",app.handler.DeleteTodo)
	return  route
}