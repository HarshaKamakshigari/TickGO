package pkg

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func getDBPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".tickgo.json")
}

func SaveTodo(task string) error {
	todos, _ := LoadTodos()
	todos = append(todos, Todo{Task: task, Done: false})
	data, _ := json.MarshalIndent(todos, "", "  ")
	return os.WriteFile(getDBPath(), data, 0644)
}

func LoadTodos() ([]Todo, error) {
	var todos []Todo
	filePath := getDBPath()
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &todos)
	return todos, err
}
