package models

type PaginationRequest struct {
	Page uint32 `json:"page"`
	Size uint32 `json:"size"`
}

type ReadUsersRequest struct {
	PaginationRequest
	IDList *[]uint32 `json:"idList"`
	Name   *string   `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	ID   uint32  `json:"id" binding:"required"`
	Name *string `json:"name"`
}

type ReadQuizRequest struct {
	PaginationRequest
	IDList *[]uint32 `json:"idList"`
	Name   *string   `json:"name"`
}

type CreateQuizRequest struct {
	Name      string                  `json:"name" binding:"required"`
	Questions []CreateQuestionRequest `json:"questions" bindind:"required"`
}
type CreateQuestionRequest struct {
	Question string                 `json:"question" binding:"required"`
	Options  *[]CreateOptionRequest `json:"options" `
}
type CreateOptionRequest struct {
	Value     string `json:"value"`
	IsCorrect bool   `json:"isCorrect"`
}

type UpdateQuizRequest struct {
	ID   uint32  `json:"id" binding:"required"`
	Name *string `json:"name"`
}

type BeginQuizRequest struct {
	QuizID uint32 `json:"quizId" binding:"required"`
	UserID uint32 `json:"userId" binding:"required"`
}

type AnswerQuizQuestionRequest struct {
	OptionID      uint32 `json:"optionId" binding:"required"`
	ProgressionID uint32 `json:"progressionId" binding:"required"`
}

type FinalizeQuizRequest struct {
	ProgressionID uint32 `json:"progressionId" binding:"required"`
}
