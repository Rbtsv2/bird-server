package bird

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// RouteConfig stocke les configurations pour une route, y compris le fichier et les paramètres.
type RouteConfig struct {
	File     string
	Param    string
	FuncName string // Ajout du nom de la fonction pour permettre l'exécution dynamique.
}

// RouteHandler gère le stockage et le chargement des routes.
type RouteHandler struct {
	Routes map[string]RouteConfig // Map des chemins de route vers les configurations de route correspondantes
}

// NewRouteHandler crée une nouvelle instance de RouteHandler.
func NewRouteHandler() *RouteHandler {
	return &RouteHandler{
		Routes: make(map[string]RouteConfig),
	}
}

func (r *RouteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if routeConfig, ok := r.Routes[path]; ok {
		fullPath := filepath.Join("/var/www/server/template", routeConfig.File)
		log.Printf("Serving file: %s with params: %s", fullPath, routeConfig.Param)

		// Parse the template file
		tmpl, err := template.ParseFiles(fullPath)
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Create a data map based on parameters
		data := map[string]interface{}{}
		paramParts := strings.Split(routeConfig.Param, "=")
		if len(paramParts) == 2 {
			data[paramParts[0]] = paramParts[1]
		}

		// Execute the template with data
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		log.Printf("No route found for path: %s", path)
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// LoadRoutesFromFile charge les routes à partir d'un fichier .bird spécifié.
func (r *RouteHandler) LoadRoutesFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s, %v", filePath, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentRoute string
	var blockIndentation int = -1 // Start with no block indentation

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		currentIndentation := len(line) - len(trimmedLine)

		// Detect route definition
		if strings.HasPrefix(trimmedLine, "route(") && strings.HasSuffix(trimmedLine, ")") {
			routePath := strings.Trim(trimmedLine[len("route("):len(trimmedLine)-1], "\"")
			currentRoute = routePath
			blockIndentation = currentIndentation + 4 // Assuming block indentation is 4 spaces more
			continue
		}

		// Check if we are inside a route block by indentation
		if blockIndentation >= 0 && currentIndentation == blockIndentation && strings.HasPrefix(trimmedLine, "return ") {
			parts := strings.SplitN(trimmedLine[len("return "):], ",", 2)
			htmlFile := strings.Trim(parts[0], " \"")
			param := ""
			if len(parts) > 1 {
				param = strings.Trim(parts[1], " \";")
			}
			r.Routes[currentRoute] = RouteConfig{File: htmlFile, Param: param}
			log.Printf("Loaded route: %s -> %s with param: %s", currentRoute, htmlFile, param)
			blockIndentation = -1 // Exit the block
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %s, %v", filePath, err)
		return err
	}
	return nil
}

// LoadRoutesFromDirectory charge les routes à partir de tous les fichiers .bird dans un répertoire.
func (r *RouteHandler) LoadRoutesFromDirectory(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".bird" {
			filePath := filepath.Join(dirPath, file.Name())
			err := r.LoadRoutesFromFile(filePath)
			if err != nil {
				log.Printf("Error loading bird files from %s: %v", filePath, err)
			} else {
				log.Printf("Routes loaded from %s successfully.", filePath)
			}
		}
	}
	return nil
}
