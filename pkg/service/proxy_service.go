package service

import (
	"context"
	"errors"
	"kmf-proxy/pkg/domain"
	"kmf-proxy/pkg/repository/httpclient"
	"kmf-proxy/pkg/repository/storage"
	"net/http"
)

type Service struct {
	httpClient    domain.HttpClientRepo
	storageClient domain.StorageClientRepo
}

func NewProxyService(httpClient *httpclient.Client, storageClient *storage.Client) *Service {
	return &Service{
		httpClient:    httpClient,
		storageClient: storageClient,
	}
}

func (p *Service) Proxy(ctx context.Context, request domain.Request) (domain.Response, error) {
	uuid, ok := ctx.Value("uuid").(string)
	if !ok {
		return domain.Response{}, errors.New("couldn't generate uuid")
	}

	resp, err := p.httpClient.Do(ctx, request)
	if err != nil {
		return domain.Response{}, err
	}

	go func(id string, request domain.Request, response http.Response) {
		p.storageClient.Add(id, request, *resp)
	}(uuid, request, *resp)

	return domain.Response{
		ID:      uuid,
		Status:  resp.Status,
		Headers: resp.Header,
		Length:  resp.ContentLength,
	}, nil
}
