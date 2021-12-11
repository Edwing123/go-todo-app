package mysql

const (
	insertUser        = "INSERT INTO user(username, hashed_password) VALUES(?, ?);"
	selectUserForAuth = `SELECT id, hashed_password FROM user
    WHERE username = ?;`
)
