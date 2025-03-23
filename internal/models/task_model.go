package models

import "time"

type TaskModel struct {
	ID          int
	Name        string
	Description string
	Resources   string // json string
	Notes       string

	Category  Category // enum etc.
	Frequency int      // once every x days
	Duration  int      // in minutes
	Day       int      // task can only be done on this day of the week, XORed

	Active           bool
	NotificationType NotificationType // enum
	LastDone         int64            // timestamp
	LastIgnored      int64            // timestamp; Ignore meaning mark as done without doing
	CreatedAt        int64            // timestamp
	UpdatedAt        int64            // timestamp
}

type Category int

const (
	CategoryDrawing = iota + 1 // Sketching, Studies, Full rendering
	CategoryWriteup            // General writeup -- state of product, tech lessons
	CategoryDev                // FE website, mobile app, home server, arduino
	CategoryDnd                // Movement
	CategoryChore              // Mowing, Cleaning etc.
	CategoryFin                // Dollar cost averaging, budget report,
	CategoryAdhoc              // One off randoms
)

const (
	_         = iota
	DayMonday = iota * 2
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)

type NotificationType int

const (
	NotificationTypeNone = iota + 1
	NotificationTypeNext // e.g. Task every sunday, if not done, remind every day when task can be done.
)

func NewTask(
	name,
	description string,
	resources string,
	notes string,
	category Category,
	frequency int,
	duration int,
	day int,
	notificationType NotificationType,
) *TaskModel {
	return &TaskModel{
		Name:             name,
		Description:      description,
		Resources:        resources,
		Category:         category,
		Frequency:        frequency,
		Duration:         duration,
		Day:              day,
		Active:           true,
		NotificationType: notificationType,

		LastDone:    time.Now().Unix(),
		LastIgnored: time.Now().Unix(),
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}

type TaskStatistics struct {
	ID int

	// Last done stats
	LastDone    int64
	LastIgnored int64
	CreatedAt   int64
	UpdatedAt   int64

	// Stats
	TotalDone    int
	TotalIgnored int
}
