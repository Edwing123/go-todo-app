package mysql

const (
	// Insert queries
	insertUser     = "INSERT INTO user(username, hashed_password) VALUES(?, ?);"
	insertTodo     = "INSERT INTO todo(title, created_at, expires) VALUES(?, ?, ?);"
	insertUserTodo = "INSERT INTO user_todo(user_id, todo_id) VALUES(?, ?);"

	// Select queries
	selectUserForAuth = `SELECT id, hashed_password FROM user
    WHERE username = ?;`
	selectUserById = `SELECT id, username FROM user WHERE id = ?;`
	selectAllTodos = "SELECT * FROM todo ORDER BY created_at desc;"
	selectUserTodo = "SELECT todo_id FROM user_todo WHERE user_id = ? AND todo_id = ?;"

	// Update queries
	updateTodoToDone = "UPDATE todo SET done = TRUE WHERE id = ?;"
)
