package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		if _, err := db.CreateTask(task); err != nil {
			fmt.Println("Something went wront:", err)
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
