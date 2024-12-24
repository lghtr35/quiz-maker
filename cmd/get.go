/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lghtr35/quiz-maker/models"
	"github.com/lghtr35/quiz-maker/util"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [COMMAND] [ARGUMENTS]",
	Short: "Get entities or rankings",
}

var getQuizCmd = &cobra.Command{
	Use:   "quiz [Id]",
	Short: "Get quiz by id",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := strconv.ParseUint(args[0], 10, 32); err != nil {
			return err
		}

		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/quizzes/%s", args[0]))
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return resp.Request.Context().Err()
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.Quiz{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Quiz: %+v", unmarshalled)
		return nil
	},
}

var getQuestionCmd = &cobra.Command{
	Use:   "question [Id]",
	Short: "Get question by id",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := strconv.ParseUint(args[0], 10, 32); err != nil {
			return err
		}

		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/quizzes/questions/%s", args[0]))
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return resp.Request.Context().Err()
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.Question{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Question: %+v", unmarshalled)
		return nil
	},
}

var getScore = &cobra.Command{
	Use:   "score [UserId] [QuizId]",
	Short: "Get score by UserId and QuizId",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := strconv.ParseUint(args[0], 10, 32); err != nil {
			return err
		}
		if _, err := strconv.ParseUint(args[1], 10, 32); err != nil {
			return err
		}

		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s/quiz/%s", args[0], args[1]))
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return resp.Request.Context().Err()
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.Score{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Score: %+v", unmarshalled)
		return nil
	},
}

var getRanking = &cobra.Command{
	Use:   "ranking [UserId] [QuizId]",
	Short: "Get ranking by UserId and QuizId",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := strconv.ParseUint(args[0], 10, 32); err != nil {
			return err
		}
		if _, err := strconv.ParseUint(args[1], 10, 32); err != nil {
			return err
		}

		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s/quiz/%s/ranking", args[0], args[1]))
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return resp.Request.Context().Err()
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.GetUserRankingByScoreResponse{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Ranking: %+v", unmarshalled)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getQuizCmd)
	getCmd.AddCommand(getQuestionCmd)
	getCmd.AddCommand(getScore)
	getCmd.AddCommand(getRanking)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
