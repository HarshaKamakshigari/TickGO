package cmd

import (
	"fmt"
	"strings"

	"github.com/HarshaKamakshigari/TickGO/pkg"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new todo task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := pkg.SaveTodo(task)
		if err != nil {
			fmt.Println("❌ Failed to save todo:", err)
			return
		}
		fmt.Println("✅ Added todo:", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
