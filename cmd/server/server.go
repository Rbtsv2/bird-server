package main

import (
	"log"
	"net/http"

	"connect.com/connect/pkg/bird" // Assurez-vous que le chemin d'import est correct
)

func main() {

	handler := bird.NewRouteHandler()
	err := handler.LoadRoutesFromDirectory("/var/www/server/birdfiles")
	if err != nil {
		log.Fatalf("Failed to load routes: %v", err)
	}

	// Configure le gestionnaire pour toutes les requêtes avec "/"
	http.HandleFunc("/", handler.ServeHTTP)

	// Un point de terminaison pour déboguer les routes chargées
	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		for path, config := range handler.Routes {
			log.Printf("Route: %s, File: %s, Param: %s\n", path, config.File, config.Param)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
