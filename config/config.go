package config

import (
	"fmt"
	"os"
)

func GetTaskFile() string{
	taskFilePath := "config/todo.json"
	if _, err := os.Stat(taskFilePath); os.IsNotExist(err){
		// ab file nhi hai to create karo bancho
		file, err := os.Create(taskFilePath)
		if err != nil{
			fmt.Println("error while creating fir arpit ki gand maro 🔫", err)
		}
		os.Exit(1)
		defer file.Close()
	}
	return taskFilePath
}