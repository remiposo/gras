package dao

import "time"

type Task struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
