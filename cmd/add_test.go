package cmd

import (
	"encoding/json"
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	// Create a temporary file to simulate the task file
	tempFile, err := os.CreateTemp("", "tasks.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Cleanup after test

	// Call modified addTask function with test file
	addTask("Test Task", tempFile.Name())

	// Read back the file content
	tempFile.Seek(0, 0)
	var tasks []map[string]interface{}
	decoder := json.NewDecoder(tempFile)
	if err := decoder.Decode(&tasks); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Verify the task was added
	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0]["tsk"] != "Test Task" {
		t.Errorf("Expected task 'Test Task', got '%s'", tasks[0]["tsk"])
	}
	if tasks[0]["completed"] != false {
		t.Errorf("Expected completed to be false, got %v", tasks[0]["completed"])
	}
}
