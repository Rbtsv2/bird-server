package main

import (
	"testing"

	"connect.com/connect/pkg/bird"
)

func TestParseExpr(t *testing.T) {
	testExpressions := map[string]string{
		"42":                "42.000000",
		"7 + 3":             "10.000000",
		"\"Hello, World!\"": "Hello, World!",
		"yo(\"Ca y est, tu es un génie du code\")": "Ca y est, tu es un génie du code",
	}

	for expr, expected := range testExpressions {
		parsedExpr, err := bird.ParseExpr(expr)
		if err != nil {
			t.Errorf("Error parsing '%s': %s", expr, err)
		} else if parsedExpr.Evaluate() != expected {
			t.Errorf("Failed parsing '%s': expected %s, got %s", expr, expected, parsedExpr.Evaluate())
		}
	}
}
