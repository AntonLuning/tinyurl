package service

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsService struct {
	next Service

	addr            string
	registry        *prometheus.Registry
	createdCounter  prometheus.Counter
	fetchedCounter  *prometheus.CounterVec
	serviceDuration *prometheus.HistogramVec
}

func NewMetricsService(next Service, addr string) Service {
	s := MetricsService{
		next:     next,
		addr:     addr,
		registry: prometheus.NewRegistry(),
	}

	s.createMetrics()
	go s.exposeMetrics()

	return &s
}

func (s *MetricsService) CreateShortenURL(ctx context.Context, originalURL string) (shortenURL *string, err error) {
	defer func(start time.Time) {
		dur := time.Since(start).Seconds()
		status := "ok"
		if err != nil {
			status = "failed"
		} else {
			s.createdCounter.Inc()
		}
		s.serviceDuration.WithLabelValues("create", status).Observe(dur)
	}(time.Now())

	return s.next.CreateShortenURL(ctx, originalURL)
}

func (s *MetricsService) GetOriginalURL(ctx context.Context, shortenURL string) (originalURL *string, err error) {
	defer func(start time.Time) {
		dur := time.Since(start).Seconds()
		status := "ok"
		if err != nil {
			status = "failed"
		} else {
			s.fetchedCounter.WithLabelValues(shortenURL, *originalURL).Inc()
		}
		s.serviceDuration.WithLabelValues("fetch", status).Observe(dur)
	}(time.Now())

	return s.next.GetOriginalURL(ctx, shortenURL)
}

func (s *MetricsService) createMetrics() {
	s.registry.MustRegister(collectors.NewGoCollector())

	s.createdCounter = promauto.With(s.registry).NewCounter(prometheus.CounterOpts{
		Name: "service_create_shorten_url",
	})

	s.fetchedCounter = promauto.With(s.registry).NewCounterVec(prometheus.CounterOpts{
		Name: "service_fetch_shorten_url",
	}, []string{"shorten", "original"})

	s.serviceDuration = promauto.With(s.registry).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "service_duration_seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "status"},
	)
}

func (s *MetricsService) exposeMetrics() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(s.registry, promhttp.HandlerOpts{}))

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}
	slog.Info("Running metrics server", "address", s.addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Metrics server failed", "error", err.Error())
	}
}
