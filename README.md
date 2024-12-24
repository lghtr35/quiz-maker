# quiz-maker

Quiz maker is a simple Cobra-based quiz API that is in-memory. It can save quizzes, users, users' quiz progressions, users' quiz scores, quiz questions, question options and users' question answers.

To use build and run `quiz-maker --help`

## Build and Run

Needs Go 1.23 to build and run the application.

### Build

Run in source directory where main file resides.
`go build`

### Run

App should respond to `quiz-maker`. If not try `./quiz-maker` in executable directory.

## App flow

Flow to take a quiz and see score and rankings:

1. Create quiz, questions and options
2. Register an user
3. Begin a quiz with user
4. Answer current question in progression of quiz
5. Submit quiz (before answering all questions is possible too)
6. Get score
7. Get rankings
