package model

import "time"

// TaskStatus represents the kanban column a task is in.
type TaskStatus string

const (
	StatusBacklog    TaskStatus = "Backlog"
	StatusTodo       TaskStatus = "To Do"
	StatusInProgress TaskStatus = "In Progress"
	StatusTest       TaskStatus = "Test"
	StatusDone       TaskStatus = "Done"
)

type Priority string

const (
	PriorityLow    Priority = "Low"
	PriorityMedium Priority = "Medium"
	PriorityHigh   Priority = "High"
)

type WorkType string

const (
	TypeEpic  WorkType = "Epic"
	TypeStory WorkType = "Story"
	TypeTask  WorkType = "Task"
)

type Dependency struct {
	ID       string `json:"id"`
	Type     WorkType `json:"type"` // Epic, Story, Task
	Blocking bool   `json:"blocking"` // If true, this must be completed first
}

type WorkItem struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Type        WorkType      `json:"type"`
	ParentID    string        `json:"parent_id,omitempty"` // For child linking
	Dependencies []Dependency `json:"dependencies,omitempty"`

	Tags       []string    `json:"tags,omitempty"`
	Priority   Priority    `json:"priority,omitempty"`
	Points     int         `json:"points,omitempty"`
	Status     TaskStatus  `json:"status"`
	SprintID   string      `json:"sprint_id,omitempty"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

