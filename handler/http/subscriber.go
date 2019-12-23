package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ishvaram/betal-kamailio/driver"
	models "github.com/ishvaram/betal-kamailio/models"
	repository "github.com/ishvaram/betal-kamailio/repository"
	subs "github.com/ishvaram/betal-kamailio/repository/subscriber"
)

// NewSubscriberHandler ...
func NewSubscriberHandler(db *driver.DB) *Subscriber {
	return &Subscriber{
		repo: subs.NewSQLSubscriberRepo(db.SQL),
	}
}

// Subscriber ...
type Subscriber struct {
	repo repository.SubscriberRepo
}

// Fetch all Subscriber data
func (sub *Subscriber) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := sub.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new Subscriber
func (sub *Subscriber) Create(w http.ResponseWriter, r *http.Request) {
	subs := models.Subscriber{}
	json.NewDecoder(r.Body).Decode(&subs)

	_, err := sub.repo.Create(r.Context(), &subs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a Subscriber by id
func (sub *Subscriber) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Subscriber{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := sub.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a Subscriber details
func (sub *Subscriber) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := sub.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Subscriber not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a Subscriber
func (sub *Subscriber) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := sub.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
