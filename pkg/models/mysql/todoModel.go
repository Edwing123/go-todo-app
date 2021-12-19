package mysql

import (
	"database/sql"
	"errors"

	"github.com/Edwing123/todo-app/pkg/models"
)

// Represents the todo table that contains methods
// to work with it.
type TodoModel struct {
	DB *sql.DB
}

func (tm *TodoModel) Insert(userId int, todo models.Todo) error {
	// I have to begin a transaction because I'll be
	// inserting data into multple tables.
	tx, err := tm.DB.Begin()
	if err != nil {
		return err
	}

	// First insert the todo in the todo table.
	result, err := tx.Exec(insertTodo, todo.Title, todo.CreatedAt, todo.Expires)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Get the id of the new todo.
	todoId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert the user id and the tod id to the user_todo table.
	_, err = tx.Exec(insertUserTodo, userId, todoId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (tm *TodoModel) GetAll(userId int) ([]*models.Todo, error) {
	// The slice of todos.
	var todos []*models.Todo

	// Query the database to get all todos.
	result, err := tm.DB.Query(selectAllTodos, userId)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	// Go through every row and create a todo based on its
	// return values.
	for result.Next() {
		todo := &models.Todo{}

		err := result.Scan(&todo.Id, &todo.Title, &todo.Done, &todo.CreatedAt, &todo.Expires)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	// Check for any erro while iterating the todos.
	err = result.Err()
	if err != nil {
		return nil, err
	}

	// If everything is fine, return the slice
	// of todos and a nil error.
	return todos, nil
}

func (tm *TodoModel) MarkAsComplete(userId, todoId int) error {
	tx, err := tm.DB.Begin()
	if err != nil {
		return err
	}

	// Get the todo by its id and the id of the user.
	_, err = tx.Exec(selectUserTodo, userId, todoId)
	if err != nil {
		tx.Rollback()

		// If it doesn't exists, return a no record error.
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoRecords
		}

		//eye: Syntax errors or maybe the database is down, anyways, I'm not
		//eye: handling those errors.
		return err
	}

	// If there exist a todo, set the done column to TRUE.
	_, err = tx.Exec(updateTodoToDone, todoId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (tm *TodoModel) Delete(userId, todoId int) error {
	tx, err := tm.DB.Begin()
	if err != nil {
		return err
	}

	// First try to delete the todo from the user_todo table
	// using the user_id and todo_id.
	_, err = tx.Exec(deleteUserTodo, userId, todoId)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Now delete the todo from the todo table.
	_, err = tx.Exec(deleteTodo, todoId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
