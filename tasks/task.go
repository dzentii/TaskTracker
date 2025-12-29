package tasks

import (
	"errors"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type TaskStore struct {
	Tasks  map[int]*Task `json:"tasks"`
	NextID int           `json:"-"`
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		Tasks:  make(map[int]*Task),
		NextID: 1,
	}
}

func (store *TaskStore) NewTask(description string) (*TaskStore, error) {
	if len(description) < 3 {
		return store, errors.New("длина описания задачи не может быть менее 3 символов!")
	}
	task := Task{
		ID:          store.NextID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now().String()[:19],
		UpdatedAt:   time.Now().String()[:19],
	}
	store.Tasks[task.ID] = &task
	store.NextID++
	return store, nil
}

func (store *TaskStore) UpdateTask(taskID int, newDescription string) (*TaskStore, error) {
	if len(newDescription) < 3 {
		return store, errors.New("длина описания задачи не может быть менее 3 символов!")
	}
	task, ok := store.Tasks[taskID]
	if !ok {
		return store, errors.New("задачи с таким id нет!")
	}
	task.Description = newDescription
	return store, nil
}
