package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/remiposo/gras/domain/model"
	"github.com/remiposo/gras/domain/repository"
	"github.com/remiposo/gras/infra/dao"
)

// implement domain/repository.TaskRepo
var _ repository.TaskRepo = &TaskRepo{}

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (tr *TaskRepo) Create(ctx context.Context, task *model.Task) error {
	taskDao := &dao.Task{
		ID:        string(task.ID),
		Title:     string(task.Title),
		Status:    string(task.Status),
		CreatedAt: task.CreatedAt,
	}
	sql := `INSERT INTO tasks (id, title, status, created_at)
          VALUES (:id, :title, :status, :created_at);`
	_, err := tr.db.NamedExecContext(ctx, sql, taskDao)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepo) List(ctx context.Context) (model.Tasks, error) {
	taskDaos := []dao.Task{}
	sql := `SELECT id, title, status, created_at FROM tasks`
	if err := tr.db.SelectContext(ctx, &taskDaos, sql); err != nil {
		return nil, err
	}
	tasks := model.Tasks{}
	for _, taskDao := range taskDaos {
		task := &model.Task{
			ID:        model.UUID(taskDao.ID),
			Title:     model.TaskTitle(taskDao.Title),
			Status:    model.TaskStatus(taskDao.Status),
			CreatedAt: taskDao.CreatedAt,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
