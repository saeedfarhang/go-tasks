# Todo App in Go

A simple command-line Todo application built in Go. Manage your tasks efficiently with features like adding, completing, undoing, deleting, and tracking time spent on tasks.

## Features

- **Add Tasks**: Add new tasks with a title and optional description.
- **Complete Tasks**: Mark tasks as completed.
- **Undo Tasks**: Mark completed tasks as incomplete.
- **Delete Tasks**: Remove tasks from the list.
- **Delete Completed Tasks**: Clear all completed tasks at once.
- **Time Tracking**: Track time spent on tasks (work in progress).
- **Interactive Menu**: Navigate the app using an intuitive menu.
- **Persistent Storage**: Tasks are stored in an SQLite database for persistence.

## Installation

### Prerequisites

- Go (version 1.16 or higher)
- SQLite3

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/todo-app-go.git
   cd todo-app-go
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build the application:

   ```bash
   go build -o todo-app
   ```

4. Run the app:

   ```bash
   ./todo-app
   ```

## Usage

### Run in Interactive Mode

To use the app with an interactive menu:

```bash
./todo-app --interactive
```

### Run in Non-Interactive Mode

To simply list all tasks:

```bash
./todo-app
```

### Commands in Interactive Mode

- **Add Task**: Add a new task with a title and description.
- **Complete Task**: Mark a task as completed by entering its ID.
- **Undo Task**: Mark a completed task as incomplete by entering its ID.
- **Delete Task**: Delete a task by entering its ID.
- **Delete Completed Tasks**: Remove all completed tasks.
- **Work on Task**: Start/stop tracking time for a task (work in progress).
- **Exit**: Exit the application.

## Screenshots

![Todo App Screenshot](screenshot.png) <!-- Add a screenshot if available -->

## Database Schema

The app uses SQLite for persistent storage. The database schema is as follows:

```sql
CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    description TEXT,
    completed BOOLEAN,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE task_times (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER,
    start_time DATETIME,
    end_time DATETIME,
    FOREIGN KEY (task_id) REFERENCES tasks (id)
);
```

## Contributing

Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Go](https://golang.org/) for the awesome programming language.
- [SQLite](https://www.sqlite.org/) for lightweight and efficient database storage.
- [github.com/eiannone/keyboard](https://github.com/eiannone/keyboard) for handling keyboard input in the interactive menu.
