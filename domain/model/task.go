package model

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type TaskTitle string

func NewTaskTitle(title string) (TaskTitle, error) {
	if title == "" || utf8.RuneCountInString(title) > 64 {
		return "", fmt.Errorf("task title must be 1-64 characters")
	}

	return TaskTitle(title), nil
}

type TaskStatus string

const (
	TaskStatusToDo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID        UUID
	Title     TaskTitle
	Status    TaskStatus
	CreatedAt time.Time
}

func NewTask(title TaskTitle, created_at time.Time) *Task {
	return &Task{
		ID:        NewUUID(),
		Title:     title,
		Status:    TaskStatusToDo,
		CreatedAt: created_at,
	}
}

type Tasks []*Task
