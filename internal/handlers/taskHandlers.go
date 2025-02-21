package handlers

import (
	"Golang/internal/taskService"
	"Golang/internal/web/tasks"
	"golang.org/x/net/context"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) GetUsersUserIdTasks(_ context.Context, request tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {
	user_id := request.UserId
	userTasks, err := h.Service.GetTasksByUserID(user_id)
	if err != nil {
		return nil, err
	}
	response := tasks.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksTaskId(_ context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	taskID := request.TaskId
	if err := h.Service.DeleteTask(taskID); err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksTaskId204Response{}
	return response, nil
}

func (h *TaskHandler) PatchTasksTaskId(_ context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	taskRequest := request.Body
	taskID := request.TaskId
	TaskToUpdate := taskService.Task{}
	if taskRequest.Task != nil {
		TaskToUpdate.Task = *taskRequest.Task
	}
	if taskRequest.IsDone != nil {
		TaskToUpdate.IsDone = *taskRequest.IsDone
	}
	if taskRequest.UserId != nil {
		TaskToUpdate.UserID = *taskRequest.UserId
	}
	UpdatedTask, err := h.Service.UpdateTaskByID(taskID, TaskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &UpdatedTask.ID,
		Task:   &UpdatedTask.Task,
		IsDone: &UpdatedTask.IsDone,
		UserId: &UpdatedTask.UserID,
	}
	return response, nil
}
