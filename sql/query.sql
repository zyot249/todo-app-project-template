-- name: GetTodo :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todo
ORDER BY title;

-- name: CreateTodo :one
INSERT INTO todo (
  title, description, created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todo
WHERE id = $1;