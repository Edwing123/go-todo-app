package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Edwing123/todo-app/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Represents the user table that contains methods
// to work with.
type UserModel struct {
	DB *sql.DB
}

func (um *UserModel) Insert(username string, password string) error {
	// Create hash of the password
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	// Execute insert query.
	_, err = um.DB.Exec(insertUser, username, string(hashed_password))

	// Check for errors:
	// - Duplicate username?
	// - some other error?
	if err != nil {
		var mySQLErr *mysql.MySQLError

		// Is it a MySQL error?
		if errors.As(err, &mySQLErr) {
			// If so, what type?

			// Is it a duplicated entry error for the username,
			// I check its number and message.
			if mySQLErr.Number == 1062 && strings.Contains(mySQLErr.Message, "username") {
				return models.ErrDuplicatedUsername
			}
		}

		// I'm not handling a different error, so It must be something
		// different, for now just throw it.
		return err
	}

	return nil
}
