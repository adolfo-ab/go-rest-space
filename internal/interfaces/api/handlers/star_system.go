package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"main/internal/application/services"
	"main/internal/domain/aggregates/star_system"
	"net/http"
)

type StarSystemHandler struct {
	service services.StarSystemService
}

func NewStarSystemHandler() *StarSystemHandler {
	return &StarSystemHandler{}
}

func (h *StarSystemHandler) GetStarSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id := uuid.MustParse(idString)

	st, err := h.service.GetStarSystem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(st)
}

func (h *StarSystemHandler) CreateStarSystem(w http.ResponseWriter, r *http.Request) {
	var starSystem star_system.StarSystem
	err := json.NewDecoder(r.Body).Decode(&starSystem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateStarSystem(starSystem.GetName())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *StarSystemHandler) UpdateStarSystem(w http.ResponseWriter, r *http.Request) {
	var starSystem star_system.StarSystem
	err := json.NewDecoder(r.Body).Decode(&starSystem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateStarSystem(starSystem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *StarSystemHandler) DeleteStarSystem(w http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(r.URL.Query().Get("id"))

	err := h.service.DeleteStarSystem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
