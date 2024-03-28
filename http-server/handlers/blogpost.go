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

// GetBlogposts gets a list of all the blogposts currently in the list
func (bp BlogpostHandler) GetBlogposts(w http.ResponseWriter, r *http.Request) {
	// Get the blogposts from the model and encode it into json
	err := json.NewEncoder(w).Encode(models.GetBlogposts())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// GetBlogpost gets a specific blogpost based on a specific /{id}
func (bp BlogpostHandler) GetBlogpost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	// Convert the string id into an integer and return an error if it fails
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to convert id to integer: %v", idStr), http.StatusInternalServerError)
		return
	}
	// Get the blogpost from the model and encode it into json
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

// CreateBlogpost creates a new blogpost and adds it to the list of blogposts
func (bp BlogpostHandler) CreateBlogpost(w http.ResponseWriter, r *http.Request) {
	var blogpost models.BlogPost
	// Decode the request body into the blogpost variable and print an error if it fails
	err := json.NewDecoder(r.Body).Decode(&blogpost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Send the new blogpost to the model and encode it to send back in json
	models.CreateBlogpost(blogpost)
	err = json.NewEncoder(w).Encode(blogpost)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// UpdateBlogpost updates a specific blogpost based on a specific /{id}
func (bp BlogpostHandler) UpdateBlogpost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	// Convert the string id into an integer and return an error if it fails
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to convert id to integer: %v", idStr), http.StatusInternalServerError)
		return
	}
	var blogpost models.BlogPost
	// Decode the request body into the blogpost variable and print an error if it fails
	err = json.NewDecoder(r.Body).Decode(&blogpost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update the blogpost in the model and returns the updated blogpost or nil if it failed to find it
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
	// Convert the string id into an integer and return an error if it fails
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to convert id to integer: %v", idStr), http.StatusInternalServerError)
		return
	}
	// Delete the blogpost in the model and returns a bool to say if it was deleted successfully or not
	deleted := models.DeleteBlogpost(id)
	if deleted == false {
		http.Error(w, "Blogpost not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
