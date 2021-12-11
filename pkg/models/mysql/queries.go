package mysql

const (
	insertUser = "INSERT INTO user(username, hashed_password) VALUES(?, ?);"
)
