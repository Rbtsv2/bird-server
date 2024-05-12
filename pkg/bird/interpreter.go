package bird

import (
	"fmt"
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
	if f.Name == "yo" && len(f.Args) > 0 {
		return f.Args[0].Evaluate() // Assume yo simply returns its argument.
	}
	return "Error: Unknown function or incorrect arguments"
}

func ParseExpr(input string) (Expr, error) {
	input = strings.TrimSpace(input) // Enlève les espaces blancs superflus

	// Gestion des chaînes simples
	if strings.HasPrefix(input, "\"") && strings.HasSuffix(input, "\"") {
		return &StringExpr{Value: strings.Trim(input, "\"")}, nil
	}

	// Gestion des appels de fonction
	if strings.Contains(input, "(") && strings.Contains(input, ")") {
		functionName := strings.TrimSpace(input[:strings.Index(input, "(")])
		arguments := strings.TrimSpace(input[strings.Index(input, "(")+1 : strings.LastIndex(input, ")")])
		if functionName == "yo" {
			argExpr, err := ParseExpr(arguments)
			if err != nil {
				return nil, err
			}
			return &FunctionCallExpr{Name: functionName, Args: []Expr{argExpr}}, nil
		}
		return nil, fmt.Errorf("unsupported function: %s", functionName)
	}

	// Gestion des nombres
	if val, err := strconv.ParseFloat(input, 64); err == nil {
		return &Number{Value: val}, nil
	}

	return nil, fmt.Errorf("unsupported expression format")
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
func Interpret(input string) (string, error) {
	expr, err := ParseExpr(input) // Assurez-vous que ParseExpr et Evaluate gèrent les erreurs appropriées.
	if err != nil {
		return "", fmt.Errorf("error parsing expression: %v", err)
	}
	result := expr.Evaluate()
	return result, nil
}
