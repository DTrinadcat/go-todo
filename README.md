# go-todo

[![Go Version](https://img.shields.io/github/go-mod/go-version/DTrinadcat/go-todo)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen.svg)]()

A fast, lightweight CLI task manager and reusable in-memory Go package with persistent JSON storage.

## Installation
Make sure you have Go installed, then run:

```bash
go install github.com/DTrinadcat/go-todo/cmd/todo@latest
```

*(Ensure `$GOPATH/bin` is in your system `PATH`).*

## CLI Usage

Once installed, use the `todo` command anywhere in your terminal:

```bash
# Add a new task
todo add "Finish pet-project README"
todo add "Review pull requests"

# View all tasks
todo show

# Delete task by number or UUID
todo delete 1

# Update task text
todo update 1 "Finish and publish pet-project"

# Display help message
todo -h
```

### Example Output:
```text
$ todo show
-----------------------------
Do homework | id: a1b2c3d4-e5f6-7890-abcd-ef1234567890
Buy milk | id: f47ac10b-58cc-4372-a567-0e02b2c3d479
-----------------------------
```

## Using as a Go Library

You can also import `go-todo` into your own Go applications:

```go
package main

import (
    "fmt"
    "github.com/DTrinadcat/go-todo"
)

func main() {
    list := todo.TodoList{}
    
    // Add task
    id, err := list.Add("Buy milk")
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Created task with ID:", id)
}
```
