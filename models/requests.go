package models

type PaginationRequest struct {
	Page uint `json:"page"`
	Size uint `json:"size"`
}

type ReadUsersRequest struct {
	PaginationRequest
	IDList *[]uint `json:"idList"`
	Name   *string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	ID   uint    `json:"id" binding:"required"`
	Name *string `json:"name"`
}

type ReadQuizRequest struct {
	PaginationRequest
	IDList *[]uint `json:"idList"`
	Name   *string `json:"name"`
}

type CreateQuizRequest struct {
	Name      string                  `json:"name" binding:"required"`
	Questions []CreateQuestionRequest `json:"questions" bindind:"required"`
}
type CreateQuestionRequest struct {
	Question string                `json:"question" binding:"required"`
	Answers  []CreateAnswerRequest `json:"answers" binding:"required"`
}
type CreateAnswerRequest struct {
	Value     string `json:"value"`
	IsCorrect bool   `json:"isCorrect"`
}

type UpdateQuizRequest struct {
	ID   uint    `json:"id" binding:"required"`
	Name *string `json:"name"`
}

type BeginQuizRequest struct {
	QuizID uint `json:"quizId" binding:"required"`
	UserID uint `json:"userId" binding:"required"`
}

type AnswerQuizQuestionRequest struct {
	OptionID      uint `json:"optionId" binding:"required"`
	ProgressionID uint `json:"progressionId" binding:"required"`
}

type FinalizeQuizRequest struct {
	ProgressionID uint `json:"progressionId" binding:"required"`
}
