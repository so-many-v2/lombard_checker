package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"lombardchecker/internal/config"
	"lombardchecker/internal/entities"
	"lombardchecker/pkg/logger"
	"math/big"
	"net/http"
	"net/url"
	"time"
)

type Fetcher struct {
	proxy  string // residental proxy
	log    *logger.Logger
	client *http.Client
}

func NewFetcher(conf *config.Config, log *logger.Logger) *Fetcher {

	transport := http.Transport{}

	if conf.UseProxy {

		proxyUrl, err := url.Parse(conf.ProxyAddress)

		if err != nil {
			log.Fatal("error parse proxy url: %s", err.Error())
		}

		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	return &Fetcher{
		log: log,
		client: &http.Client{
			Timeout:   time.Second * 10,
			Transport: &transport,
		},
	}
}

func (f *Fetcher) FetchWalletAllocation(address string) (int64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	url := fmt.Sprintf("https://mainnet.prod.lombard.finance/api/v1/bard/distributor/%s/claimable", address)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("error preparing request for address: %s", address)
	}

	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Set("origin", "https://claim.lombard.finance")
	req.Header.Set("referer", "https://claim.lombard.finance")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"140\", \"Not=A?Brand\";v=\"24\", \"Google Chrome\";v=\"140\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")

	resp, err := f.client.Do(req)

	if err != nil {
		return 0, fmt.Errorf("error getting request for address: %s | %s", address, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0, fmt.Errorf("error read body from response | address : %s\n", address)
	}

	var data entities.WalletData

	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("error parse data => not eligible | address : %s\n", address)
	}

	if len(data.Data) == 0 {
		return 0, fmt.Errorf("error no data (not eligible) | address : %s\n", address)
	}

	amount := new(big.Int)
	amount, ok := amount.SetString(data.Data[0].Amount, 10)
	if !ok {
		return 0, fmt.Errorf("failed to parse big amount for address: %s\n", address)
	}

	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	realAmount := new(big.Int).Div(amount, decimals)

	return realAmount.Int64(), nil
}

func (f *Fetcher) CheckProxy() error {
	resp, err := f.client.Get("https://api.myip.com")
	if err != nil {
		return fmt.Errorf("proxy check error: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Proxy check response: %s", string(body))
	return nil
}
