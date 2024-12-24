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
	"strings"

	"github.com/lghtr35/quiz-maker/models"
	"github.com/lghtr35/quiz-maker/util"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [TYPE] [ARGUMENTS]",
	Short: "Create can be used to create entities in system. For each type it has given the arguments needed changes",
}

var createUserCmd = &cobra.Command{
	Use:   "user [name]",
	Short: "Create User with name",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("create called")
		req := models.CreateUserRequest{
			Name: args[0],
		}
		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		r := bytes.NewReader(b)
		resp, err := http.Post("http://localhost:8080/users", "application/json", r)
		if err != nil {
			return err
		}
		if resp.StatusCode != 201 {
			return resp.Request.Context().Err()
		}
		user, err := util.ReadBodyAndUnmarshal(models.User{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("UserID: %d", user.ID)
		return nil
	},
}

var createQuizCmd = &cobra.Command{
	Use:   "quiz [Name] [Questions as string seperated by comma] [Options as 2d json array]",
	Short: "Create a quiz with given parameters. Options argument is not mandatory.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		qStr := args[1]
		questions := strings.Split(qStr, ",")
		var optionRequests [][]models.CreateOptionRequest

		if len(args) > 2 {
			oStr := args[2]
			err := json.Unmarshal([]byte(oStr), &optionRequests)
			if err != nil {
				return err
			}
		}

		questionsRequests := make([]models.CreateQuestionRequest, len(questions))
		for i := 0; i < len(questionsRequests); i++ {
			questionsRequests[i].Question = questions[i]
			if len(args) > 2 {
				questionsRequests[i].Options = &optionRequests[i]
			}
		}

		req := models.CreateQuizRequest{
			Name:      name,
			Questions: questionsRequests,
		}
		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		r := bytes.NewReader(b)
		resp, err := http.Post("http://localhost:8080/quizzes", "application/json", r)
		if err != nil {
			return err
		}
		if resp.StatusCode != 201 {
			return resp.Request.Context().Err()
		}
		quiz, err := util.ReadBodyAndUnmarshal(models.Quiz{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("QuizID: %d", quiz.ID)
		return nil
	},
}

var createOptionCmd = &cobra.Command{
	Use:   "option [QuestionId] [Value] [IsCorrect]",
	Short: "Create a question option with given parameters.",
	Args:  cobra.MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		qId := args[0]
		value := args[1]
		iC := args[2]

		isCorrect, err := strconv.ParseBool(iC)
		if err != nil {
			return err
		}

		req := models.CreateOptionRequest{
			Value:     value,
			IsCorrect: isCorrect,
		}
		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		r := bytes.NewReader(b)
		resp, err := http.Post(fmt.Sprintf("http://localhost:8080/quizzes/questions/%s/options", qId), "application/json", r)
		if err != nil {
			return err
		}
		if resp.StatusCode != 201 {
			return resp.Request.Context().Err()
		}
		option, err := util.ReadBodyAndUnmarshal(models.OptionBase{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("OptionID: %d", option.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createUserCmd)
	createCmd.AddCommand(createQuizCmd)
	createCmd.AddCommand(createOptionCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
