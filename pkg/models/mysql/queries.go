package mysql

const (
	// Insert queries
	insertUser = "INSERT INTO user(username, hashed_password) VALUES(?, ?);"

	// Select queries
	selectUserForAuth = `SELECT id, hashed_password FROM user
    WHERE username = ?;`
	selectUserById = `SELECT id, username FROM user WHERE id = ?;`
)
