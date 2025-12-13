package main


//一些结构体
type CreateTodoResquest struct {
	Title string `json:"title"`
	Description  string `json:"description"`
}


type UpdateTodoRequest struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed *bool `json:"completed,omitempty"`
}