package repository_ipinfo

import (
	"context"
	"encoding/json"
	"net/http"

	repository_config "github.com/hengkyawijaya/simple-go/repository/config"
)

type IPInfo struct {
	cfg    *repository_config.Config
	client *http.Client
}

type IPInfoResponse struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func NewIPInfo(ctx context.Context, cfg *repository_config.Config) *IPInfo {
	return &IPInfo{
		cfg:    cfg,
		client: http.DefaultClient,
	}
}

func (i *IPInfo) GetIPInfo() (string, error) {
	req, err := http.NewRequest(http.MethodGet, i.cfg.Repository.IPInfo.Endpoint, nil)
	if err != nil {
		return "", err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var ipInfo IPInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&ipInfo)
	if err != nil {
		return "", err
	}

	return ipInfo.Ip, nil
}
