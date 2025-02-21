package userService

import (
	"Golang/internal/taskService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email" gorm:"unique;not null"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
