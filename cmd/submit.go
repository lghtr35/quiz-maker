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

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit [ProgressionId]",
	Short: "Finalize a quiz progression qith a calculation of a score.",
	Long:  `Finalize a quiz progression qith a calculation of a score. Submit takes progressionId argument to finalize a quiz and returns score.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("submit called")

		progressionId, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return err
		}

		req := models.FinalizeQuizRequest{
			ProgressionID: uint32(progressionId),
		}

		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		r := bytes.NewReader(b)
		resp, err := http.Post("http://localhost:8080/quizzes/submit", "application/json", r)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return resp.Request.Context().Err()
		}
		unmarshalled, err := util.ReadBodyAndUnmarshal(models.FinalizeQuizResponse{}, resp.Body)
		if err != nil {
			return err
		}
		log.Printf("Score: %+v", unmarshalled.Score)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
