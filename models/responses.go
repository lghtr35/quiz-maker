package models

type PaginationResponse struct {
	Page    uint  `json:"page"`
	Size    uint  `json:"size"`
	Content []any `json:"content"`
}

type QuestionWithOptionsResponse struct {
	Base
	Question string       `json:"question"`
	Answers  []OptionBase `json:"options"`
	QuizID   uint         `json:"quizId"`
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
	Rank    uint    `json:"rank"`
	Percent float32 `json:"percent"`
	Message string  `json:"message"`
	Score   *Score  `json:"userScore"`
}
