package cmd

import (
	"fmt"
	"os"

	"github.com/igweo/jeeves/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jeeves",
	Short: "Jeeves gives you the ability to ask questions using GenAI right in your terminal",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "ask",
			Short: "Use GPT3.5Turbo to answer a question.",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				question := args[0]
				fmt.Println(utils.ChatCompletion(question))
			},
		},
	)
}
