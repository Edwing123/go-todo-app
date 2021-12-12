package mysql

import (
	"database/sql"

	"github.com/Edwing123/todo-app/pkg/models"
)

// Represents the todo table that contains methods
// to work with it.
type TodoModel struct {
	DB *sql.DB
}

func (tm *TodoModel) Insert() {

}

func (tm *TodoModel) GetAll() ([]*models.Todo, error) {
	var todos []*models.Todo

	result, err := tm.DB.Query(selectAllTodos)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		todo := &models.Todo{}

		err := result.Scan(&todo.Id, &todo.Title, &todo.Done, &todo.CreatedAt, &todo.Expires)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
