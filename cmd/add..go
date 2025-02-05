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
		addTask(task)
	}, // on running ye chalega
}

func addTask(task string) {
	// ye task file lega from the config file
	taskfilepath := config.GetTaskFile()

	var tasks []map[string]interface{}
	file, err := os.OpenFile(taskfilepath, os.O_RDWR|os.O_CREATE, 0666) // ye to khol dega 😼 file bhai
	if err != nil{
		fmt.Print("nhi khul rha bro....file meri jaan, aap kya smjhe")
		return // return kar do bhai
	}
	defer file.Close()

	// ab decode to karna padega na jo user input kiya hai bolo hai na
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error()!= "EOF" {
		fmt.Println("shittt bhai error aa rhi hai file read karne me....return ho jayo")
		return
	}
	// add karo task
	tasks = append(tasks, map[string]interface{}{
		"id" : 1,
		"tsk" : task,
		"completed" : false,
	})
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("error aa gya jayo arpit ka gand maro kuch to hag diya hai config me")
		return
	}
	// ho gaya ho gaya 😎😎😎😎
	fmt.Println("added bhai jayo yaad rakhna add kiye the")
}

