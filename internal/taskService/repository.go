package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	//CreateTask - передаем в функцию task типа Task из orm.go
	//возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	//GetAllTasks - возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	//UpdateTaskByID - передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	//DeleteTask - передаем id для удаления, возращаем только ошибку
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
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
	var local Task
	//находим есть ли такая запись и пишем ее в local
	if finderr := r.db.First(&local, id).Error; finderr != nil {
		return finderr
	}
	//удаляем local из БД
	if delerr := r.db.Delete(&local).Error; delerr != nil {
		return delerr
	}
	return nil
}
