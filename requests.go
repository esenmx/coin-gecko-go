package gocko

import (
	"fmt"
)

func (c *Client) Ping() (Ping, error) {
	var p Ping
	err := c.do(fmt.Sprintf("%s/ping", baseURL), nil, &p)
	return p, err
}

//
// Simple
//

func (c *Client) SimpleSupportedVsCurrencies() ([]string, error) {
	var scs []string
	err := c.do(fmt.Sprintf("%s/simple/supported_vs_currencies", baseURL), nil, &scs)
	return scs, err
}

func (c *Client) SimplePrice(p SimplePriceParams) (SimplePrices, error) {
	sps := SimplePrices{vsCurrencies: p.VsCurrencies}
	err := c.do(fmt.Sprintf("%s/simple/price", baseURL), p, &sps)
	return sps, err
}

//
// Coins
//

func (c Client) CoinsList(p CoinsParams) ([]Coin, error) {
	var cs []Coin
	err := c.do(fmt.Sprintf("%s/coins/list", baseURL), p, &cs)
	if len(cs[0].Id) == 0 {
		cs = cs[1:]
	}
	return cs, err
}

func (c *Client) CoinsMarkets(p CoinsMarketsParams) ([]Market, error) {
	var ms []Market
	err := c.do(fmt.Sprintf("%s/coins/markets", baseURL), p, &ms)
	return ms, err
}

func (c *Client) CoinsData(p CoinsDataParams) (CoinData, error) {
	var cd CoinData
	err := c.do(fmt.Sprintf("%s/coins/%s", baseURL, p.Id), p, &cd)
	return cd, err
}

func (c *Client) CoinsCharts(p CoinsChartsParams) (Charts, error) {
	var ccs Charts
	err := c.do(fmt.Sprintf("%s/coins/%s/market_chart", baseURL, p.Id), p, &ccs)
	return ccs, err
}

func (c *Client) CoinsOHLC(p CoinsOHLCParams) (OHLC, error) {
	var ohlc OHLC
	err := c.do(fmt.Sprintf("%s/coins/%s/ohlc", baseURL, p.Id), p, &ohlc)
	return ohlc, err
}

//
// Exchanges
//

func (c *Client) ExchangesList() (ExchangeList, error) {
	var el ExchangeList
	err := c.do(fmt.Sprintf("%s/exchanges/list", baseURL), nil, &el)
	return el, err
}

func (c *Client) Exchanges(p ExchangesParams) ([]Exchange, error) {
	var es []Exchange
	err := c.do(fmt.Sprintf("%s/exchanges", baseURL), p, &es)
	return es, err
}
