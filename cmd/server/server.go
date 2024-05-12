package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"connect.com/connect/pkg/bird" // Assurez-vous que le chemin d'import est correct
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Construction du chemin du fichier basé sur l'URL demandée
		// Nous utilisons /var/www/server comme base pour les fichiers .bird
		filePath := filepath.Join("/var/www/server", r.URL.Path)

		// Lecture du contenu du fichier .bird
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			// Si le fichier n'est pas trouvé ou qu'une erreur de lecture se produit, renvoyer une erreur 404
			http.Error(w, "File not found or error reading file.", http.StatusNotFound)
			return
		}

		// Interpréter le contenu du fichier .bird avec l'interpréteur du langage Bird
		output, err := bird.Interpret(string(content))
		if err != nil {
			// Gérer les erreurs d'interprétation en renvoyant une erreur 500
			http.Error(w, fmt.Sprintf("Error interpreting content: %s", err), http.StatusInternalServerError)
			return
		}

		// Envoyer le résultat de l'interprétation comme réponse HTTP
		fmt.Fprintf(w, output)
	})

	// Démarrer le serveur sur le port 8080 et loguer les erreurs
	log.Fatal(http.ListenAndServe(":8080", nil))
}
