package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"todocli/config"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "You Lazy Shit 😮‍💨, Added to your to-do list",
	Long: `kya he likhe add kar diye thumara list
	

	KOI MATLAB HE NHI HAI AUR LIKHNE KA HEHE
	`,
	Args: cobra.ExactArgs(1), // sirf ek he arguments leta hai (jo bhi task description hai)
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		taskFile := config.GetTaskFile() // Fetch the task file path
		addTask(task, taskFile)
	}, // on running ye chalega
}

   func addTask(task, taskFile string) {
	var tasks []map[string]interface{}

	// Open the task file
	file, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("nhi khul rha bro....file meri jaan, aap kya smjhe")
		return
	}
	defer file.Close()

	// Decode existing tasks
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error() != "EOF" {
		fmt.Println("shittt bhai error aa rhi hai file read karne me....return ho jayo")
		return
	}

	// Add new task
	tasks = append(tasks, map[string]interface{}{
		"id":        len(tasks) + 1,
		"tsk":       task,
		"completed": false,
	})

	// Write back to the file
	file.Seek(0, 0)
	file.Truncate(0) // Clear the file before writing new data
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("error aa gya jayo arpit ka gand maro kuch to hag diya hai config me")
		return
	}

	fmt.Println("added bhai jayo yaad rakhna add kiye the")
}
