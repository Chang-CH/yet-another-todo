package usecase

import (
	"yata-api/internal/models"
	"yata-api/internal/repository"
)

type TaskUseCase struct {
	Repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		Repo: repo,
	}
}

func (u *TaskUseCase) CreateTask(
    writer http.ResponseWriter, request *http.Request,
) error {
    // TODO: get values from request
	name string,
	description string,
	resources string, // json string
	notes string,
	category models.Category,
	frequency int,
	duration int,
	day int,
	notificationType models.NotificationType,

	task := models.NewTask(
		name,
		description,
		resources,
		notes,
		category,
		frequency,
		duration,
		day,
		notificationType,
	)
	_, err := u.Repo.CreateTask(task)
	return err
}

func (u *TaskUseCase) DeleteTask(id int) error {
	return u.Repo.DeleteTask(id)
}

func (u *TaskUseCase) ModifyTask(
	id int,
	name string,
	description string,
	resources string, // json string
	notes string,
	category models.Category,
	frequency int, // nullable: 0
	duration int, // nullable: 0
	day int, // nullable: (all days)
	notificationType models.NotificationType,
) error {
	if day == 0 {
		day = models.DayMonday | models.DayTuesday | models.DayWednesday | models.DayThursday | models.DayFriday | models.DaySaturday | models.DaySunday
	}
	task := models.NewTask(
		name,
		description,
		resources,
		notes,
		category,
		frequency,
		duration,
		day,
		notificationType,
	)
	return u.Repo.UpdateTask(id, task)
}

func (u *TaskUseCase) GetTask(id int) (*models.TaskModel, error) {
	return u.Repo.GetTask(id)
}

func (u *TaskUseCase) GetActiveTasks() ([]*models.TaskModel, error) {
	return u.Repo.GetActiveTasks()
}
