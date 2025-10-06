package Interfaces

import (
	"ToDoList/Models"
)

type TaskRepository interface {
	Add(t *Models.Task) (int, error)
	GetAll() ([]*Models.Task, error)
	GetById(id int) (*Models.Task, error)
	Delete(id int) error
}
