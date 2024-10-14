package grpc

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/AntonLuning/tiny-url/api"
	"github.com/AntonLuning/tiny-url/proto"
	"github.com/AntonLuning/tiny-url/service"
)

const (
	DOMAIN = "localhost"
	PORT   = 9988
)

func TestMain(m *testing.M) {
	setupTestServer()

	code := m.Run()

	os.Exit(code)
}

func setupTestServer() {
	urlService := service.NewShortenURLService("testing.com")

	grpcServer := api.NewGRPCAPIServer(fmt.Sprintf("%s:%d", DOMAIN, PORT), urlService)

	go func() {
		if err := grpcServer.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(100 * time.Millisecond) // Give the server a moment to start
}

func TestCreateShortenURL(t *testing.T) {
	client, err := New(fmt.Sprintf("%s:%d", DOMAIN, PORT))
	if err != nil {
		t.Fatal(err)
	}

	_, err = createShortenURL(*client, "example.com/my/test/url")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetOriginalURL(t *testing.T) {
	client, err := New(fmt.Sprintf("%s:%d", DOMAIN, PORT))
	if err != nil {
		t.Fatal(err)
	}

	exampleURL := "example.com/my/test/url"

	shortenURL, err := createShortenURL(*client, exampleURL)
	if err != nil {
		t.Fatal(err)
	}

	originalURL, err := getOriginalURL(*client, *shortenURL)
	if err != nil {
		t.Fatal(err)
	}

	if exampleURL != *originalURL {
		t.Fatalf("original URL did not match what server responed with original: %s respOriginal: %s", exampleURL, *originalURL)
	}
}

func createShortenURL(client proto.ServiceClient, originalURL string) (*string, error) {
	resp, err := client.CreateShortenURL(context.Background(), &proto.ShortenURLRequest{Original: originalURL})
	if err != nil {
		return nil, err
	}

	if resp.Original != originalURL {
		return nil, fmt.Errorf("original URL did not match what server responed with original: %s respOriginal: %s", originalURL, resp.Original)
	}

	return &resp.Shorten, nil
}

func getOriginalURL(client proto.ServiceClient, shortenURL string) (*string, error) {
	resp, err := client.GetOriginalURL(context.Background(), &proto.OriginalURLRequest{Shorten: shortenURL})
	if err != nil {
		return nil, err
	}

	if resp.Shorten != shortenURL {
		return nil, fmt.Errorf("shorten URL did not match what server responed with shorten: %s respShorten: %s", shortenURL, resp.Shorten)
	}

	return &resp.Original, nil
}
