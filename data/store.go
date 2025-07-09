package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/winslowb/onebill-kanban/model"
	"github.com/google/uuid"
)

var (
	baseDir     = ".onebill"
	workitemDir = filepath.Join(baseDir, "workitems")
	sprintDir   = filepath.Join(baseDir, "sprints")
)

func InitStore() error {
	dirs := []string{baseDir, workitemDir, sprintDir}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", d, err)
		}
	}
	return nil
}

func SaveWorkItem(w *model.WorkItem) error {
	if w.ID == "" {
		w.ID = uuid.New().String()
		w.CreatedAt = time.Now()
	}
	w.UpdatedAt = time.Now()

	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return err
	}

	file := filepath.Join(workitemDir, fmt.Sprintf("%s.json", w.ID))
	return os.WriteFile(file, data, 0644)
}

func LoadAllWorkItems() ([]model.WorkItem, error) {
	files, err := os.ReadDir(workitemDir)
	if err != nil {
		return nil, err
	}

	var items []model.WorkItem
	for _, f := range files {
		path := filepath.Join(workitemDir, f.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var item model.WorkItem
		if err := json.Unmarshal(data, &item); err == nil {
			items = append(items, item)
		}
	}
	return items, nil
}

