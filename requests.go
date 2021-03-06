package gocko

import (
	"fmt"
)

func (c *Client) Ping() (Ping, error) {
	var p Ping
	err := c.Do(fmt.Sprintf("%s/ping", baseURL), nil, &p)
	return p, err
}

//
// Simple
//

func (c *Client) SimpleSupportedVsCurrencies() ([]string, error) {
	var scs []string
	err := c.Do(fmt.Sprintf("%s/simple/supported_vs_currencies", baseURL), nil, &scs)
	return scs, err
}

func (c *Client) SimplePrice(p SimplePriceParams) (SimplePrices, error) {
	sps := SimplePrices{vsCurrencies: p.VsCurrencies}
	err := c.Do(fmt.Sprintf("%s/simple/price", baseURL), p, &sps)
	return sps, err
}

//
// Coins
//

func (c Client) CoinsList(p CoinsParams) ([]Coin, error) {
	var cs []Coin
	err := c.Do(fmt.Sprintf("%s/coins/list", baseURL), p, &cs)
	if len(cs[0].Id) == 0 {
		cs = cs[1:]
	}
	return cs, err
}

func (c *Client) CoinsMarkets(p CoinsMarketsParams) ([]Market, error) {
	var ms []Market
	err := c.Do(fmt.Sprintf("%s/coins/markets", baseURL), p, &ms)
	return ms, err
}

func (c *Client) CoinsID(p CoinsDataParams) (CoinData, error) {
	var cd CoinData
	err := c.Do(fmt.Sprintf("%s/coins/%s", baseURL, p.Id), p, &cd)
	return cd, err
}

func (c *Client) CoinsMarketCharts(p CoinsChartsParams) (Charts, error) {
	var ccs Charts
	err := c.Do(fmt.Sprintf("%s/coins/%s/market_chart", baseURL, p.Id), p, &ccs)
	return ccs, err
}

func (c *Client) CoinsOHLC(p CoinsOHLCParams) (OHLC, error) {
	var ohlc OHLC
	err := c.Do(fmt.Sprintf("%s/coins/%s/ohlc", baseURL, p.Id), p, &ohlc)
	return ohlc, err
}

//
// Exchanges
//

func (c *Client) ExchangesList() (ExchangeList, error) {
	var el ExchangeList
	err := c.Do(fmt.Sprintf("%s/exchanges/list", baseURL), nil, &el)
	return el, err
}

func (c *Client) Exchanges(p ExchangesParams) ([]Exchange, error) {
	var es []Exchange
	err := c.Do(fmt.Sprintf("%s/exchanges", baseURL), p, &es)
	return es, err
}
