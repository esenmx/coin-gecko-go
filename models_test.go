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
	err := unmarshalModel("simple_price", &sps)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(sps))
	assert.Equal(t, -9.42491022906792, *sps["polkadot"].Usd24HChange)
	assert.Equal(t, 24.56, sps["solana"].Usd)
	assert.Nil(t, sps["solana"].Usd24HVol)
}

func TestMarkets(t *testing.T) {
	var ms []Market
	err := unmarshalModel("coins_markets", &ms)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(ms))
	eth := ms[0]
	usdt := ms[1]
	assert.Equal(t, ID("ethereum"), eth.Id)
	assert.Equal(t, ID("tether"), usdt.Id)
	assert.Nil(t, usdt.SparklineIn7D)
	assert.Nil(t, usdt.PriceChangePercentage1HInCurrency)
	assert.Nil(t, usdt.Roi)
	assert.Equal(t, Currency("btc"), eth.Roi.Currency)
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
