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

// Authenticate user by its username and password,
// if the credentials are fine, then return the id
// of the user, otherwise return with an error.
func (um *UserModel) Authenticate(username string, password string) (int, error) {
	var id int
	var hashed_password []byte

	err := um.DB.QueryRow(selectUserForAuth, username).Scan(&id, &hashed_password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		}

		return 0, err
	}

	// Compare the user password and the hashed password from
	// the database, if they do not match return credentials error.
	err = bcrypt.CompareHashAndPassword(hashed_password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}

		return 0, err
	}

	// Return the user id and a nil error if
	// the provided credentials are correct.
	return id, nil
}

func (um *UserModel) GetById(id int) (*models.User, error) {
	user := &models.User{}

	err := um.DB.QueryRow(selectUserById, id).Scan(&user.Id, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecords
		}

		return nil, err
	}

	return user, nil
}
