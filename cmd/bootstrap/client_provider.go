package bootstrap

import (
	"fmt"
	"myApiController/cmd/configs"
	"myApiController/internal/domain"
	infrastructure "myApiController/internal/infrastructure/client"
	"net/http"
	"time"
)

var clients = map[string]func(c configs.Client) (domain.DataRowClient, error){
	configs.Resource_GetRestApi:    buildHttpGetRestClient,
	configs.Resource_PostRestApi:   buildHttpPostRestClient,
	configs.Resource_DeleteRestApi: buildHttpDeleteRestClient,
}

func GetDataRowClient(c configs.Client) (domain.DataRowClient, error) {
	client, exists := clients[c.Type]
	if !exists {
		return nil, fmt.Errorf("unable to build %s client", c.Name)
	}

	return client(c)
}

func buildHttpGetRestClient(c configs.Client) (domain.DataRowClient, error) {
	httpClient := http.Client{
		Timeout: time.Second * 1000,
	}
	return infrastructure.NewRestApi(c.Path, http.MethodGet, c.Headers, httpClient), nil
}

func buildHttpPostRestClient(c configs.Client) (domain.DataRowClient, error) {
	httpClient := http.Client{
		Timeout: time.Second * 1000,
	}
	return infrastructure.NewRestApi(c.Path, http.MethodPost, c.Headers, httpClient), nil
}

func buildHttpDeleteRestClient(c configs.Client) (domain.DataRowClient, error) {
	httpClient := http.Client{
		Timeout: time.Second * 1000,
	}
	return infrastructure.NewRestApi(c.Path, http.MethodDelete, c.Headers, httpClient), nil
}
