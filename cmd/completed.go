package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"todocli/config"

	"github.com/spf13/cobra"
)
var completeCmd = &cobra.Command{
	Use:   "complete [task_id]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1), // Accept exactly one argument (the task ID)
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		completeTask(taskID)
	},
}

func completeTask(taskID string) {
	// Get task file path from config
	taskFilePath := config.GetTaskFile()

	// Read existing tasks
	var tasks []map[string]interface{}
	file, err := os.OpenFile(taskFilePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the tasks
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error() != "EOF" {
		fmt.Println("Error reading tasks:", err)
		return
	}

	// Find and mark task as completed
	for _, task := range tasks {
		if fmt.Sprintf("%v", task["id"]) == taskID {
			task["done"] = true
			fmt.Printf("Task '%s' marked as completed.\n", task["tsk"])
			break
		}
	}

	// Write updated tasks back to file
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Error saving task:", err)
		return
	}
}
