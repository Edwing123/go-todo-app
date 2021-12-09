# Server informacion

## Routes descriptions

| Path           | Method | Handler        | Description                   |
| :------------- | :----- | :------------- | :---------------------------- |
| /              | Get    | home           | Display home page             |
| /register      | Get    | registerForm   | Display sign up form page     |
| /auth/register | Post   | registerUser   | Register user                 |
| /login         | Get    | loginForm      | Display login form page       |
| /auth/login    | Post   | loginUser      | Authenticate user             |
| /todos/list    | Get    | todosManager   | Display todos manager page    |
| /todos/create  | Get    | createTodoForm | Display create todo form page |
| /todos/create  | Post   | createTodo     | Create new todo               |
| /auth/logout   | Post   | logoutUser     | Logout user                   |
