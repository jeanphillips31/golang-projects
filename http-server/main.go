package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jeanphillips31/golang-projects/http-server/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Get port from environment file
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environment")
	}

	r := chi.NewRouter()
	// Set cors options (this allows basically everything so should only be run locally)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	// Mount blogpost routes
	v1Router.Mount("/blogposts", BlogpostRoutes())
	// Mouse v1 route (can be used when updating the api to create a v2, v3, etc.)
	r.Mount("/v1", v1Router)

	log.Println("Starting server....")
	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

func BlogpostRoutes() chi.Router {
	r := chi.NewRouter()
	blogpostHandler := handlers.BlogpostHandler{}
	r.Get("/", blogpostHandler.GetBlogposts)
	r.Post("/", blogpostHandler.CreateBlogpost)
	r.Get("/{id}", blogpostHandler.GetBlogpost)
	r.Put("/{id}", blogpostHandler.UpdateBlogpost)
	r.Delete("/{id}", blogpostHandler.DeleteBlogpost)
	return r
}
