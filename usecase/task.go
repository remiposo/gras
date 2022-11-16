package usecase

import (
	"context"
	"time"

	"github.com/remiposo/gras/domain/model"
	"github.com/remiposo/gras/domain/repository"
)

type TaskUc struct {
	taskRepo repository.TaskRepo
}

func NewTaskUc(taskRepo repository.TaskRepo) *TaskUc {
	return &TaskUc{
		taskRepo: taskRepo,
	}
}

type TaskAddDto struct {
	Title string
}

func (tuc *TaskUc) Add(ctx context.Context, data TaskAddDto) (string, error) {
	title, err := model.NewTaskTitle(data.Title)
	if err != nil {
		return "", err
	}
	task := model.NewTask(title, time.Now())
	if err := tuc.taskRepo.Create(ctx, task); err != nil {
		return "", err
	}

	return string(task.ID), nil
}

type TaskShowDto struct {
	ID     string
	Title  string
	Status string
}

type TaskListDto []TaskShowDto

func (tuc *TaskUc) List(ctx context.Context) (TaskListDto, error) {
	tasks, err := tuc.taskRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	data := make(TaskListDto, 0, len(tasks))
	for _, task := range tasks {
		d := TaskShowDto{
			ID:     string(task.ID),
			Title:  string(task.Title),
			Status: string(task.Status),
		}
		data = append(data, d)
	}
	return data, nil
}
