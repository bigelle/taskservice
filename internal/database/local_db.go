package database

import (
	"fmt"
	"slices"
	"sync"
	"time"
)

type LocalDB struct {
	Tasks []Task
	mu    sync.Mutex
}

func (d *LocalDB) Create(name, desc string) (uint, error) {
	if name == "" {
		return 0, fmt.Errorf("%w: task name must not be empty", ErrInvalidData)
	}
	if desc == "" {
		return 0, fmt.Errorf("%w: description must not be empty", ErrInvalidData)
	}
	return d.create(name, desc), nil
}

func (d *LocalDB) create(name, desc string) (ID uint) {
	d.mu.Lock()
	defer d.mu.Unlock()

	ID = 1
	if len(d.Tasks) != 0 {
		last := d.Tasks[len(d.Tasks)-1]

		ID = last.ID + 1
	}

	task := Task{
		ID:          ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Name:        name,
		Description: desc,
		Status:      StatusPending,
	}

	d.Tasks = append(d.Tasks, task)

	return ID
}

func (d *LocalDB) View(taskID uint) (Task, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if taskID == 0 {
		return Task{}, fmt.Errorf("%w: task ID must not be empty", ErrInvalidData)
	}

	for _, task := range d.Tasks {
		if task.ID == taskID {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("%w: no tasks with ID %d", ErrNoRecord, taskID)
}

func (d *LocalDB) UpdateStatus(taskID uint, status string) (Task, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if taskID == 0 {
		return Task{}, fmt.Errorf("%w: task ID must not be empty", ErrInvalidData)
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

	return Task{}, fmt.Errorf("%w: no tasks with ID %d", ErrNoRecord, taskID)
}

func (d *LocalDB) UpdateResult(taskID uint, result string) (Task, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if taskID == 0 {
		return Task{}, fmt.Errorf("%w: task ID must not be empty", ErrInvalidData)
	}
	if result == "" {
		return Task{}, fmt.Errorf("%w: task ID must not be empty", ErrInvalidData)
	}

	for i := range d.Tasks {
		if d.Tasks[i].ID == taskID {
			d.Tasks[i].Result = &result
			d.Tasks[i].UpdatedAt = time.Now()
			return d.Tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("%w: no tasks with ID %d", ErrNoRecord, taskID)
}

func (d *LocalDB) Delete(taskID uint) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if taskID == 0 {
		return fmt.Errorf("%w: task ID must not be empty", ErrInvalidData)
	}

	upd := slices.DeleteFunc(d.Tasks, func(t Task) bool { return t.ID == taskID })
	if len(upd) == len(d.Tasks) {
		return fmt.Errorf("%w: no tasks with ID %d", ErrNoRecord, taskID)
	}
	d.Tasks = upd

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
