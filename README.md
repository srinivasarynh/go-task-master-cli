# TaskMaster CLI

A production-ready command-line task management application built with Go.

## Features

- ✅ Create, update, and delete tasks
- 🎯 Priority levels (low, medium, high)
- 📅 Due date management
- 🔄 Status tracking (pending, completed, canceled)
- 📊 Clean table output with emojis
- 🗂️ Filter tasks by status
- ⚠️ Overdue task detection
- 📝 Detailed task information
- 🔧 Configurable via YAML
- 📋 Structured logging

## Installation

### From Source

```bash
git clone https://github.com/yourusername/taskmaster.git
cd taskmaster
go mod tidy
make build
make install
```

### Using Go Install

```bash
go install github.com/yourusername/taskmaster@latest
```

## Usage

### Add a new task

```bash
# Basic task
taskmaster add -t "Complete project documentation"

# Task with description and priority
taskmaster add -t "Review code" -d "Review PR #123" -p high

# Task with due date
taskmaster add -t "Submit report" --due 2024-12-31
```

### List tasks

```bash
# List all tasks
taskmaster list

# List pending tasks
taskmaster list --status pending

# List completed tasks
taskmaster list --status completed

# List overdue tasks
taskmaster list --overdue

# Detailed view
taskmaster list --detailed
```

### Complete a task

```bash
taskmaster complete 1
```

### Update a task

```bash
# Update title
taskmaster update 1 -t "New title"

# Update priority and due date
taskmaster update 1 -p high --due 2024-12-25
```

### Delete a task

```bash
taskmaster delete 1
```

## Configuration

TaskMaster uses a YAML configuration file located at `~/.taskmaster/config.yaml`.

```yaml
database:
  path: ~/.taskmaster/tasks.db

logging:
  level: info
  file: ~/.taskmaster/taskmaster.log
```

### Environment Variables

You can override configuration using environment variables:

- `TASKMASTER_DATABASE_PATH`
- `TASKMASTER_LOGGING_LEVEL`
- `TASKMASTER_LOGGING_FILE`

## Architecture

The application follows Clean Architecture principles:

```
├── cmd/                 # CLI commands (Cobra)
├── internal/
│   ├── config/         # Configuration management
│   ├── database/       # Database connection & migrations
│   ├── domain/         # Business entities
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic layer
│   └── ui/             # User interface helpers
└── pkg/                # Shared packages
```

## Development

### Prerequisites

- Go 1.21 or later
- SQLite3

### Running Tests

```bash
make test
```

### Building

```bash
make build
```

### Development Mode

```bash
make dev  # Requires air for hot reloading
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `make lint` and `make test`
6. Submit a pull request

## License

MIT License - see LICENSE file for details.
```

## Additional Production Files

### .gitignore
```gitignore
# Binaries
build/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE files
.vscode/
.idea/
*.swp
*.swo

# OS files
.DS_Store
Thumbs.db

# Application files
*.db
*.log
config.yaml
```

### docker/Dockerfile
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o taskmaster .

FROM alpine:latest
RUN apk --no-cache add ca-certificates sqlite
WORKDIR /root/

COPY --from=builder /app/taskmaster .
COPY --from=builder /app/configs ./configs

CMD ["./taskmaster"]
```

### .github/workflows/ci.yml
```yaml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Install dependencies
      run: go mod tidy
    
    - name: Run tests
      run: go test -v ./...
    
    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest
    
    - name: Build
      run: go build -v ./...
```
