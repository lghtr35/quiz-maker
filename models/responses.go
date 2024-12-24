package models

type PaginationResponse struct {
	Page    uint32 `json:"page"`
	Size    uint32 `json:"size"`
	Content []any  `json:"content"`
}

type QuestionWithOptionsResponse struct {
	Base
	Question string       `json:"question"`
	Answers  []OptionBase `json:"options"`
	QuizID   uint32       `json:"quizId"`
}

type BeginQuizResponse struct {
	Progression Progression `json:"progression"`
}

type AnswerQuizQuestionResponse struct {
	Progression Progression `json:"progression"`
}

type FinalizeQuizResponse struct {
	Score Score `json:"score"`
}

type GetUserRankingByScoreResponse struct {
	Rank    uint32  `json:"rank"`
	Percent float32 `json:"percent"`
	Message string  `json:"message"`
	Score   Score   `json:"userScore"`
}
