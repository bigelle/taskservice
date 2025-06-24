package database

import (
	"fmt"
	"sync"
	"time"
)

type LocalDB struct {
	Tasks []Task
}

func (d *LocalDB) Create(creator, name, desc string) (uint, error) {
	if creator == "" {
		return 0, fmt.Errorf("creator name must not be empty")
	}
	if name == "" {
		return 0, fmt.Errorf("task name must not be empty")
	}
	if desc == "" {
		return 0, fmt.Errorf("description must not be empty")
	}
	return d.create(creator, name, desc)
}

func (d *LocalDB) create(creator, name, desc string) (id uint, err error) {
	var newID uint = 1
	if len(d.Tasks) != 0 {
		last := d.Tasks[len(d.Tasks)-1]

		newID = last.ID + 1
	} 

	task := Task{
		ID:              newID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CreatorName:     creator,
		TaskName:        name,
		TaskDescription: desc,
	}

	d.Tasks = append(d.Tasks, task)

	return newID, nil
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
