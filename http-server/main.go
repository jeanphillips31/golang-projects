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

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environment")
	}

	r := chi.NewRouter()

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
	v1Router.Mount("/blogposts", BlogpostRoutes())
	r.Mount("/v1", v1Router)

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
