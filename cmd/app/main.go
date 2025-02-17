package main

import (
	"Golang/internal/database"
	"Golang/internal/handlers"
	"Golang/internal/taskService"
	"Golang/internal/userService"
	"Golang/internal/web/tasks"
	"Golang/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&taskService.Task{})

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskServ := taskService.NewService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskServ)

	userRepo := userService.NewUserRepository(database.DB)
	userServ := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userServ)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
