package gocko

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://api.coingecko.com/api/v3"

type Client struct {
	httpClient *http.Client
}

func DefaultClient() *Client {
	return &Client{httpClient: http.DefaultClient}
}

// CustomClient you may consider using: https://github.com/hashicorp/go-retryablehttp
func CustomClient(client *http.Client) *Client {
	return &Client{httpClient: client}
}

func (c *Client) do(url string, params QueryParams, ptr interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if params != nil {
		qp, err := params.toQuery()
		if err != nil {
			return err
		}
		q := req.URL.Query()
		for k, v := range qp {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(ptr)
}

type api interface {
	Ping() (Ping, error)
	SimpleSupportedVsCurrencies() ([]string, error)
	SimplePrice(SimplePriceParams) (SimplePrices, error)
	CoinsList(CoinsParams) ([]Coin, error)
	CoinsMarkets(CoinsMarketsParams) ([]Market, error)
	CoinsData(CoinsDataParams) (CoinData, error)
	CoinsCharts(CoinsChartsParams) (Charts, error)
	CoinsOHLC(CoinsOHLCParams) (OHLC, error)
	ExchangesList() (ExchangeList, error)
	Exchanges(params ExchangesParams) ([]Exchange, error)
}

var _ api = (*Client)(nil)
