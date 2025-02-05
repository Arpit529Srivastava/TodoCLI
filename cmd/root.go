/*
Copyright © 2025 Arpit Srivastava
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todocli",
	Short: "hey ! everyone i am Arpit Srivastava here TODOCLI for adding your todolist in the cli",
	Long: `TODOCLI is a tool for managing your todo list 😅 which i know you will forget lol
	i know your are to lazy to make a website and then do it but here it is 

for example run:
	todocli add "Finally finished my work 😎" 
		or
	todocli -a "Finally made the project"
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todocli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
