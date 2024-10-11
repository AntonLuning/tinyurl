package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/AntonLuning/tiny-url/models"
	"github.com/AntonLuning/tiny-url/service"
	"github.com/AntonLuning/tiny-url/utils"
)

type apiFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	addr    string
	service service.Service
}

func NewJSONAPIServer(addr string, service service.Service) *JSONAPIServer {
	return &JSONAPIServer{
		addr:    addr,
		service: service,
	}
}

func (s *JSONAPIServer) Run() error {
	http.Handle("/api/v1/", http.StripPrefix("/api/v1", s.v1Mux()))

	slog.Info("JSON API server starting", "address", s.addr)

	return http.ListenAndServe(s.addr, nil)
}

func (s *JSONAPIServer) v1Mux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /url", makeHTTPHandlerFunc(s.handleCreateShortenURL))
	mux.HandleFunc("GET /url", makeHTTPHandlerFunc(s.handleGetShortenURL))

	return mux
}

func (s *JSONAPIServer) handleCreateShortenURL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	originalURL := r.URL.Query().Get("original")
	if originalURL == "" {
		return fmt.Errorf("query param \"original\" must not be empty")
	}

	shortenURL, err := s.service.CreateShortenURL(ctx, originalURL)
	if err != nil {
		return err
	}

	resp := models.CreateShortenURLResponse{
		Original: originalURL,
		Shorten:  *shortenURL,
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *JSONAPIServer) handleGetShortenURL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	shortenURL := r.URL.Query().Get("shorten")
	if shortenURL == "" {
		return fmt.Errorf("query param \"shorten\" must not be empty")
	}

	originalURL, err := s.service.GetOriginalURL(ctx, shortenURL)
	if err != nil {
		return err
	}

	resp := models.GetOriginalURLResponse{
		Shorten:  shortenURL,
		Original: *originalURL,
	}

	return writeJSON(w, http.StatusOK, resp)
}

func makeHTTPHandlerFunc(apiFn apiFunc) http.HandlerFunc {
	ctx := context.WithValue(context.Background(), utils.REQUEST_ID_KEY, 1)

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			// TODO: handle errors more dynamically (custom error type)
			_ = writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, content any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(content)
}
