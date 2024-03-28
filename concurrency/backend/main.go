package main

import (
	"bytes"
	"encoding/base64"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/disintegration/gift"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create new sse with cors headers
	s := sse.NewServer(&sse.Options{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, OPTIONS",
			"Access-Control-Allow-Headers": "Keep-Alive,X-Requested-With,Cache-Control,Content-Type,Last-Event-ID",
		},
	})
	defer s.Shutdown()
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
	// Mount the SSE server to handle SSE events
	v1Router.Mount("/events/", s)
	// Handle image upload
	v1Router.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		// Parse multipart form with a max memory size
		err := r.ParseMultipartForm(10 << 20) // Max 10 MB files
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get file from form data
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Read uploaded file
		imgData, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Decode image (only accepting png on the frontend)
		img, err := png.Decode(bytes.NewReader(imgData))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get filter type
		filterType := r.FormValue("filter")

		// Process image
		go func() {
			// Process image with filter
			processedImg := applyFilter(img, filterType)
			// Convert the processedImg to a format suitable for SSE (Base64)
			var buf bytes.Buffer
			err = png.Encode(&buf, processedImg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
			w.WriteHeader(http.StatusAccepted)
			//Send sse event with the encoded image
			s.SendMessage("/v1/events/imageProcessed", sse.SimpleMessage(encodedStr))
		}()
	})
	r.Mount("/v1", v1Router)

	log.Println("Starting server....")
	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

// applyFilter applies the selected filter to the image and returns the image with the filter applied
func applyFilter(img image.Image, filterType string) image.Image {
	g := gift.New()
	switch filterType {
	case "pixelate":
		g.Add(gift.Pixelate(5))
	case "grayscale":
		g.Add(gift.Grayscale())
	case "invert":
		g.Add(gift.Invert())
	case "sepia":
		g.Add(gift.Sepia(100))
	}

	dst := image.NewNRGBA(g.Bounds(img.Bounds()))
	g.Draw(dst, img)
	return dst
}
