# Todo CLI

A simple command-line To-Do list manager built with Go and Cobra. This CLI allows you to add, list, and mark tasks as completed, helping you to keep track of your to-do list from the terminal.

## Features

- **Add a task**: Add new tasks to your to-do list.
- **List tasks**: View all tasks in your to-do list.
- **Mark tasks as complete**: Mark tasks as completed to track your progress.

## Prerequisites

- Go 1.18 or later installed.
- `go.mod` and `go.sum` are managed by Go Modules.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/arpit529srivastava/todo-cli.git
   ```

2. ```bash
   go mod tidy
   ```

## Usage

### Add a Task

To add a new Task, use the `add` command followed by the task description

```bash
todocli add "Your task description"
```

### List All Tasks

To list all tasks, use the `list` command:

```bash
todocli list
```

### Mark a Task as Completed

To mark a task as completed, use the `complete` command with the task ID:

```
todocli complete [task_id]
```

### Task File

Tasks are saved in the `config/todo.json` file in JSON format. If the file does not exist, it will be created automatically when you first add a task.

### Project Structure

```
todo-cli/
│── cmd/                    
│   ├── add.go              # Command to add a task
│   ├── list.go             # Command to list tasks
│   ├── complete.go         # Command to mark a task as complete
│   ├── root.go             # Root command of the CLI
│── config/                 
│   ├── config.go           # Configuration management for tasks
│   ├── todo.json           # Task storage file (JSON format)
│── main.go                 # Entry point of the application
│── go.mod                  # Go module file
│── go.sum                  # Dependencies file
```
### License
This project is licensed under the MIT License. 