package gocko

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://api.coingecko.com/api/v3"

type Client struct {
	httpClient *http.Client
}

type Option func(*Client)

func NewClient(options ...Option) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		option(c)
	}
	return c
}
func WithHttpClient(hc *http.Client) Option { return func(c *Client) { c.httpClient = hc } }

func (c *Client) Do(url string, params QueryParams, ptr interface{}) error {
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

type Api interface {
	Ping() (Ping, error)
	SimpleSupportedVsCurrencies() ([]string, error)
	SimplePrice(SimplePriceParams) (SimplePrices, error)
	CoinsList(CoinsParams) ([]Coin, error)
	CoinsMarkets(CoinsMarketsParams) ([]Market, error)
	CoinsID(CoinsDataParams) (CoinData, error)
	CoinsMarketCharts(CoinsChartsParams) (Charts, error)
	CoinsOHLC(CoinsOHLCParams) (OHLC, error)
	ExchangesList() (ExchangeList, error)
	Exchanges(params ExchangesParams) ([]Exchange, error)
}

func assertApiInterface() {
	var _ Api = (*Client)(nil)
}
