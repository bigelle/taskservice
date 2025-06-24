package database

import (
	"fmt"
	"slices"
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

	return Task{}, fmt.Errorf("no tasks with ID %d", taskID)
}

func (d *LocalDB) UpdateStatus(taskID uint, status string) (Task, error) {
	if taskID == 0 {
		return Task{}, fmt.Errorf("task ID must not be empty")
	}
	if TaskStatusFromString(status) == StatusUndefined {
		return Task{}, fmt.Errorf("undefined task status: %s", status)
	}

	for i := range d.Tasks {
		if d.Tasks[i].ID == taskID {
			if TaskStatusFromString(status) == d.Tasks[i].Status {
				return d.Tasks[i], nil
			}

			d.Tasks[i].Status = TaskStatusFromString(status)
			d.Tasks[i].UpdatedAt = time.Now()
			return d.Tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("no tasks with ID %d", taskID)
}

func (d *LocalDB) UpdateResult(taskID uint, result string) (Task, error) {
	if taskID == 0 {
		return Task{}, fmt.Errorf("task ID must not be empty")
	}
	if result == "" {
		return Task{}, fmt.Errorf("result must not be empty")
	}

	for i := range d.Tasks {
		if d.Tasks[i].ID == taskID {
			d.Tasks[i].Result = &result
			d.Tasks[i].UpdatedAt = time.Now()
			return d.Tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("no tasks with ID %d", taskID)
}

func (d *LocalDB) Delete(taskID uint) error {
	if taskID == 0 {
		return fmt.Errorf("task ID must not be empty")
	}

	d.Tasks = slices.DeleteFunc(d.Tasks, func(t Task) bool { return t.ID == taskID})
	
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
