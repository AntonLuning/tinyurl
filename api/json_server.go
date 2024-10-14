package api

import (
	"context"
	"encoding/json"
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
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", s.v1Mux()))

	slog.Info("JSON API server starting", "address", s.addr)

	return http.ListenAndServe(s.addr, mux)
}

func (s *JSONAPIServer) v1Mux() http.Handler {
	v1Mux := http.NewServeMux()

	v1Mux.HandleFunc("POST /url", makeHTTPHandlerFunc(s.handleCreateShortenURL))
	v1Mux.HandleFunc("GET /url", makeHTTPHandlerFunc(s.handleGetShortenURL))

	return v1Mux
}

func (s *JSONAPIServer) handleCreateShortenURL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	originalURL := r.URL.Query().Get("original")
	if originalURL == "" {
		return NewValidationError("original", "query param missing")
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
		return NewValidationError("shorten", "query param missing")
	}

	originalURL, err := s.service.GetOriginalURL(ctx, shortenURL)
	if err != nil {
		switch e := err.(type) {
		case *service.ShortenNotExistError:
			return NewNotFoundError("shorten url", e.Value)
		default:
			return err
		}
	}

	resp := models.GetOriginalURLResponse{
		Shorten:  shortenURL,
		Original: *originalURL,
	}

	return writeJSON(w, http.StatusOK, resp)
}

func makeHTTPHandlerFunc(apiFn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := utils.SetContextValues(context.Background(), "JSON")

		if err := apiFn(ctx, w, r); err != nil {
			switch e := err.(type) {
			case *ValidationError:
				writeJSON(w, http.StatusBadRequest, map[string]any{"error": "bad request", "message": e})
			case *NotFoundError:
				writeJSON(w, http.StatusNotFound, map[string]any{"error": "not found", "message": e})
			default:
				writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
			}
		}
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, content any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(content)
}
