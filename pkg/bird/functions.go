package bird

import (
	"fmt"
	"log"
)

// FunctionMap est une carte des noms de fonction aux fonctions exécutables.
type FunctionMap struct {
	functions map[string]func(args ...Expr) (string, error)
}

// NewFunctionMap crée une nouvelle instance de FunctionMap.
func NewFunctionMap() *FunctionMap {
	fm := &FunctionMap{
		functions: make(map[string]func(args ...Expr) (string, error)),
	}
	fm.registerDefaults()
	return fm
}

// registerDefaults enregistre les fonctions par défaut.
func (fm *FunctionMap) registerDefaults() {

	fm.functions["log"] = func(args ...Expr) (string, error) {

		if len(args) < 1 {
			return "", fmt.Errorf("not enough arguments to log")
		}
		log.Printf("Log: %s", args[0].Evaluate()) // Affiche le log dans la console pour le débogage
		// Simple log implementation that just returns the string representation of the first argument.
		return args[0].Evaluate(), nil

	}

}

// Call exécute la fonction avec les arguments donnés.
func (fm *FunctionMap) Call(funcName string, args ...Expr) (string, error) {
	if function, exists := fm.functions[funcName]; exists {
		return function(args...)
	}
	return "", fmt.Errorf("function %s not defined", funcName)
}

// AddFunction ajoute une fonction personnalisée à la carte.
func (fm *FunctionMap) AddFunction(name string, function func(args ...Expr) (string, error)) {
	fm.functions[name] = function
}
