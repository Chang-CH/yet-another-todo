package usecase

import (
	"yata-api/internal/models"
)

type TaskRepository interface {
    CreateTask(task *models.TaskModel) error
    UpdateTask(id int, task *models.TaskModel) error
    DeleteTask(id int) error
    GetActiveTasks() ([]*models.TaskModel, error)
    GetTask(id int) (*models.TaskModel, error)
}

type TaskUseCase struct {
    Repo TaskRepository
}

func (u *TaskUseCase) CreateTask(name, email string) error {
    task := models.NewTask(
        name, 
        "description",
        []string{"resource1", "resource2"},
        "notes",
        models.CategoryAdhoc,
        0,
        30,
        models.DayMonday | models.DayTuesday | models.DayWednesday | models.DayThursday | models.DayFriday | models.DaySaturday | models.DaySunday,
        models.NotificationTypeNone,
    )
    return u.Repo.CreateTask(task)
}

func (u *TaskUseCase) GetUser(id int) (*models.TaskModel, error) {
    return u.Repo.GetTask(id)
}