package database

import (
	"fmt"
	"sync"
	"time"
)

type LocalDB struct {
	Tasks []Task
}

func (d *LocalDB) Create(name, desc string) (uint, error) {
	if name == "" {
		return 0, fmt.Errorf("task name must not be empty")
	}
	if desc == "" {
		return 0, fmt.Errorf("description must not be empty")
	}
	return d.create(name, desc)
}

func (d *LocalDB) create(name, desc string) (id uint, err error) {
	var newID uint = 1
	if len(d.Tasks) != 0 {
		last := d.Tasks[len(d.Tasks)-1]

		newID = last.ID + 1
	}

	task := Task{
		ID:          newID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Name:        name,
		Description: desc,
		Status:      StatusPending,
	}

	d.Tasks = append(d.Tasks, task)

	return newID, nil
}

func (d *LocalDB) View(taskID uint) (Task, error) {
	if taskID == 0 {
		return Task{}, fmt.Errorf("task ID must not be empty")
	}

	for _, task := range d.Tasks {
		if task.ID == taskID {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("no tasks with id %d", taskID)
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
