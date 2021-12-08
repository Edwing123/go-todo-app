# Data models

In this document I'll be writing things about the data access layer, like the models for the data base, the fields they will have, and the different relationships.

At first I want to start off with just two models/schemas:

- User
- Todo
- UserTodo

## User model

The user model represents user informacion:

- id (pk)
- username
- hashed_password

# Todo model

The todo model represents a todo (a task that will be completed, hopefully):

- id (pk)
- title
- created_at
- expires

## UserTodo model

This model represents the relationship between a user and a todo:

- user_id (fk)
- todo_id (fk)

Each user may have zero or more todos.
