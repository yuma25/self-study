package taskmanager

import (
	"errors"
	"fmt"
)

// Task 構造体
type Task struct {
	ID   int
	Title string
	Done  bool
}

// Manager 構造体 (タスクの集合を管理)
type Manager struct {
	tasks  []Task
	nextID int
}

func NewManager() *Manager {
	return &Manager{
		tasks:  []Task{},
		nextID: 1,
	}
}

// Add: タスクの追加 (ポインタレシーバ)
func (m *Manager) Add(title string) {
	task := Task{
		ID:    m.nextID,
		Title: title,
		Done:  false,
	}
	m.tasks = append(m.tasks, task)
	m.nextID++
}

// Complete: タスクの完了
func (m *Manager) Complete(id int) error {
	for i := range m.tasks {
		if m.tasks[i].ID == id {
			m.tasks[i].Done = true
			return nil
		}
	}
	return errors.New("task not found")
}

// List: タスクの表示 (Stringerインターフェースに近い役割)
func (m *Manager) List() {
	if len(m.tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, t := range m.tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d: %s %s\n", t.ID, status, t.Title)
	}
}
