package dao

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}
