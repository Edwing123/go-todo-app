# Server informacion

## Routes descriptions

| Path            | Protected? | Method | Handler            | Description                   |
| :-------------- | :--------- | :----- | :----------------- | :---------------------------- |
| /               | No         | Get    | home               | Display home page             |
| /register       | No         | Get    | registerForm       | Display sign up form page     |
| /auth/register  | No         | Post   | registerUser       | Register user                 |
| /login          | No         | Get    | loginForm          | Display login form page       |
| /auth/login     | No         | Post   | loginUser          | Authenticate user             |
| /todos/list     | Yes        | Get    | todosManager       | Display todos manager page    |
| /todos/create   | Yes        | Get    | createTodoForm     | Display create todo form page |
| /todos/create   | Yes        | Post   | createTodo         | Create new todo               |
| /todos/complete | Yes        | Post   | markTodoAsComplete | Mark a todo as complete       |
| /todos/delete   | Yes        | Post   | deleteTodo         | Delete todo                   |
| /auth/logout    | Yes        | Post   | logoutUser         | Logout user                   |
