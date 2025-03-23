package repository

import (
	"database/sql"
	"fmt"
	"log"
	"yata-api/internal/models"
)

type TaskRepository interface {
	CreateTask(task *models.TaskModel) (int, error)
	UpdateTask(id int, task *models.TaskModel) error
	DeleteTask(id int) error
	GetActiveTasks() ([]*models.TaskModel, error)
	GetTask(id int) (*models.TaskModel, error)

	Close() error
}

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(
	dbUser string,
	dbPassword string,
	dbHost string,
	dbPort string,
	dbName string,
) TaskRepository {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)

	// open the database
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// create the Tasks table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS TASKS(
    ID SERIAL PRIMARY KEY,
    NAME TEXT,DESCRIPTION TEXT,
    RESORUCES TEXT,
    NOTES TEXT,
    CATEGORY INT,
    FREQUENCY INT,
    DURATION INT,
    DAY INT,
    IS_ACTIVE BOOLEAN,
    NOTIFICATION_TYPE INT);`) // TODO: validate
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	// create the Tasks statistics table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS TASK_STATISTICS(ID INT PRIMARY KEY,
     LAST_DONE TIMESTAMP,
     LAST_IGNORED TIMESTAMP,
     CREATED_AT   TIMESTAMP,
     UPDATED_AT   TIMESTAMP,
     DONE INT,
     IGNORED INT);`) // TODO: validate
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	return &TaskRepositoryImpl{
		db: db,
	}
}

func (t TaskRepositoryImpl) CreateTask(task *models.TaskModel) (int, error) {
	var id int
	err := t.db.QueryRow(`INSERT INTO TASKS(
    NAME, 
    DESCRIPTION, 
    RESOURCES, 
    NOTES, 
    CATEGORY, 
    FREQUENCY, 
    DURATION, 
    DAY, 
    IS_ACTIVE, 
    NOTIFICATION_TYPE) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING ID;`, // TODO: validate
		task.Name,
		task.Description,
		task.Resources,
		task.Notes,
		int(task.Category),
		task.Frequency,
		task.Duration,
		task.Day,
		task.Active,
		int(task.NotificationType)).Scan(&id)

	if err != nil {
		log.Fatal("Failed to create task: ", err)
	}

	return id, err
}

func (t TaskRepositoryImpl) UpdateTask(id int, task *models.TaskModel) error {
	_, err := t.db.Exec(`UPDATE TASKS SET NAME = $1, 
    DESCRIPTION = $2, 
    RESOURCES = $3, 
    NOTES = $4, 
    CATEGORY = $5, 
    FREQUENCY = $6, 
    DURATION = $7, 
    DAY = $8, 
    IS_ACTIVE = $9, 
    NOTIFICATION_TYPE = $10
    WHERE ID = $11;`, // TODO: validate
		task.Name,
		task.Description,
		task.Resources,
		task.Notes,
		int(task.Category),
		task.Frequency,
		task.Duration,
		task.Day,
		task.Active,
		int(task.NotificationType),
		id)

	if err != nil {
		log.Fatal("Failed to update task: ", err)
	}

	return err
}

func (t TaskRepositoryImpl) DeleteTask(id int) error {
	_, err := t.db.Exec(`DELETE FROM TASKS WHERE ID = $1;`, id) // TODO: validate

	if err != nil {
		log.Fatal("Failed to delete task: ", err)
	}

	return err
}

func (t TaskRepositoryImpl) GetActiveTasks() ([]*models.TaskModel, error) {
	rows, err := t.db.Query(`SELECT * FROM TASKS WHERE IS_ACTIVE = TRUE;`)
	if err != nil {
		log.Fatal("Failed to get active tasks: ", err)
	}
	defer rows.Close()

	tasks := []*models.TaskModel{}

	for rows.Next() {
		var task models.TaskModel
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Resources, &task.Notes, &task.Category, &task.Frequency, &task.Duration, &task.Day, &task.Active, &task.NotificationType)
		if err != nil {
			log.Fatal("Failed to scan task: ", err)
		}
		tasks = append(tasks, &task)
	}

	return tasks, err
}

func (t TaskRepositoryImpl) GetTask(id int) (*models.TaskModel, error) {
	var task models.TaskModel
	err := t.db.QueryRow(`SELECT * FROM TASKS WHERE ID = $1;`, id).Scan(&task.ID, &task.Name, &task.Description, &task.Resources, &task.Notes, &task.Category, &task.Frequency, &task.Duration, &task.Day, &task.Active, &task.NotificationType)
	if err != nil {
		log.Fatal("Failed to get task: ", err)
	}

	return &task, err
}

func (t TaskRepositoryImpl) Close() error {
	return t.db.Close()
}
