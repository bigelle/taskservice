package database

import (
	"sync"
	"time"
)

type Task struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at,omitzero"`
	UpdatedAt       time.Time `json:"updated_at,omitzero"`
	CreatorName     string    `json:"creator_name"`
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
}

func NewDB() TaskDB {
	return NewLocalDB()
}

type TaskDB interface {
	Create(creator, name, desc string) (uint, error)
	View(taskID uint) (Task, error)
	Update(taskID uint, task Task) (Task, error)
	Delete(taskID uint) error
}

type LocalDB struct {
	Tasks []Task
}

func (d *LocalDB) Create(creator, name, desc string) (uint, error) {
	// TODO:
	return 0, nil
}

func (d *LocalDB) View(taskID uint) (Task, error) {
	// TODO:
	return Task{}, nil
}

func (d *LocalDB) Update(taskID uint, task Task) (Task, error) {
	// TODO:
	return Task{}, nil
}

func (d *LocalDB) Delete(taskID uint) error {
	// TODO:
	return nil
}

var (
	db   LocalDB
	once sync.Once
)

func NewLocalDB() *LocalDB {
	once.Do(func() {
		db = LocalDB{}
	})
	return &db
}
