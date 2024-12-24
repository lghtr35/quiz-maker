package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/lghtr35/quiz-maker/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserHandler struct {
	db      *gorm.DB
	decoder schema.Decoder
}

func newUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db, decoder: *schema.NewDecoder()}
}

func (h *UserHandler) ConfigureSelf(m *http.ServeMux) *http.ServeMux {
	m.HandleFunc("GET /users", h.readUsers)
	m.HandleFunc("GET /users/{id}", h.readUserWithID)
	m.HandleFunc("POST /users", h.createUsers)
	m.HandleFunc("PATCH /users", h.updateUsers)
	m.HandleFunc("DELETE /users/{id}", h.deleteUser)
	m.HandleFunc("GET /users/{userId}/quiz/{quizId}/ranking", h.getUserRankingByScore)
	m.HandleFunc("GET /users/{userId}/quiz/{quizId}", h.getUserScoreForQuiz)

	return m
}

// readUsers
// @Summary Get a list of users
// @Description Retrieves a paginated list of users based on optional filters.
// @Tags Users
// @Accept json
// @Produce json
// @Param id_list query []uint false "List of user IDs"
// @Param name query string false "Name to search for"
// @Param page query int true "Page number"
// @Param size query int true "Page size"
// @Success 200 {array} models.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [get]
func (h *UserHandler) readUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => ReadUsers invoked", r.Method, r.URL.Path)
	var request models.ReadUsersRequest
	err := h.decoder.Decode(&request, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []models.User
	q := h.db.Model(&models.User{})
	if request.IDList != nil && len(*request.IDList) > 0 {
		q = q.Where("ID IN ?", *request.IDList)
	}
	if request.Name != nil && *request.Name != "" {
		// obtain a search string like '%name%'
		nameLike := fmt.Sprintf("%%%s%%", *request.Name)
		q = q.Where("name LIKE ?", nameLike)
	}
	offset := (request.Page - 1) * request.Size
	q = q.Offset(int(offset)).Limit(int(request.Size))
	res := q.Preload(clause.Associations).Find(&users)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(b)
}

// createUsers
// @Summary Create a new user
// @Description Creates a new user in the system.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User details"
// @Success 201 {object} models.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
func (h *UserHandler) createUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => CreateUsers invoked", r.Method, r.URL.Path)
	var request models.CreateUserRequest

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(b, &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Name: request.Name,
	}
	res := h.db.Create(user)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err = json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

// updateUsers
// @Summary Update an existing user
// @Description Updates user details based on the provided request body.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UpdateUserRequest true "Updated user details"
// @Success 200 {object} models.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [patch]
func (h *UserHandler) updateUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => UpdateUsers invoked", r.Method, r.URL.Path)
	var request models.UpdateUserRequest

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(b, &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := new(models.User)
	res := h.db.First(user, request.ID)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	if request.Name != nil && *request.Name != "" {
		user.Name = *request.Name
	}

	res = h.db.Save(user)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err = json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// readUserWithID
// @Summary Get a user by ID
// @Description Retrieves a user by their ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/{id} [get]
func (h *UserHandler) readUserWithID(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => ReadUserWithID invoked", r.Method, r.URL.Path)
	id := r.PathValue("id")

	var user models.User
	res := h.db.Preload(clause.Associations).First(&user, id)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// deleteUser
// @Summary Delete a user
// @Description Deletes a user by their ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/{id} [delete]
func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => DeleteUser invoked", r.Method, r.URL.Path)
	id := r.PathValue("id")

	res := h.db.Delete(&models.User{}, id)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
}

// getUserScoreForQuiz godoc
// @Summary      Get user's score for a specific quiz
// @Description  Retrieves the user's score for a specific quiz.
// @Tags         Users
// @Param        userId  path      string  true  "User ID"
// @Param        quizId  path      string  true  "Quiz ID"
// @Success      200     {object}  models.Score
// @Failure      500     {string}  string  "Internal server error"
// @Router       /users/{userId}/quiz/{quizId}  [get]
func (h *UserHandler) getUserScoreForQuiz(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => GetUserScoreForQuiz invoked", r.Method, r.URL.Path)
	userId := r.PathValue("userId")
	quizId := r.PathValue("quizId")

	var score models.Score
	res := h.db.Where("user_id = ? AND quiz_id = ?", userId, quizId).First(&score)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// getUserRankingByScore godoc
// @Summary      Get user's ranking by score in a specific quiz
// @Description  Retrieves the user's ranking, score, and percentage of quizzers they outperformed in a specific quiz.
// @Tags         Users
// @Param        userId  path      string  true  "User ID"
// @Param        quizId  path      string  true  "Quiz ID"
// @Success      200     {object}  models.GetUserRankingByScoreResponse
// @Failure      500     {string}  string  "Internal server error"
// @Router       /users/{userId}/quiz/{quizId}/ranking [get]
func (h *UserHandler) getUserRankingByScore(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s => GetUserRankingByScore invoked", r.Method, r.URL.Path)
	userIdStr := r.PathValue("userId")
	quizId := r.PathValue("quizId")

	temp, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userId := uint(temp)

	var scores []models.Score
	//Query
	res := h.db.Where("quiz_id = ?", quizId).Order("score desc").Find(&scores)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	totalOpponentCount := len(scores)
	userPlace := 0
	userScore := new(models.Score)
	for i, s := range scores {
		if s.UserID == userId {
			userPlace = i + 1
			userScore = &scores[i]
			break
		}
	}
	percent := (1 - (float32(userPlace) / float32(totalOpponentCount))) * 100
	response := models.GetUserRankingByScoreResponse{
		Percent: percent,
		Message: fmt.Sprintf("You were better than %.2f%% of all quizzers", percent),
		Score:   userScore,
		Rank:    uint(userPlace),
	}

	if percent == 0 && userPlace == 1 {
		response.Message = "You were the only person to finish this quiz yet."
	}

	if userPlace == 0 {
		response.Message = "Score of this quiz has not been found."
		response.Percent = 0
		response.Score = nil
		response.Rank = 0
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}
