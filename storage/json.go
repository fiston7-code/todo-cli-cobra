package storage

import (
	"github.com/fiston7-code/todo-cli-cobra/models"

	"encoding/json"
	"os"
	"path/filepath"
)

const filePath = "data/tasks.json"

func Save(tasks []models.Task) error {
	// Créer le dossier data s'il n'existe pas
	dir := filepath.Dir(filePath)
	os.MkdirAll(dir, 0755)

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func Load() ([]models.Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
