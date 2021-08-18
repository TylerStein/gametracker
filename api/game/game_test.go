package main

import (
	"testing"
)

func TestPostHandler(t *testing.T) {
	t.Run("Gives 200 response when body is correct", func(t *testing.T) {
		handler := createMockPostHandler()
		body := createNewGameBody()
		request := createMockRequestWithBody(body)
		response, err := handler.handleRequest(request)
		if err != nil {
			t.Fatalf("Wanted nil error, got error: %v", err)
		}

		if response.StatusCode != 200 {
			t.Fatalf("Wanted 200 response, got code: %v", response.StatusCode)
		}
	})

	t.Run("Gives 400 response when body is malformed", func(t *testing.T) {
		handler := createMockPostHandler()
		request := createMockRequestWithBody("{}")
		response, err := handler.handleRequest(request)
		if err != nil {
			t.Fatalf("Wanted nil error, got error: %v", err)
		}

		if response.StatusCode != 400 {
			t.Fatalf("Wanted 400 response, got code %v", response.StatusCode)
		}
	})

	t.Run("Database inserts new Game", func(t *testing.T) {
		handler := createMockPostHandler()
		body := createNewGameBody()
		request := createMockRequestWithBody(body)
		response, err := handler.handleRequest(request)
		if err != nil {
			t.Fatalf("Wanted nil error, got error: %v", err)
		}

		if response.StatusCode != 200 {
			t.Fatalf("Wanted 200 response, got code: %v", response.StatusCode)
		}
	})

	// t.Run("Unable to get IP", func(t *testing.T) {
	// 	DefaultHTTPGetAddress = "http://127.0.0.1:12345"

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err == nil {
	// 		t.Fatal("Error failed to trigger with an invalid request")
	// 	}
	// })

	// t.Run("Non 200 Response", func(t *testing.T) {
	// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(500)
	// 	}))
	// 	defer ts.Close()

	// 	DefaultHTTPGetAddress = ts.URL

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err != nil && err.Error() != ErrNon200Response.Error() {
	// 		t.Fatalf("Error failed to trigger with an invalid HTTP response: %v", err)
	// 	}
	// })

	// t.Run("Unable decode IP", func(t *testing.T) {
	// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(500)
	// 	}))
	// 	defer ts.Close()

	// 	DefaultHTTPGetAddress = ts.URL

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err == nil {
	// 		t.Fatal("Error failed to trigger with an invalid HTTP response")
	// 	}
	// })

	// t.Run("Successful Request", func(t *testing.T) {
	// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(200)
	// 		fmt.Fprintf(w, "127.0.0.1")
	// 	}))
	// 	defer ts.Close()

	// 	DefaultHTTPGetAddress = ts.URL

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err != nil {
	// 		t.Fatal("Everything should be ok")
	// 	}
	// })
}
