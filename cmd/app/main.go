package main

import (
	"Golang/internal/database"
	"Golang/internal/handlers"
	"Golang/internal/taskService"
	"Golang/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
