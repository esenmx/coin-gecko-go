package gocko

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestSimplePrices(t *testing.T) {
	var sps SimplePrices
	sps.vsCurrencies = []string{"usd", "aud"}
	err := unmarshalModel("simple_price", &sps)
	require.NoError(t, err)
	require.Equal(t, 2, len(sps.Prices))
	require.Equal(t, 12.2, sps.Prices["polkadot"].CurrencyPrice["usd"].Price)
	require.Equal(t, 16.59, sps.Prices["polkadot"].CurrencyPrice["aud"].Price)
	require.Equal(t, 12307700548.971405, *sps.Prices["polkadot"].CurrencyPrice["usd"].MarketCap)
	require.Equal(t, 16723457351.931345, *sps.Prices["polkadot"].CurrencyPrice["aud"].MarketCap)
	require.Equal(t, 911849298.6238484, *sps.Prices["polkadot"].CurrencyPrice["usd"].Vol24h)
	require.Equal(t, 1239366417.8542655, *sps.Prices["polkadot"].CurrencyPrice["aud"].Vol24h)
	require.Equal(t, 9.111769973736187, *sps.Prices["polkadot"].CurrencyPrice["usd"].Change24h)
	require.Equal(t, 8.739671660788401, *sps.Prices["polkadot"].CurrencyPrice["aud"].Change24h)
	require.Equal(t, 25.63, sps.Prices["solana"].CurrencyPrice["usd"].Price)
	require.Equal(t, 34.84, sps.Prices["solana"].CurrencyPrice["aud"].Price)
	require.Nil(t, sps.Prices["solana"].CurrencyPrice["usd"].Vol24h)
	require.Nil(t, sps.Prices["solana"].CurrencyPrice["aud"].Vol24h)
}

func TestMarkets(t *testing.T) {
	var ms []Market
	err := unmarshalModel("coins_markets", &ms)
	require.NoError(t, err)
	require.Equal(t, 2, len(ms))
	eth := ms[0]
	usdt := ms[1]
	require.Equal(t, "ethereum", eth.Id)
	require.Equal(t, "tether", usdt.Id)
	require.Nil(t, usdt.SparklineIn7D)
	require.Nil(t, usdt.PriceChangePercentage1HInCurrency)
	require.Nil(t, usdt.Roi)
	require.Equal(t, "btc", eth.Roi.Currency)
	require.Equal(t, 168, len(eth.SparklineIn7D.Price))
	require.Equal(t, time.Date(2015, 3, 2, 0, 0, 0, 0, time.UTC), usdt.AtlDate)
}

func TestCoinData(t *testing.T) {
	var cd CoinData
	err := unmarshalModel("coins_id", &cd)
	require.NoError(t, err)
	require.Equal(t, "https://github.com/ethereum/go-ethereum",
		((cd.Links["repos_url"]).(map[string]interface{})["github"]).([]interface{})[0])
}

func unmarshalModel(filename string, ptr interface{}) error {
	bs, err := os.ReadFile(fmt.Sprintf("mock/%s.json", filename))
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}
