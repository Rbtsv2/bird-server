package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"connect.com/connect/pkg/bird" // Assurez-vous que le chemin d'import est correct
)

func main() {
	handler := bird.NewRouteHandler()
	fm := bird.NewFunctionMap() // Créez une instance de FunctionMap

	err := handler.LoadRoutesFromDirectory("/var/www/bird/birdfiles")
	if err != nil {
		log.Fatalf("Failed to load routes: %v", err)
	}

	// Configurez le gestionnaire pour toutes les requêtes avec "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requested path: %s", r.URL.Path)

		if routeConfig, ok := handler.Routes[r.URL.Path]; ok {
			// Tentative d'exécuter des fonctions si définies
			output, params, err := bird.Interpret(routeConfig.File, fm)
			if err != nil {
				log.Printf("Error executing function: %v", err)
			}

			// Prépare et rend le template HTML avec les paramètres dynamiques
			tmplPath := filepath.Join("/var/www/bird/template", routeConfig.File)
			tmpl, err := template.ParseFiles(tmplPath)
			if err != nil {
				log.Printf("Error parsing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			data := map[string]interface{}{
				"Content": output, // Résultat de `log()` ou autre fonction
				"Params":  params, // Paramètres supplémentaires de la fonction
			}

			// Injecte les paramètres de l'URL dans data
			paramParts := strings.Split(routeConfig.Param, "=")
			if len(paramParts) == 2 {
				data[paramParts[0]] = paramParts[1]
			}

			tmpl.Execute(w, data) // Exécute le template avec data
		} else {
			log.Printf("No route found for path: %s", r.URL.Path)
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})

	// Un point de terminaison pour déboguer les routes chargées
	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		for path, config := range handler.Routes {
			log.Printf("Route: %s, File: %s, Param: %s\n", path, config.File, config.Param)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
