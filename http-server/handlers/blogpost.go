package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jeanphillips31/golang-projects/http-server/models"
	"net/http"
	"strconv"
)

type BlogpostHandler struct {
}

func (bp BlogpostHandler) GetBlogposts(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.GetBlogposts())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (bp BlogpostHandler) GetBlogpost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to convert id to integer: %v", idStr), http.StatusInternalServerError)
		return
	}
	blogpost := models.GetBlogpost(id)
	if blogpost == nil {
		http.Error(w, fmt.Sprintf("Blogpost not found for id: %v", idStr), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(blogpost)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (bp BlogpostHandler) CreateBlogpost(w http.ResponseWriter, r *http.Request) {
	var blogpost models.BlogPost
	err := json.NewDecoder(r.Body).Decode(&blogpost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateBlogpost(blogpost)
	err = json.NewEncoder(w).Encode(blogpost)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (bp BlogpostHandler) UpdateBlogpost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to convert id to integer: %v", idStr), http.StatusInternalServerError)
		return
	}
	var blogpost models.BlogPost
	err = json.NewDecoder(r.Body).Decode(&blogpost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBlogpost := models.UpdateBlogpost(id, blogpost)
	if updatedBlogpost == nil {
		http.Error(w, "Blogpost not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedBlogpost)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (bp BlogpostHandler) DeleteBlogpost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to convert id to integer: %v", idStr), http.StatusInternalServerError)
		return
	}

	deleted := models.DeleteBlogpost(id)
	if deleted == false {
		http.Error(w, "Blogpost not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
