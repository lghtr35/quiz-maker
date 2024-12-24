package models

import "time"

type Base struct {
	ID        uint32    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Score struct {
	Base
	QuizID uint32  `json:"quizId"`
	UserID uint32  `json:"userId"`
	Score  float32 `json:"score"`
}

type Progression struct {
	Base
	UserID            uint32 `json:"userId"`
	QuizID            uint32 `json:"quizId"`
	IsFinished        bool   `json:"isFinished"`
	CurrentQuestionID uint32 `json:"currentQuestionId"`
	QuestionNumber    int    `json:"questionNumber"`
}

type Answer struct {
	Base
	UserID   uint32 `json:"userId"`
	OptionID uint32 `json:"optionId"`
	QuizID   uint32 `json:"quizId"`
}

type OptionBase struct {
	Base
	QuestionID uint32   `json:"questionId"`
	Value      string   `json:"value"`
	Answers    []Answer `json:"answers"`
}

type Option struct {
	OptionBase
	IsCorrect bool `json:"isCorrect"`
}

type Question struct {
	Base
	Question     string        `json:"question"`
	QuizID       uint32        `json:"quizId"`
	Options      []Option      `json:"options"`
	Progressions []Progression `gorm:"foreignKey:CurrentQuestionID" json:"progressions"`
}

type Quiz struct {
	Base
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
	Answers   []Answer   `json:"answers"`
}

type User struct {
	Base
	Name    string   `json:"name"`
	Answers []Answer `json:"answer"`
}
