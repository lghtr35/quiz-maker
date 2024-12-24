/*
Copyright © 2024 Serdil Cagin Cakmak serdilcakmak@gmail.com
*/
package cmd

import (
	"log"
	"net/http"

	_ "github.com/lghtr35/quiz-maker/docs"
	"github.com/lghtr35/quiz-maker/handlers"
	"github.com/lghtr35/quiz-maker/models"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts to listen for connections",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		err = db.AutoMigrate(
			&models.User{},
			&models.Quiz{},
			&models.Question{},
			&models.Progression{},
			&models.Option{},
			&models.Score{},
			&models.Answer{},
		)
		if err != nil {
			panic(err)
		}

		handlers := handlers.InitializeHandlers(db)

		mux := http.NewServeMux()

		mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		))
		for _, h := range handlers {
			mux = h.ConfigureSelf(mux)
		}

		log.Println("Started listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", mux))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
