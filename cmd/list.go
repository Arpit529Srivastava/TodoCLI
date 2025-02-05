package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"todocli/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "You Lazy Shit 😮‍💨, Here is you f**(King) list 😓 ",
	Long: ` Dekho itna karna hai kab karoge kab bhai 
		
		JAO KARO BSDK KARO NA 😠💢
FOR EXAMPLE :
	todocli list 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		listall()
	}, // on running ye chalega
}

func listall(){
	taskfilepath := config.GetTaskFile()
	var tasks []map[string]interface{}
	file, err := os.OpenFile(taskfilepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("error aa rha hai bhai", err)
		return
	}
	defer file.Close()

	// now decode the task bhai
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error() != "EOF" {
		fmt.Println("error aa rha hai maa chodo developer ka nhi to simple ye issue daal to 😃", err)
		return
	}

	// now find all the compeleted tasks bhaijaan
	if len(tasks) == 0{
		fmt.Println("there are no task!!! how can this be possible? 😱😱😱")
		return
	}

	/// ye rha aapka list meri jaaannn
	for _, task := range tasks{
		done := "Not Done"
		if task["done"].(bool){
			done = "Completed"
		}
		fmt.Printf("ID : %d | Task : %s | Status : %s \n", int(task["id"].(float64)), task["tsk"], done)
	}
}
