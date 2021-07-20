package gocko

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var client = DefaultClient()

func TestClient_Ping(t *testing.T) {
	p, err := client.Ping()
	assert.NoError(t, err)
	assert.True(t, strings.Contains(p.GeckoSays, "To the Moon!"))
}

func TestClient_SupportedVsCurrencies(t *testing.T) {
	scs, err := client.SimpleSupportedVsCurrencies()
	assert.NoError(t, err)
	assert.NotEmpty(t, len(scs))
	assert.NotEmpty(t, len(scs[0]))
}

func TestClient_SimplePrice(t *testing.T) {
	res, err := client.SimplePrice(SimplePriceParams{Ids: []ID{"polkadot", "solana", "chainlink", "kusama"}})
	assert.NoError(t, err)
	assert.Equal(t, 4, len(res))
	assert.NotNil(t, res["polkadot"])
	assert.NotNil(t, res["solana"])
	assert.NotNil(t, res["chainlink"])
	assert.NotNil(t, res["kusama"])
}

func TestClient_Coins(t *testing.T) {
	cs, err := client.CoinsList(CoinsParams{})
	assert.NoError(t, err)
	assert.NotEmpty(t, cs)
	assert.NotEmpty(t, cs[0].Id)
	assert.NotEmpty(t, cs[0].Symbol)
	assert.NotEmpty(t, cs[0].Name)
}

func TestClient_CoinsMarkets(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		ms, err := client.CoinsMarkets(CoinsMarketsParams{
			VsCurrency:            "usd",
			PerPage:               42,
			Sparkline:             true,
			PriceChangePercentage: "1h,24h,7d",
		})
		assert.NoError(t, err)
		assert.Equal(t, 42, len(*ms))
		first := (*ms)[0]
		assert.Equal(t, 168, len(first.SparklineIn7D.Price))
		assert.NotNil(t, first.PriceChangePercentage1HInCurrency)
		assert.NotNil(t, first.PriceChangePercentage24HInCurrency)
		assert.NotNil(t, first.PriceChangePercentage7DInCurrency)
		for _, v := range *ms {
			assert.NotEmpty(t, v.Id)
			assert.NotEmpty(t, v.Symbol)
			assert.NotEmpty(t, v.Name)
			assert.NotEmpty(t, v.Image)
			assert.NotEmpty(t, v.CurrentPrice)
			assert.NotEmpty(t, v.MarketCap)
			//assert.NotEmpty(t, v.TotalVolume)
			//assert.NotEmpty(t, v.CirculatingSupply)
			//assert.NotEmpty(t, v.TotalSupply)
			//assert.NotEmpty(t, v.MaxSupply)
			assert.NotNil(t, v.SparklineIn7D)
		}
	})
	t.Run("Ids", func(t *testing.T) {
		ms, err := client.CoinsMarkets(CoinsMarketsParams{VsCurrency: "usd", Ids: []ID{"polkadot", "solana"}})
		assert.NoError(t, err)
		assert.Equal(t, 2, len(*ms))
	})
}

func TestClient_CoinsCharts(t *testing.T) {
	tsIntegrity := func(ccs *Charts) {
		for i, v := range ccs.Prices {
			assert.Equal(t, v[0], ccs.MarketCaps[i][0])
			assert.Equal(t, v[0], ccs.TotalVolumes[i][0])
		}
	}
	t.Run("Minutely", func(t *testing.T) {
		ccs, err := client.CoinsCharts(CoinsChartsParams{Id: "polkadot", VsCurrency: "usd", Days: "1"})
		assert.NoError(t, err)
		assert.Equal(t, 12*24+1, len(ccs.Prices))
		assert.Equal(t, 12*24+1, len(ccs.TotalVolumes))
		assert.Equal(t, 12*24+1, len(ccs.MarketCaps))
		tsIntegrity(ccs)
	})
	t.Run("Hourly", func(t *testing.T) {
		ccs, err := client.CoinsCharts(CoinsChartsParams{Id: "polkadot", VsCurrency: "usd", Days: "7"})
		assert.NoError(t, err)
		assert.Equal(t, 7*24+1, len(ccs.Prices))
		assert.Equal(t, 7*24+1, len(ccs.TotalVolumes))
		assert.Equal(t, 7*24+1, len(ccs.MarketCaps))
		tsIntegrity(ccs)
	})
	t.Run("Daily", func(t *testing.T) {
		ccs, err := client.CoinsCharts(CoinsChartsParams{Id: "polkadot", VsCurrency: "usd", Days: "100"})
		assert.NoError(t, err)
		assert.Equal(t, 100+1, len(ccs.Prices))
		assert.Equal(t, 100+1, len(ccs.TotalVolumes))
		assert.Equal(t, 100+1, len(ccs.MarketCaps))
		tsIntegrity(ccs)
	})
}

func TestClient_CoinsData(t *testing.T) {
	cd, err := client.CoinsData(CoinsDataParams{Id: "solana"})
	assert.NoError(t, err)
	assert.Equal(t, "solana", cd.Id)
	assert.Equal(t, "sol", cd.Symbol)
	assert.Equal(t, "Solana", cd.Name)
}

func TestClient_CoinsOHLC(t *testing.T) {
	t.Run("Minutely", func(t *testing.T) {
		ohlc, err := client.CoinsOHLC(CoinsOHLCParams{Id: "polkadot", VsCurrency: "usd", Days: "1"})
		assert.NoError(t, err)
		assert.Equal(t, 2*24+1, len(*ohlc))
	})
	t.Run("Hourly", func(t *testing.T) {
		ohlc, err := client.CoinsOHLC(CoinsOHLCParams{Id: "polkadot", VsCurrency: "usd", Days: "7"})
		assert.NoError(t, err)
		assert.Equal(t, 6*7+1, len(*ohlc))
	})
	t.Run("Daily", func(t *testing.T) {
		ohlc, err := client.CoinsOHLC(CoinsOHLCParams{Id: "polkadot", VsCurrency: "usd", Days: "90"})
		assert.NoError(t, err)
		assert.Equal(t, 22+2+1, len(*ohlc))
	})
}

func TestClient_ExchangesList(t *testing.T) {
	el, err := client.ExchangesList()
	assert.NoError(t, err)
	assert.NotEmpty(t, el)
	assert.NotEmpty(t, (*el)[0].Id)
	assert.NotEmpty(t, (*el)[0].Name)
}

func TestClient_Exchanges(t *testing.T) {
	es, err := client.Exchanges(ExchangesParams{})
	assert.NoError(t, err)
	assert.NotEmpty(t, es)
	assert.NotEmpty(t, (*es)[0].Id)
	assert.NotEmpty(t, (*es)[0].Name)
}
