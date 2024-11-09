package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ApiError struct {
	Error string
}

type APIServer struct {
	store      Storage
	listenAddr string
}

func MakeHTTPHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	indentedJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(indentedJSON)
	return err
}

func newAPIServer(store Storage, listenAddr string) *APIServer {
	return &APIServer{
		store:      store,
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/clothes", MakeHTTPHandler(s.handleGetClothes))
	router.HandleFunc("POST /api/clothes", MakeHTTPHandler(s.handleCreateClothes))
	router.HandleFunc("GET /api/clothes/{id}", MakeHTTPHandler(s.handleGetClothesByID))
	router.HandleFunc("DELETE /api/clothes/{id}", MakeHTTPHandler(s.handleDeleteClothes))
	router.HandleFunc("PATCH /api/clothes/{id}", MakeHTTPHandler(s.handleWearClothes))
	router.HandleFunc("PUT /api/clothes/{id}", MakeHTTPHandler(s.handleUpdateClothes))

	log.Println("Listening on", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func getID(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid ID: %s", idStr)
	}
	return id, nil
}

// Get /clothes
func (s *APIServer) handleGetClothes(w http.ResponseWriter, r *http.Request) error {
	clothes, err := s.store.GetClothing()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, clothes)
}

// Get /clothes by ID
func (s *APIServer) handleGetClothesByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	article, err := s.store.GetArticleByID(id)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, ApiError{Error: "Article not found: " + r.PathValue("id")})
	}

	return WriteJSON(w, http.StatusOK, article)
}

// Post /clothes
func (s *APIServer) handleCreateClothes(w http.ResponseWriter, r *http.Request) error {
	article := new(Clothing)
	if err := json.NewDecoder(r.Body).Decode(article); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid JSON body"})
	}

	article.UpdateCPW()
	if err := s.store.CreateArticle(article); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to create article"})
	}

	if err := WriteJSON(w, http.StatusCreated, article); err != nil {
		return err
	}

	return nil
}

// Delete /clothes by ID
func (s *APIServer) handleDeleteClothes(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	err = s.store.DeleteArticle(id)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, ApiError{Error: "Article not found: " + r.PathValue("id")})
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// Patch /clothes by ID, incrment wears and last worn
func (s *APIServer) handleWearClothes(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	article, err := s.store.GetArticleByID(id)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, ApiError{Error: "Article not found: " + r.PathValue("id")})
	}
	article.IncrementWears()

	err = s.store.UpdateArticle(article)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to update article"})
	}

	return WriteJSON(w, http.StatusOK, article)
}

func (s *APIServer) handleUpdateClothes(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	article := new(Clothing)
	if err := json.NewDecoder(r.Body).Decode(article); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid JSON body"})
	}

	existingArticle, err := s.store.GetArticleByID(id)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, ApiError{Error: "Article not found"})
	}

	lastWear := existingArticle.LastWorn
	article.LastWorn = lastWear

	article.ID = id
	article.UpdateCPW()
	err = s.store.UpdateArticle(article)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to update article"})
	}

	return WriteJSON(w, http.StatusOK, article)
}
