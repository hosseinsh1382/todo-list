package Repositories

import (
	"ToDoList/Interfaces"
	"ToDoList/Models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresTaskRepository struct {
	db *gorm.DB
}

func NewPostgresTaskRepository() (Interfaces.TaskRepository, error) {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=todo sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	if err := db.AutoMigrate(&Models.Task{}); err != nil {
		return nil, fmt.Errorf("failed to migrate tasks table %v", err)
	}

	return &PostgresTaskRepository{db: db}, nil
}

func (p PostgresTaskRepository) Add(newTask *Models.Task) (int, error) {
	result := p.db.Create(&newTask)
	if result.Error != nil {
		return 0, fmt.Errorf("error adding task to database: %w", result.Error)
	}
	fmt.Println("Rows affected:", result.RowsAffected)
	return newTask.ID, nil
}

func (p PostgresTaskRepository) GetAll() ([]*Models.Task, error) {
	var tasks []*Models.Task
	result := p.db.Find(&tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("error in getting all tasks: %w", result.Error)
	}
	return tasks, nil
}

func (p PostgresTaskRepository) GetById(id int) (*Models.Task, error) {
	var task Models.Task
	result := p.db.First(task, id)
	if result.Error != nil {
		return nil, fmt.Errorf("error in getting task by id: %w", result.Error)
	}
	return &task, nil
}

func (p PostgresTaskRepository) Delete(id int) error {
	var task Models.Task
	result := p.db.First(&task, id)
	if result.Error != nil {
		return fmt.Errorf("error in getting task by id: %w", result.Error)
	}

	task.Deleted = true
	result = p.db.Save(&task)
	if result.Error != nil {
		return fmt.Errorf("error in deleting task: %w", result.Error)
	}
	return nil

}

/*
func (repo *PostgresTaskRepository) GetAll() ([]*Models.Task, error) {
	err := repo.db.Ping()
	if err != nil {
		log.Printf("Error connecting to Postgres")
		return nil, err
	}

	tasksRow, err := repo.db.Query("SELECT * FROM Tasks")
	if err != nil {
		log.Printf("Error occured in querying tasks table %s", err)
		return nil, err
	}
	allTasks := []*Models.Task{}

	for tasksRow.Next() {
		var task Models.Task
		err := tasksRow.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt, &task.StartAt, &task.EndAt, &task.Deleted)
		if err != nil {
			log.Printf("Error occured in querying tasks table %s", err)
			return nil, err
		}
		allTasks = append(allTasks, &task)
	}

	return allTasks, nil
}
func (repo *PostgresTaskRepository) GetById(id int) (*Models.Task, error) {
	err := repo.db.Ping()
	if err != nil {
		log.Printf("Error connecting to Postgres %s", err)
		return nil, err
	}

	row, err := repo.db.Query("SELECT * FROM Tasks WHERE id=?", id)
	if err != nil {
		log.Printf("Error occured in querying tasks table %s", err)
		return nil, err
	}
	var task *Models.Task
	err = row.Scan(&task.Id, task.Title, &task.Description, &task.CreatedAt, &task.StartAt, &task.EndAt)

	if err != nil {
		log.Printf("Error occured in querying tasks table %s", err)
		return nil, err
	}

	return task, nil
}

func (repo *PostgresTaskRepository) Add(t *Models.Task) (int, error) {
	err := repo.db.Ping()
	if err != nil {
		log.Printf("Error connecting to Postgres %s", err)
	}

	result, err := repo.db.Exec("INSERT INTO tasks (title,description,start_at,end_at,created_at) VALUES (?,?,?,?,?)", t.Title, t.Description, t.StartAt, t.EndAt, t.CreatedAt)

	if err != nil {
		return 0, fmt.Errorf("error occured in adding task %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error occured in querying tasks table %v", err)
	}
	return int(id), nil
}*/
