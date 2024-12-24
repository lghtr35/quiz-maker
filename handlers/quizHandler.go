package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/lghtr35/quiz-maker/models"
	"github.com/lghtr35/quiz-maker/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuizHandler struct {
	db *gorm.DB
}

func newQuizHandler(db *gorm.DB) *QuizHandler {
	return &QuizHandler{db: db}
}

func (h *QuizHandler) ConfigureSelf(m *http.ServeMux) *http.ServeMux {
	m.HandleFunc("GET /quizzes/questions/{id}", h.getQuestion)
	m.HandleFunc("GET /quizzes/{id}", h.readQuizWithID)
	m.HandleFunc("POST /quizzes", h.createQuiz)
	m.HandleFunc("PATCH /quizzes", h.updateQuiz)

	m.HandleFunc("DELETE /quizzes/{id}", h.deleteQuiz)

	m.HandleFunc("POST /quizzes/begin", h.beginQuiz)
	m.HandleFunc("POST /quizzes/answer", h.answerQuizQuestion)
	m.HandleFunc("POST /quizzes/submit", h.calculateScore)

	return m
}

// createQuiz
// @Summary Create a new quiz
// @Description Creates a new quiz along with its questions and answers.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param quiz body models.CreateQuizRequest true "Quiz details"
// @Success 201 {object} models.Quiz
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes [post]
func (h *QuizHandler) createQuiz(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => CreateQuiz invoked", r.Method, r.URL.Path)
	var request models.CreateQuizRequest

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(b, &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quiz := models.Quiz{
		Name: request.Name,
	}
	res := h.db.Create(&quiz)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, q := range request.Questions {
		question := models.Question{
			Question: q.Question,
			QuizID:   quiz.ID,
		}
		res := h.db.Create(&question)
		if res.Error != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
			return
		}
		for _, a := range q.Answers {
			answer := models.Option{
				OptionBase: models.OptionBase{
					Value:      a.Value,
					QuestionID: question.ID,
				},
				IsCorrect: a.IsCorrect,
			}
			res := h.db.Create(&answer)
			if res.Error != nil {
				http.Error(w, res.Error.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	b, err = json.Marshal(quiz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

// updateQuiz
// @Summary Update an existing quiz
// @Description Updates the details of an existing quiz.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param quiz body models.UpdateQuizRequest true "Updated quiz details"
// @Success 200 {object} models.Quiz
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes [patch]
func (h *QuizHandler) updateQuiz(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => UpdateQuiz invoked", r.Method, r.URL.Path)
	var request models.UpdateQuizRequest

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(b, &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quiz := new(models.Quiz)
	res := h.db.First(quiz, request.ID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	if request.Name != nil && *request.Name != "" {
		quiz.Name = *request.Name
	}

	res = h.db.Save(quiz)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err = json.Marshal(quiz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// readQuizWithID
// @Summary Get a quiz by ID
// @Description Retrieves a quiz by its ID, including its questions and answers.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param id path string true "Quiz ID"
// @Success 200 {object} models.Quiz
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes/{id} [get]
func (h *QuizHandler) readQuizWithID(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => ReadQuizWithID invoked", r.Method, r.URL.Path)
	id := r.PathValue("id")

	var quiz models.Quiz
	res := h.db.Preload(clause.Associations).First(&quiz, id)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(quiz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// deleteQuiz
// @Summary Delete a quiz
// @Description Deletes a quiz by its ID, including its associated questions and answers.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param id path string true "Quiz ID"
// @Success 204 "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes/{id} [delete]
func (h *QuizHandler) deleteQuiz(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => DeleteQuiz invoked", r.Method, r.URL.Path)
	id := r.PathValue("id")

	res := h.db.Select(clause.Associations).Delete(&models.Quiz{}, id)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
}

// beginQuiz
// @Summary Begin a quiz
// @Description Starts a quiz session for a user, initializing the progression with the first question.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param beginQuiz body models.BeginQuizRequest true "Quiz start details"
// @Success 201 {object} models.BeginQuizResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes/begin [post]
func (h *QuizHandler) beginQuiz(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => BeginQuiz invoked", r.Method, r.URL.Path)
	request, err := util.ReadBodyAndUnmarshal(models.BeginQuizRequest{}, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get quiz and check if it is okay to start progressing on it
	var quiz models.Quiz
	res := h.db.Preload(clause.Associations).First(&quiz, request.QuizID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}
	if len(quiz.Questions) < 1 {
		http.Error(w, "quizHandler: quiz does not have any questions,", http.StatusInternalServerError)
		return
	}

	// Create a new progression for user to keep track of where we are at
	progression := models.Progression{
		UserID:            request.UserID,
		QuizID:            request.QuizID,
		IsFinished:        false,
		CurrentQuestionID: quiz.Questions[0].ID,
		QuestionNumber:    0,
	}
	res = h.db.Create(&progression)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := models.BeginQuizResponse{
		Progression: progression,
	}
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

// answerQuizQuestion
// @Summary Answer a quiz question
// @Description Submits an answer for the current question and updates the progression.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param answerQuizQuestion body models.AnswerQuizQuestionRequest true "Answer details"
// @Success 200 {object} models.AnswerQuizQuestionResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes/answer [post]
func (h *QuizHandler) answerQuizQuestion(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => AnswerQuizQuestion invoked", r.Method, r.URL.Path)
	request, err := util.ReadBodyAndUnmarshal(models.AnswerQuizQuestionRequest{}, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get progression to check if it is okay to answer new questions
	// if it is ok, get question that we are going to answer
	var progression models.Progression
	res := h.db.First(&progression, request.ProgressionID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	if progression.IsFinished {
		http.Error(w, "quizHandler: quiz is already finished", http.StatusBadRequest)
		return
	}

	// check if question belongs to the quiz that is being done
	// check if question has that option that user is trying to select
	// if all good select option and save answer
	var question models.Question
	res = h.db.Preload("Options").First(&question, progression.CurrentQuestionID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	if question.QuizID != progression.QuizID {
		http.Error(w, "quizHandler: question does not belong to this quiz", http.StatusBadRequest)
		return
	}

	isOptionInQuestion := false
	for _, o := range question.Options {
		if o.ID == request.OptionID {
			isOptionInQuestion = true
			break
		}
	}
	if !isOptionInQuestion {
		http.Error(w, "quizHandler: chosen option does not belong to this question", http.StatusBadRequest)
		return
	}

	// Save new answer to db
	answer := models.Answer{
		UserID:   progression.UserID,
		OptionID: request.OptionID,
		QuizID:   progression.QuizID,
	}
	res = h.db.Create(&answer)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Get quiz to fetch new question for progression or finish the progression
	var quiz models.Quiz
	res = h.db.Preload(clause.Associations).First(&quiz, progression.QuizID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	progression.QuestionNumber++
	if len(quiz.Questions) > progression.QuestionNumber {
		progression.CurrentQuestionID = quiz.Questions[progression.QuestionNumber].ID
	} else {
		progression.IsFinished = true
	}

	// save progression
	res = h.db.Save(&progression)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := models.AnswerQuizQuestionResponse{
		Progression: progression,
	}
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// finalizeQuiz
// @Summary Finalize a quiz
// @Description Marks a quiz as finished and calculates the score based on correct answers.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param finalizeQuiz body models.FinalizeQuizRequest true "Quiz finalization details"
// @Success 200 {object} models.FinalizeQuizResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes/submit [post]
func (h *QuizHandler) calculateScore(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => CalculateScore invoked", r.Method, r.URL.Path)
	request, err := util.ReadBodyAndUnmarshal(models.FinalizeQuizRequest{}, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var progression models.Progression
	res := h.db.Preload(clause.Associations).First(&progression, request.ProgressionID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}
	progression.IsFinished = true

	var quiz models.Quiz
	res = h.db.Preload(clause.Associations).First(&quiz, progression.QuizID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	totalQuestionCount := len(quiz.Questions)
	if totalQuestionCount == 0 {
		http.Error(w, "quizHandler: quiz does not have any questions", http.StatusInternalServerError)
		return
	}
	// get answers given to the quiz by the user
	var answers []models.Answer
	res = h.db.Where("user_id = ? AND quiz_id = ?", progression.UserID, progression.QuizID).Find(&answers)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}
	optionIds := make([]uint, len(answers))
	for i, a := range answers {
		optionIds[i] = a.OptionID
	}
	// get options where answers belong to user and quiz
	var options []models.Option
	res = h.db.Where("id IN ?", optionIds).Find(&options)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	correctAnswerCount := 0
	for _, o := range options {
		if o.IsCorrect {
			correctAnswerCount++
		}
	}
	calculatedScore := float32(correctAnswerCount) / float32(totalQuestionCount)
	score := models.Score{
		QuizID: progression.QuizID,
		UserID: progression.UserID,
		Score:  calculatedScore,
	}
	res = h.db.Create(&score)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Quiz has been submitted so progression is not needed anymore
	res = h.db.Delete(&models.Progression{}, progression.ID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}
	response := models.FinalizeQuizResponse{
		Score: score,
	}
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)

}

// getQuestion
// @Summary Get a quiz question by ID
// @Description Retrieves a question by its ID along with its answer options.
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} models.QuestionWithOptionsResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /quizzes/questions/{id} [get]
func (h *QuizHandler) getQuestion(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => GetQuestion invoked", r.Method, r.URL.Path)
	questionId := r.PathValue("id")

	var question models.Question
	res := h.db.Preload("Options").First(&question, questionId)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	options := make([]models.OptionBase, len(question.Options))
	for i, o := range question.Options {
		options[i] = models.OptionBase{
			Base:       o.Base,
			QuestionID: o.QuestionID,
			Value:      o.Value,
		}
	}

	response := models.QuestionWithOptionsResponse{
		Base:     question.Base,
		Question: question.Question,
		Answers:  options,
		QuizID:   question.QuizID,
	}
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}
