package bird

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Interface Expr pour toutes les expressions.
type Expr interface {
	Evaluate() string
}

// Number pour les nombres.
type Number struct {
	Value float64
}

func (n *Number) Evaluate() string {
	return fmt.Sprintf("%f", n.Value)
}

// StringExpr pour les chaînes.
type StringExpr struct {
	Value string
}

func (s *StringExpr) Evaluate() string {
	return s.Value
}

// NewStringExpr crée une nouvelle expression de chaîne.
func NewStringExpr(value string) Expr {
	return &StringExpr{Value: value}
}

// BinaryExpr pour les opérations binaires.
type BinaryExpr struct {
	Op    string
	Left  Expr
	Right Expr
}

func (b *BinaryExpr) Evaluate() string {
	leftVal, _ := strconv.ParseFloat(b.Left.Evaluate(), 64)
	rightVal, _ := strconv.ParseFloat(b.Right.Evaluate(), 64)
	switch b.Op {
	case "+":
		return fmt.Sprintf("%f", leftVal+rightVal)
	case "-":
		return fmt.Sprintf("%f", leftVal-rightVal)
	case "*":
		return fmt.Sprintf("%f", leftVal*rightVal)
	case "/":
		return fmt.Sprintf("%f", leftVal/rightVal)
	default:
		panic(fmt.Sprintf("unsupported operator: %s", b.Op))
	}
}

// FunctionCallExpr pour les appels de fonction.
type FunctionCallExpr struct {
	Name string
	Args []Expr
}

func (f *FunctionCallExpr) Evaluate() string {
	if f.Name == "log" && len(f.Args) > 0 {
		return f.Args[0].Evaluate() // Assume yo simply returns its argument.
	}
	return "Error: Unknown function or incorrect arguments"
}

// PArseExp est concu pour traiter des expressions mathématiques ou des appels de fonction
func ParseExpr(input string) (Expr, error) {
	log.Printf("Parsing expression: %s", input)
	input = strings.TrimSpace(input) // Enlève les espaces blancs superflus

	// Si l'entrée est un chemin de fichier ou une chaîne littérale non fonctionnelle
	if strings.HasSuffix(input, ".html") || !strings.Contains(input, "(") {
		return &StringExpr{Value: input}, nil
	}

	// Gestion des chaînes de caractères
	if strings.HasPrefix(input, "\"") && strings.HasSuffix(input, "\"") {
		value := input[1 : len(input)-1]
		log.Printf("Detected string expression: %s", value)
		return &StringExpr{Value: value}, nil
	}

	// Gestion des appels de fonction
	if strings.Contains(input, "(") && strings.Contains(input, ")") {
		functionName := strings.TrimSpace(input[:strings.Index(input, "(")])
		argumentContent := input[strings.Index(input, "(")+1 : strings.LastIndex(input, ")")]
		log.Printf("Detected function call: %s with arguments %s", functionName, argumentContent)
		args, err := ParseArguments(argumentContent)
		if err != nil {
			return nil, err
		}
		return &FunctionCallExpr{Name: functionName, Args: args}, nil
	}

	// Gestion des nombres
	if val, err := strconv.ParseFloat(input, 64); err == nil {
		log.Printf("Detected numeric expression: %f", val)
		return &Number{Value: val}, nil
	}

	log.Printf("Failed to parse expression: unsupported format")
	return nil, fmt.Errorf("unsupported expression format")
}

// ParseArguments analyse les arguments d'un appel de fonction.
func ParseArguments(input string) ([]Expr, error) {
	log.Printf("Parsing arguments: %s", input)
	parts := strings.Split(input, ",")
	var args []Expr
	for _, part := range parts {
		expr, err := ParseExpr(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		args = append(args, expr)
	}
	return args, nil
}

// Evaluate prend une chaîne en entrée et renvoie le résultat de son évaluation.
func Evaluate(input string) string {
	expr, err := ParseExpr(input)
	if err != nil {
		return ""
	}
	return expr.Evaluate()
}

// Interpret prend une chaîne en entrée et interprète le code comme du langage Bird, puis renvoie le résultat.
// Interpret fonction à jour pour utiliser FunctionMap.
func Interpret(input string, fm *FunctionMap) (string, error) {
	expr, err := ParseExpr(input)
	if err != nil {
		return "", fmt.Errorf("error parsing expression: %v", err)
	}
	if callExpr, ok := expr.(*FunctionCallExpr); ok {
		return fm.Call(callExpr.Name, callExpr.Args...)
	}
	return expr.Evaluate(), nil
}
