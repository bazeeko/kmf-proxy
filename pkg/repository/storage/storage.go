package storage

import (
	"kmf-proxy/pkg/domain"
	"net/http"
	"sync"
)

type item struct {
	request  domain.Request
	response http.Response
}

type Client struct {
	cache map[string]item
	mx    sync.Mutex
}

func NewRepository() *Client {
	return &Client{
		cache: make(map[string]item),
	}
}

func (s *Client) Add(id string, request domain.Request, response http.Response) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.cache[id] = item{
		request:  request,
		response: response,
	}
}

func (s *Client) Get(id string) (request domain.Request, response http.Response, exists bool) {
	s.mx.Lock()
	defer s.mx.Unlock()

	it, exists := s.cache[id]
	if !exists {
		return domain.Request{}, http.Response{}, false
	}

	return it.request, it.response, true
}
