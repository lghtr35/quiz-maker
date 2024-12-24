# quiz-maker

Quiz maker is a simple Cobra-based quiz API that is in-memory. It can save quizzes, users, users' quiz progressions, users' quiz scores, quiz questions, question options and users' question answers.

To use build and run `quiz-maker --help`

## App flow

Flow to take a quiz and see score and rankings:

1. Create quiz, questions and options
2. Register an user
3. Begin a quiz with user
4. Answer current question in progression of quiz
5. Submit quiz (before answering all questions is possible too)
6. Get score
7. Get rankings
