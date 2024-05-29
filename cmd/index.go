package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/igweo/jeeves/internal/utils"
	"github.com/spf13/cobra"
)

var (
	modelFlag string
	d         = color.New(color.FgCyan, color.Bold)
	rootCmd   = &cobra.Command{
		Use:   "jeeves",
		Short: "Jeeves gives you the ability to ask questions using GenAI right in your terminal",
	}
	askCmd = &cobra.Command{
		Use:   "ask",
		Short: "Use GPT3.5T or GPT4 to answer a question.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			modelFlagValue, _ := cmd.Flags().GetString("model")
			question := args[0]
			d.Printf("%s\n", utils.ChatCompletion(question, modelFlagValue))
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(askCmd)
	askCmd.Flags().StringVar(&modelFlag, "model", "", "Choose to use 'gpt4o' for the model, gpt3 is the default")
}
