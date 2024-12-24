package main

import (
	"github.com/lghtr35/quiz-maker/cmd"
	_ "github.com/lghtr35/quiz-maker/docs"
)

// @title           Quiz Maker API
// @version         0.0.1
// @description     Quiz Maker API for Fast Track Interview
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
func main() {
	cmd.Execute()
}
