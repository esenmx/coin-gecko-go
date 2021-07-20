package gocko

import (
	"errors"
	"strconv"
	"strings"
	"unsafe"
)

var MissingParameterError = errors.New("missing parameter")
var InvalidParameterError = errors.New("invalid parameter")

type QueryParams interface {
	toQuery() (map[string]string, error)
}

type CoinsParams struct {
	includePlatform bool
}

func (c CoinsParams) toQuery() (map[string]string, error) {
	return map[string]string{"include_platform": strconv.FormatBool(c.includePlatform)}, nil
}

type SimplePriceParams struct {
	Ids []ID // required
	//VsCurrencies         []Currency // required
	IncludeMarketCap     bool
	Include24hrVol       bool
	Include24hrChange    bool
	IncludeLastUpdatedAt bool
}

func (p SimplePriceParams) toQuery() (map[string]string, error) {
	if len(p.Ids) == 0 {
		return nil, MissingParameterError
	}
	return map[string]string{
		"ids": strings.Join(*(*[]string)(unsafe.Pointer(&p.Ids)), ","),
		//"vs_currencies":           strings.Join(*(*[]string)(unsafe.Pointer(&p.VsCurrencies)), ","),
		"vs_currencies":           "usd",
		"include_market_cap":      strconv.FormatBool(p.IncludeMarketCap),
		"include_24hr_vol":        strconv.FormatBool(p.Include24hrVol),
		"include_24hr_change":     strconv.FormatBool(p.Include24hrChange),
		"include_last_updated_at": strconv.FormatBool(p.IncludeLastUpdatedAt),
	}, nil

}

type CoinsMarketsParams struct {
	VsCurrency            Currency // required usd, eur, jpy, etc
	Ids                   []ID
	Category              string // decentralized_finance_defi, stablecoins
	Order                 string // gecko_desc, gecko_asc, market_cap_asc, market_cap_desc, volume_asc, volume_desc, id_asc, id_desc
	PerPage               int    // max 250
	Page                  int
	PriceChangePercentage string // 1h, 24h, 7d, 14d, 30d, 200d, 1y (eg. '1h,24h,7d' comma-separated)
	Sparkline             bool
}

func (c CoinsMarketsParams) toQuery() (map[string]string, error) {
	if len(c.VsCurrency) == 0 {
		return nil, MissingParameterError
	}
	if c.Page < 0 || c.PerPage < 0 || c.PerPage > 250 {
		return nil, InvalidParameterError
	}
	q := map[string]string{}
	q["vs_currency"] = string(c.VsCurrency)
	if len(c.Category) > 0 {
		q["category"] = c.Category
	}
	if c.Ids != nil {
		if len(c.Ids) > 0 {
			q["ids"] = strings.Join(*(*[]string)(unsafe.Pointer(&c.Ids)), ",")
		}
	}
	if len(c.Order) > 0 {
		q["order"] = c.Order
	}
	if c.PerPage > 0 {
		q["per_page"] = strconv.Itoa(c.PerPage)
	}
	if c.Page > 0 {
		q["page"] = strconv.Itoa(c.Page)
	}
	if len(c.PriceChangePercentage) > 0 {
		q["price_change_percentage"] = c.PriceChangePercentage
	}
	q["sparkline"] = strconv.FormatBool(c.Sparkline)
	return q, nil
}

type CoinsDataParams struct {
	Id ID
	//Localization  bool
	//Tickers       bool
	//MarketData    bool
	//CommunityData bool
	//DeveloperData bool
	//Sparkline     bool
}

func (c CoinsDataParams) toQuery() (map[string]string, error) {
	if len(c.Id) == 0 {
		return nil, MissingParameterError
	}
	return map[string]string{
		//"localization":   strconv.FormatBool(c.Localization),
		//"tickers":        strconv.FormatBool(c.Tickers),
		//"market_data":    strconv.FormatBool(c.MarketData),
		//"community_data": strconv.FormatBool(c.CommunityData),
		//"developer_data": strconv.FormatBool(c.DeveloperData),
		//"sparkline":      strconv.FormatBool(c.Sparkline),
	}, nil
}

type CoinsChartsParams struct {
	Id         ID       // required
	VsCurrency Currency // required
	Days       string   // required (eg. 1,14,30,max) 5min interval 1 day, 1h interval 1-90days, 1d interval 90+days
}

func (c CoinsChartsParams) toQuery() (map[string]string, error) {
	if len(c.Id) == 0 || len(c.VsCurrency) == 0 || len(c.Days) == 0 {
		return nil, MissingParameterError
	}
	return map[string]string{
		"vs_currency": string(c.VsCurrency),
		"days":        c.Days,
	}, nil
}

type CoinsOHLCParams struct {
	Id         ID       // required
	VsCurrency Currency // required
	Days       string   // required 1/7/14/30/90/180/365/max, intervals: 1-2d:30m, 3-30d:4h, 31+d:4d
}

func (c CoinsOHLCParams) toQuery() (map[string]string, error) {
	if len(c.Id) == 0 || len(c.VsCurrency) == 0 || len(c.Days) == 0 {
		return nil, MissingParameterError
	}
	return map[string]string{
		"vs_currency": string(c.VsCurrency),
		"days":        c.Days,
	}, nil
}

type ExchangesParams struct {
	PerPage int // max 250
	Page    int
}

func (e ExchangesParams) toQuery() (map[string]string, error) {
	if e.PerPage > 250 || e.PerPage < 0 || e.Page < 0 {
		return nil, InvalidParameterError
	}
	return map[string]string{"per_page": strconv.Itoa(e.PerPage), "page": strconv.Itoa(e.Page)}, nil
}
