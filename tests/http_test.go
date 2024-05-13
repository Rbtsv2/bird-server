package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"connect.com/connect/pkg/bird" // Assurez-vous que ce chemin est correct
)

func TestRouteHandler(t *testing.T) {
	handler := bird.NewRouteHandler() // Notez l'utilisation de bird. pour accéder à NewRouteHandler
	err := handler.LoadRoutesFromDirectory("/var/www/bird/birdfiles")
	if err != nil {
		t.Fatalf("Failed to load routes: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(handler.ServeHTTP))
	defer server.Close()

	tests := []struct {
		route  string
		expect string
		status int
	}{
		{"/contact", "expected content for contact", http.StatusOK},
		{"/nonexistent", "", http.StatusNotFound},
	}

	for _, test := range tests {
		resp, err := http.Get(server.URL + test.route)
		if err != nil {
			t.Errorf("Failed to get %s: %v", test.route, err)
			continue
		}
		if resp.StatusCode != test.status {
			t.Errorf("Expected status %d for route %s, got %d", test.status, test.route, resp.StatusCode)
		}
		// Additional checks can be added here to validate response body etc.
	}
}
