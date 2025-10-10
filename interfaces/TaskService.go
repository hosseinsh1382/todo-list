package interfaces

import "ToDoList/Models"

type TaskService interface {
	Add(task Models.Task) (int, error)
	GetAll() ([]*Models.Task, error)
	GetById(id int) (*Models.Task, error)
	Delete(id int) error
}
