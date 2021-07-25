package gocko

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestSimplePrices(t *testing.T) {
	var sps SimplePrices
	sps.vsCurrencies = []string{"usd", "aud"}
	err := unmarshalModel("simple_price", &sps)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(sps.Prices))
	assert.Equal(t, 12.2, sps.Prices["polkadot"].CurrencyPrice["usd"].Price)
	assert.Equal(t, 16.59, sps.Prices["polkadot"].CurrencyPrice["aud"].Price)
	assert.Equal(t, 12307700548.971405, *sps.Prices["polkadot"].CurrencyPrice["usd"].MarketCap)
	assert.Equal(t, 16723457351.931345, *sps.Prices["polkadot"].CurrencyPrice["aud"].MarketCap)
	assert.Equal(t, 911849298.6238484, *sps.Prices["polkadot"].CurrencyPrice["usd"].Vol24h)
	assert.Equal(t, 1239366417.8542655, *sps.Prices["polkadot"].CurrencyPrice["aud"].Vol24h)
	assert.Equal(t, 9.111769973736187, *sps.Prices["polkadot"].CurrencyPrice["usd"].Change24h)
	assert.Equal(t, 8.739671660788401, *sps.Prices["polkadot"].CurrencyPrice["aud"].Change24h)
	assert.Equal(t, 25.63, sps.Prices["solana"].CurrencyPrice["usd"].Price)
	assert.Equal(t, 34.84, sps.Prices["solana"].CurrencyPrice["aud"].Price)
	assert.Nil(t, sps.Prices["solana"].CurrencyPrice["usd"].Vol24h)
	assert.Nil(t, sps.Prices["solana"].CurrencyPrice["aud"].Vol24h)
}

func TestMarkets(t *testing.T) {
	var ms []Market
	err := unmarshalModel("coins_markets", &ms)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(ms))
	eth := ms[0]
	usdt := ms[1]
	assert.Equal(t, "ethereum", eth.Id)
	assert.Equal(t, "tether", usdt.Id)
	assert.Nil(t, usdt.SparklineIn7D)
	assert.Nil(t, usdt.PriceChangePercentage1HInCurrency)
	assert.Nil(t, usdt.Roi)
	assert.Equal(t, "btc", eth.Roi.Currency)
	assert.Equal(t, 168, len(eth.SparklineIn7D.Price))
	assert.Equal(t, time.Date(2015, 3, 2, 0, 0, 0, 0, time.UTC), usdt.AtlDate)
}

func TestCoinData(t *testing.T) {
	var cd CoinData
	err := unmarshalModel("coins_id", &cd)
	assert.NoError(t, err)
	assert.Equal(t, "https://github.com/ethereum/go-ethereum",
		((cd.Links["repos_url"]).(map[string]interface{})["github"]).([]interface{})[0])
}

func unmarshalModel(filename string, ptr interface{}) error {
	bs, err := os.ReadFile(fmt.Sprintf("mock/%s.json", filename))
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}
