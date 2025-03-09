package _interface

import "yata-api/internal/models"

type TaskRepository interface {
    CreateTask(task *models.TaskModel) error
    UpdateTask(id int, task *models.TaskModel) error
    DeleteTask(id int) error
    GetActiveTasks() ([]*models.TaskModel, error)
    GetTask(id int) (*models.TaskModel, error)
}