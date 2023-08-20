package handler

import (
	"kmf-proxy/pkg/domain"
	"kmf-proxy/pkg/service"
	"net/http"
)

type Handler struct {
	proxyService domain.ProxyService
}

func NewHandler(proxyService *service.Service) *Handler {
	return &Handler{
		proxyService: proxyService,
	}
}

func (h *Handler) Init() http.Handler {
	router := http.NewServeMux()

	router.Handle("/proxy", middlewareRequestID(http.HandlerFunc(h.Proxy)))

	return router
}
