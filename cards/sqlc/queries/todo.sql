-- name: CreateTodo :one
INSERT INTO todos(title,description,completed)
VALUES($1,$2,$3)
RETURNING *;

-- name: GetTodoByID :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: UpdateTodo :exec
UPDATE 
    todos
SET 
    title = COALESCE($2,title),
    description =COALESCE($3,description),
    completed = COALESCE($4,completed),
    updated_at = $5
WHERE id = $1;


-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;