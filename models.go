package gocko

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

type Coin struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type Image struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type Supply struct {
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	MaxSupply         float64 `json:"max_supply"`
}

type Price struct {
	Price     float64
	MarketCap *float64
	Vol24h    *float64
	Change24h *float64
}

type SimplePrice struct {
	CurrencyPrice map[string]Price
	LastUpdatedAt *int64
}

type SimplePrices struct {
	vsCurrencies []string
	Prices       map[string]SimplePrice
}

func (r *SimplePrices) UnmarshalJSON(bs []byte) error {
	if len(r.vsCurrencies) == 0 {
		return MissingParameterError
	}
	var data map[string]map[string]float64
	err := json.Unmarshal(bs, &data)
	if err != nil {
		return err
	}
	r.Prices = make(map[string]SimplePrice, len(data))
	for k, v := range data {
		lu := int64(v["last_updated_at"])
		ps := SimplePrice{LastUpdatedAt: &lu, CurrencyPrice: make(map[string]Price, len(r.vsCurrencies))}
		for _, vsc := range r.vsCurrencies {
			parser := func(k string) *float64 {
				if p, ok := v[fmt.Sprintf("%s_%s", vsc, k)]; ok {
					return &p
				}
				return nil
			}
			price := v[string(vsc)]
			ps.CurrencyPrice[vsc] = Price{
				Price:     price,
				MarketCap: parser("market_cap"),
				Vol24h:    parser("24h_vol"),
				Change24h: parser("24h_change"),
			}
		}
		r.Prices[string(k)] = ps
	}
	return nil
}

type StatusUpdate struct {
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	User        string    `json:"user"`
	UserTitle   string    `json:"user_title"`
	Pin         bool      `json:"pin"`
	Project     Project   `json:"project"`
}

type Description string
type Platforms map[string]string
type Project struct {
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Image  Image  `json:"image"`
}
type CoinData struct {
	Id                           string                 `json:"id"`
	Symbol                       string                 `json:"symbol"`
	Name                         string                 `json:"name"`
	AssetPlatformId              *string                `json:"asset_platform_id"`
	Platforms                    Platforms              `json:"platforms"`
	BlockTimeInMinutes           int                    `json:"block_time_in_minutes"`
	HashingAlgorithm             string                 `json:"hashing_algorithm"`
	Categories                   []string               `json:"categories"`
	Description                  Description            `json:"description"`
	Links                        map[string]interface{} `json:"links"`
	Image                        Image                  `json:"image"`
	CountryOrigin                string                 `json:"country_origin"`
	GenesisDate                  string                 `json:"genesis_date"`
	SentimentVotesUpPercentage   float64                `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64                `json:"sentiment_votes_down_percentage"`
	MarketCapRank                int                    `json:"market_cap_rank"`
	CoingeckoRank                int                    `json:"coingecko_rank"`
	CoingeckoScore               float64                `json:"coingecko_score"`
	//PublicNotice                 interface{}            `json:"public_notice"`
	//AdditionalNotices            []interface{}          `json:"additional_notices"`
	//DeveloperScore               float64                `json:"developer_score"`
	//CommunityScore               float64                `json:"community_score"`
	//LiquidityScore               float64                `json:"liquidity_score"`
	//PublicInterestScore          float64                `json:"public_interest_score"`
	//PublicInterestStats          struct {
	//	AlexaRank   int `json:"alexa_rank"`
	//	BingMatches int `json:"bing_matches"`
	//} `json:"public_interest_stats"`
	StatusUpdates []StatusUpdate `json:"status_updates"`
	LastUpdated   time.Time      `json:"last_updated"`
}

func (r *Description) UnmarshalJSON(bs []byte) error {
	data := map[string]string{}
	err := json.Unmarshal(bs, &data)
	if err != nil {
		return err
	}
	*r = Description(data["en"])
	return nil
}

func (r *Platforms) UnmarshalJSON(bs []byte) error {
	data := map[string]string{}
	err := json.Unmarshal(bs, &data)
	if err != nil {
		return err
	}
	*r = map[string]string{}
	for k, v := range data {
		if len(k) > 0 && len(v) > 0 {
			(*r)[k] = v
		}
	}
	return nil
}

type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}
type SparklineIn7D struct {
	Price []float64 `json:"price"`
}
type Market struct {
	Coin
	Supply
	Image                               string         `json:"image"`
	CurrentPrice                        float64        `json:"current_price"`
	MarketCap                           float64        `json:"market_cap"`
	MarketCapRank                       int            `json:"market_cap_rank"`
	FullyDilutedValuation               float64        `json:"fully_diluted_valuation"`
	TotalVolume                         float64        `json:"total_volume"`
	High24H                             float64        `json:"high_24h"`
	Low24H                              float64        `json:"low_24h"`
	PriceChange24H                      float64        `json:"price_change_24h"`
	PriceChangePercentage24H            float64        `json:"price_change_percentage_24h"`
	MarketCapChange24H                  float64        `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H        float64        `json:"market_cap_change_percentage_24h"`
	Ath                                 float64        `json:"ath"`
	AthChangePercentage                 float64        `json:"ath_change_percentage"`
	AthDate                             time.Time      `json:"ath_date"`
	Atl                                 float64        `json:"atl"`
	AtlChangePercentage                 float64        `json:"atl_change_percentage"`
	AtlDate                             time.Time      `json:"atl_date"`
	Roi                                 *ROI           `json:"roi"`
	LastUpdated                         time.Time      `json:"last_updated"`
	SparklineIn7D                       *SparklineIn7D `json:"sparkline_in_7d"`
	PriceChangePercentage1HInCurrency   *float64       `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24HInCurrency  *float64       `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage14DInCurrency  *float64       `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage7DInCurrency   *float64       `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage30DInCurrency  *float64       `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage200DInCurrency *float64       `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1YInCurrency   *float64       `json:"price_change_percentage_1y_in_currency"`
}

type Charts struct {
	Prices       [][2]float64 `json:"prices"`
	MarketCaps   [][2]float64 `json:"market_caps"`
	TotalVolumes [][2]float64 `json:"total_volumes"`
}

type OHLC [][5]float64

type ExchangeList []struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Exchange struct {
	Id                          string   `json:"id"`
	Name                        string   `json:"name"`
	YearEstablished             *int     `json:"year_established"`
	Country                     *string  `json:"country"`
	Description                 *string  `json:"description"`
	Url                         string   `json:"url"`
	Image                       *string  `json:"image"`
	HasTradingIncentive         *bool    `json:"has_trading_incentive"`
	TrustScore                  *int     `json:"trust_score"`
	TrustScoreRank              *int     `json:"trust_score_rank"`
	TradeVolume24HBtc           *float64 `json:"trade_volume_24h_btc"`
	TradeVolume24HBtcNormalized *float64 `json:"trade_volume_24h_btc_normalized"`
}
