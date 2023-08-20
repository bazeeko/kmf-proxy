package domain

import (
	"context"
	"net/http"
)

type HttpClientRepo interface {
	Do(ctx context.Context, request Request) (*http.Response, error)
}

type StorageClientRepo interface {
	Add(id string, request Request, response http.Response)
	Get(id string) (request Request, response http.Response, exists bool)
}

type ProxyService interface {
	Proxy(ctx context.Context, request Request) (Response, error)
}
