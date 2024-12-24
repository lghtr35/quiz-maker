/*
Copyright © 2024 Serdil Cagin Cakmak serdilcakmak@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-maker [COMMAND]",
	Short: "Quiz-Maker is a simple Quiz api that serves quizzes",
	Long:  `Quiz-Maker is a simple Quiz api that serves quizzes with multiple choice answered questions and calculates score and ranks users.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
