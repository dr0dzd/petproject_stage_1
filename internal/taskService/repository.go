package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID uint) ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := r.db.Find(&tasks).Error
	return tasks, result
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var local Task
	//ищем есть ли запись для обновления, если есть пишем ее в local
	if finderr := r.db.First(&local, id).Error; finderr != nil {
		return Task{}, finderr
	}
	//обновляем local
	if uperr := r.db.Model(&local).Updates(&task).Error; uperr != nil {
		return Task{}, uperr
	}
	return local, nil
}

func (r *taskRepository) DeleteTask(id uint) error {
	result := r.db.Unscoped().Delete(&Task{}, id)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
