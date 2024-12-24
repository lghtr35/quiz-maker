/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lghtr35/quiz-maker/models"
	"github.com/lghtr35/quiz-maker/util"
	"github.com/spf13/cobra"
)

// beginCmd represents the begin command
var beginCmd = &cobra.Command{
	Use:   "begin [UserId] [QuizId]",
	Short: "Begin a quiz as an user",
	Long:  `Begin a quiz using QuizId as an user with UserId. It will return a Progression object for that user and quiz specifically`,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("begin called")

		userId, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return err
		}
		quizId, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			return err
		}

		req := models.BeginQuizRequest{
			QuizID: uint32(quizId),
			UserID: uint32(userId),
		}

		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		r := bytes.NewReader(b)
		resp, err := http.Post("http://localhost:8080/quizzes/begin", "application/json", r)
		if err != nil {
			return err
		}
		if resp.StatusCode != 201 {
			return resp.Request.Context().Err()
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.BeginQuizResponse{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Progression: %+v", unmarshalled.Progression)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(beginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// beginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// beginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
