package models

import "time"

type Base struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Score struct {
	Base
	QuizID uint    `json:"quizId"`
	UserID uint    `json:"userId"`
	Score  float32 `json:"score"`
}

type Progression struct {
	Base
	UserID            uint `json:"userId"`
	QuizID            uint `json:"quizId"`
	IsFinished        bool `json:"isFinished"`
	CurrentQuestionID uint `json:"currentQuestionId"`
	QuestionNumber    int  `json:"questionNumber"`
}

type Answer struct {
	Base
	UserID   uint `json:"userId"`
	OptionID uint `json:"optionId"`
	QuizID   uint `json:"quizId"`
}

type OptionBase struct {
	Base
	QuestionID uint     `json:"questionId"`
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
	QuizID       uint          `json:"quizId"`
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
