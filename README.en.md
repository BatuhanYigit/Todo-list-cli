# Todo List

A simple todo list application written in Go.

## Version
0.1

## Date
2023-07-20

## Description

This project is a simple todo list application that helps users manage their tasks. The application supports two different database options: sqlite3 or a text file.

## Requirements

To run the application, you must have the following:

* Go 1.18 or later
* A sqlite3 or text file

## Installation

To install the application, run the following command:


go install [github.com/](https://github.com/)[your-username]/todo-list
```

## Usage

To use the application, follow these steps:

1. Select a database.
2. Add a task.
3. List all tasks.
4. Update a task.
5. Delete a task.

## Example

```
# Select a database

1. sqlite3
2. text file

# Select sqlite3 database

sqlite3

# Add a task

Task name: Grocery shopping
Status: todo

# List all tasks

ID | Task | Status
---|---|---|
1 | Grocery shopping | todo

# Update a task

Task ID: 1
New status: in progress

# Delete a task

Task ID: 1

# Task deleted


## Future features

* I plan to add the following features to the application in the future:
    * Improve the user interface
    * Add dates and times to tasks
    * Add reminders to tasks
    * Support multiple users