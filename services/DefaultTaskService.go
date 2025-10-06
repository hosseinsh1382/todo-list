package services

import (
	"ToDoList/Models"
	"ToDoList/interfaces"
	"fmt"
)

type DefaultTaskService struct {
	repository interfaces.TaskRepository
}

func NewDefaultTaskService(repository interfaces.TaskRepository) interfaces.TaskService {
	return &DefaultTaskService{
		repository: repository,
	}
}

func (s DefaultTaskService) Add(task Models.Task) (int, error) {
	id, err := s.repository.Add(&task)
	if err != nil {
		return 0, fmt.Errorf("Error in DefaultTaskService - Add:\n %v", err)
	}
	return id, nil
}

func (s DefaultTaskService) GetAll() ([]*Models.Task, error) {
	tasks, err := s.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("Error in DefaultTaskService - GetAll:\n %v", err)
	}
	return tasks, nil
}

func (s DefaultTaskService) GetById(id int) (*Models.Task, error) {
	task, err := s.repository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("Error in DefaultTaskService - GetById(%v):\n %v", id, err)
	}
	return task, nil
}

func (s DefaultTaskService) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error in DefaultTaskService - Delete(%v):\n %v", id, err)
	}
	return nil
}
