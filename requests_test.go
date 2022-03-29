package gocko

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var client = NewClient()

func TestClient_Ping(t *testing.T) {
	p, err := client.Ping()
	require.NoError(t, err)
	require.True(t, strings.Contains(p.GeckoSays, "To the Moon!"))
}

func TestClient_SupportedVsCurrencies(t *testing.T) {
	scs, err := client.SimpleSupportedVsCurrencies()
	require.NoError(t, err)
	require.NotEmpty(t, len(scs))
	require.NotEmpty(t, len(scs[0]))
}

func TestClient_SimplePrice(t *testing.T) {
	res, err := client.SimplePrice(SimplePriceParams{
		Ids:          []string{"polkadot", "solana", "chainlink", "kusama"},
		VsCurrencies: []string{"usd", "aud"},
	})
	require.NoError(t, err)
	require.Equal(t, 4, len(res.Prices))
	require.Equal(t, 2, len(res.vsCurrencies))
	require.NotNil(t, res.Prices["polkadot"])
	require.NotNil(t, res.Prices["solana"])
	require.NotNil(t, res.Prices["chainlink"])
	require.NotNil(t, res.Prices["kusama"])
}

func TestClient_Coins(t *testing.T) {
	cs, err := client.CoinsList(CoinsParams{})
	require.NoError(t, err)
	require.NotEmpty(t, cs)
	require.NotEmpty(t, cs[0].Id)
	require.NotEmpty(t, cs[0].Symbol)
	require.NotEmpty(t, cs[0].Name)
}

func TestClient_CoinsMarkets(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		ms, err := client.CoinsMarkets(CoinsMarketsParams{
			VsCurrency:            "usd",
			PerPage:               42,
			Sparkline:             true,
			PriceChangePercentage: "1h,24h,7d",
		})
		require.NoError(t, err)
		require.Equal(t, 42, len(ms))
		first := ms[0]
		require.Equal(t, 168, len(first.SparklineIn7D.Price))
		require.NotNil(t, first.PriceChangePercentage1HInCurrency)
		require.NotNil(t, first.PriceChangePercentage24HInCurrency)
		require.NotNil(t, first.PriceChangePercentage7DInCurrency)
		for _, v := range ms {
			require.NotEmpty(t, v.Id)
			require.NotEmpty(t, v.Symbol)
			require.NotEmpty(t, v.Name)
			require.NotEmpty(t, v.Image)
			require.NotEmpty(t, v.CurrentPrice)
			require.NotEmpty(t, v.MarketCap)
			//require.NotEmpty(t, v.TotalVolume)
			//require.NotEmpty(t, v.CirculatingSupply)
			//require.NotEmpty(t, v.TotalSupply)
			//require.NotEmpty(t, v.MaxSupply)
			require.NotNil(t, v.SparklineIn7D)
		}
	})
	t.Run("Ids", func(t *testing.T) {
		ms, err := client.CoinsMarkets(CoinsMarketsParams{VsCurrency: "usd", Ids: []string{"polkadot", "solana"}})
		require.NoError(t, err)
		require.Equal(t, 2, len(ms))
	})
}

func TestClient_CoinsCharts(t *testing.T) {
	tsIntegrity := func(ccs Charts) {
		for i, v := range ccs.Prices {
			require.Equal(t, v[0], ccs.MarketCaps[i][0])
			require.Equal(t, v[0], ccs.TotalVolumes[i][0])
		}
	}
	t.Run("Minutely", func(t *testing.T) {
		ccs, err := client.CoinsMarketCharts(CoinsChartsParams{Id: "polkadot", VsCurrency: "usd", Days: "1"})
		require.NoError(t, err)
		require.Equal(t, 12*24+1, len(ccs.Prices))
		require.Equal(t, 12*24+1, len(ccs.TotalVolumes))
		require.Equal(t, 12*24+1, len(ccs.MarketCaps))
		tsIntegrity(ccs)
	})
	t.Run("Hourly", func(t *testing.T) {
		ccs, err := client.CoinsMarketCharts(CoinsChartsParams{Id: "polkadot", VsCurrency: "usd", Days: "7"})
		require.NoError(t, err)
		require.Equal(t, 7*24+1, len(ccs.Prices))
		require.Equal(t, 7*24+1, len(ccs.TotalVolumes))
		require.Equal(t, 7*24+1, len(ccs.MarketCaps))
		tsIntegrity(ccs)
	})
	t.Run("Daily", func(t *testing.T) {
		ccs, err := client.CoinsMarketCharts(CoinsChartsParams{Id: "polkadot", VsCurrency: "usd", Days: "100"})
		require.NoError(t, err)
		require.Equal(t, 100+1, len(ccs.Prices))
		require.Equal(t, 100+1, len(ccs.TotalVolumes))
		require.Equal(t, 100+1, len(ccs.MarketCaps))
		tsIntegrity(ccs)
	})
}

func TestClient_CoinsData(t *testing.T) {
	cd, err := client.CoinsID(CoinsDataParams{Id: "solana"})
	require.NoError(t, err)
	require.Equal(t, "solana", cd.Id)
	require.Equal(t, "sol", cd.Symbol)
	require.Equal(t, "Solana", cd.Name)
}

func TestClient_CoinsOHLC(t *testing.T) {
	t.Run("Minutely", func(t *testing.T) {
		ohlc, err := client.CoinsOHLC(CoinsOHLCParams{Id: "polkadot", VsCurrency: "usd", Days: "1"})
		require.NoError(t, err)
		require.LessOrEqual(t, 48, len(ohlc))
	})
	t.Run("Hourly", func(t *testing.T) {
		ohlc, err := client.CoinsOHLC(CoinsOHLCParams{Id: "polkadot", VsCurrency: "usd", Days: "7"})
		require.NoError(t, err)
		require.LessOrEqual(t, 6*7, len(ohlc))
	})
	t.Run("Daily", func(t *testing.T) {
		ohlc, err := client.CoinsOHLC(CoinsOHLCParams{Id: "polkadot", VsCurrency: "usd", Days: "90"})
		require.NoError(t, err)
		require.LessOrEqual(t, 90/4, len(ohlc))
	})
}

func TestClient_ExchangesList(t *testing.T) {
	el, err := client.ExchangesList()
	require.NoError(t, err)
	require.NotEmpty(t, el)
	require.NotEmpty(t, el[0].Id)
	require.NotEmpty(t, el[0].Name)
}

func TestClient_Exchanges(t *testing.T) {
	es, err := client.Exchanges(ExchangesParams{})
	require.NoError(t, err)
	require.NotEmpty(t, es)
	require.NotEmpty(t, es[0].Id)
	require.NotEmpty(t, es[0].Name)
}
