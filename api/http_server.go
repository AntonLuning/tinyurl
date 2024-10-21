package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/AntonLuning/tiny-url/models"
	"github.com/AntonLuning/tiny-url/service"
	"github.com/AntonLuning/tiny-url/utils"
)

type apiFunc func(context.Context, http.ResponseWriter, *http.Request) error

type HTTPServer struct {
	addr    string
	service service.Service
	withAPI bool
}

func NewHTTPServer(addr string, service service.Service, withAPI bool) *HTTPServer {
	return &HTTPServer{
		addr:    addr,
		service: service,
		withAPI: withAPI,
	}
}

func (s *HTTPServer) Run(basePath string) error {
	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("GET %s/{id}", basePath), makeHTTPHandlerFunc(s.handleUsage))

	if s.withAPI {
		mux.Handle("/api/v1/", http.StripPrefix("/api/v1", s.v1Mux()))
		slog.Info("server starting, including JSON API", "address", s.addr)
	} else {
		slog.Info("server starting", "address", s.addr)
	}

	return http.ListenAndServe(s.addr, mux)
}

func (s *HTTPServer) v1Mux() http.Handler {
	v1Mux := http.NewServeMux()

	v1Mux.HandleFunc("POST /url", makeHTTPHandlerFunc(s.handleCreateShortenURL))
	v1Mux.HandleFunc("GET /url", makeHTTPHandlerFunc(s.handleGetShortenURL))

	return v1Mux
}

func (s *HTTPServer) handleCreateShortenURL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	originalURL := r.URL.Query().Get("original")

	shortenURL, err := s.service.CreateShortenURL(ctx, originalURL)
	if err != nil {
		switch e := err.(type) {
		case *service.EmptyInputError:
			return NewValidationError(e.Value, "query param missing")
		default:
			return err
		}
	}

	resp := models.CreateShortenURLResponse{
		Original: originalURL,
		Shorten:  *shortenURL,
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *HTTPServer) handleGetShortenURL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	shortenURL := r.URL.Query().Get("shorten")
	originalURL, err := s.fetchOriginal(ctx, shortenURL)
	if err != nil {
		return err
	}

	resp := models.GetOriginalURLResponse{
		Shorten:  shortenURL,
		Original: *originalURL,
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *HTTPServer) handleUsage(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	shortenURL := r.URL.String()
	originalURL, err := s.fetchOriginal(ctx, shortenURL)
	if err != nil {
		return err
	}

	http.Redirect(w, r, ensureProtocolPrefix(*originalURL), http.StatusFound)

	return nil
}

func makeHTTPHandlerFunc(apiFn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := utils.SetContextValues(context.Background(), "HTTP")

		slog.Info("Incoming request", "server", "HTTP", "method", r.Method, "path", r.URL.Path, "requestID", ctx.Value(utils.REQUEST_ID_KEY))

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

func (s *HTTPServer) fetchOriginal(ctx context.Context, shortenURL string) (*string, error) {
	originalURL, err := s.service.GetOriginalURL(ctx, shortenURL)
	if err != nil {
		switch e := err.(type) {
		case *service.EmptyInputError:
			return nil, NewValidationError(e.Value, "query param missing")
		case *service.ShortenNotExistError:
			return nil, NewNotFoundError("shorten url", e.Value)
		default:
			return nil, err
		}
	}

	return originalURL, nil
}

func writeJSON(w http.ResponseWriter, statusCode int, content any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(content)
}

func ensureProtocolPrefix(targetURL string) string {
	if strings.HasPrefix(targetURL, "http://") || strings.HasPrefix(targetURL, "https://") {
		return targetURL
	}
	if !strings.HasPrefix(targetURL, "//") {
		return "//" + targetURL
	}
	return targetURL
}
