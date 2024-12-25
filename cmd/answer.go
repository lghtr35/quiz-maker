/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/lghtr35/quiz-maker/models"
	"github.com/lghtr35/quiz-maker/util"
	"github.com/spf13/cobra"
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer [ProgressionId] [OptionId]",
	Short: "Answer a question to progress in quiz",
	Long:  `Answer a question to progress in quiz. It takes an optionId and progressionId to save an answer to the current question in quiz.`,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("answer called")

		progressionId, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return err
		}
		optionId, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			return err
		}

		req := models.AnswerQuizQuestionRequest{
			OptionID:      uint32(optionId),
			ProgressionID: uint32(progressionId),
		}

		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		r := bytes.NewReader(b)
		resp, err := http.Post("http://localhost:8080/quizzes/answer", "application/json", r)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			s, err := util.ReadBodyAndGetString(resp.Body)
			if err != nil {
				return err
			}
			log.Printf("Status: %d, Error: %s", resp.StatusCode, s)
			return nil
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.AnswerQuizQuestionResponse{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Progression: %+v", unmarshalled.Progression)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(answerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
