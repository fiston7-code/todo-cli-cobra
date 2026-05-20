Markdown

# Todo-CLI 🚀

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Built with Cobra](https://img.shields.io/badge/Built%20with-Cobra-1e77b4?style=for-the-badge)](https://github.com/spf13/cobra)

A powerful, production-ready Command-Line Interface (CLI) todo list manager built from scratch in Go using the **Cobra** framework. This project follows clean engineering principles with a strict separation of concerns (Models, JSON Storage, and CLI Layer).

---

## 🛠️ Features

* **Persistent Storage**: All your tasks are safely encoded and saved in a formatted `data/tasks.json` file.
* **Auto-Incrementing IDs**: Tasks automatically receive a unique ID tracking system.
* **Dynamic Flags**: Support for short/long flags (`--desc` or `-d`) to add deep descriptions.
* **Smart Auto-generated Help**: Thanks to Cobra, the application automatically builds its own documentation.

---

## 🚀 How to Run (Development)

To run the application directly using the Go compiler without building a binary:

```bash
# Display the general help menu
go run cmd/todo/main.go --help

# Display help for a specific command
go run cmd/todo/main.go list --help

📦 Compilation & Production Binary

If you want to compile the project into a fast, standalone executable binary (perfect for distribution or putting it on your server):
Bash

# Compile the binary
go build -o todo cmd/todo/main.go

# Run the compiled binary directly
./todo list

⌨️ Usage & Examples
1. List all tasks
Bash

go run cmd/todo/main.go list
# Or with binary: ./todo list

2. Add a new task (with description flag)
Bash

go run cmd/todo/main.go add "Deploy KALKAN TECH Landing Page" -d "Push the new responsive design to the production VPS"

3. Delete a task by its ID
Bash

go run cmd/todo/main.go delete 1

📂 Architecture
Plaintext

├── cmd
│   └── todo
│       └── main.go       # CLI Commands & App Orchestration (Cobra)
├── models
│   └── task.go         # Domain Models & In-Memory Store Logic
├── storage
│   └── json.go         # JSON Serialization / Disk Input-Output Operations
└── data
    └── tasks.json      # Your local database (auto-generated)

Designed with 💡 by KALKAN TECH.