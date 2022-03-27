# coingecko-go

Golang Client for [CoinGecko](https://coingecko.com/)

- Dependency-Free
- Required parameters handled before the requests
- Messy payloads simplified, not well organized fields omitted(see also `/coins/{id}`), open `issue` for your needs
- `Nullable` fields have pointer types

## Progress Tracker

### Ping

- [X] ping Simple

### Simple

- [X] simple/price
- [ ] simple/token_price/{id}
- [X] simple/supported_vs_currencies

### Coins

- [X] coins/list
- [X] coins/markets
- [X] coins/{id}
- [ ] coins/{id}/tickers
- [ ] coins/{id}/history
- [X] coins/{id}/market_chart
- [ ] coins/{id}/market_chart/range
- [ ] coins/{id}/status_updates
- [X] coins/{id}/ohlc

### Contract

- [ ] coins/{id}/contract/{contract_address}
- [ ] coins/{id}/contract/{contract_address}/market_chart/
- [ ] coins/{id}/contract/{contract_address}/market_chart/range

### Asset Platforms

- [ ] asset_platforms

### Categories

- [ ] coins/categories/list
- [ ] coins/categories

### Exchanges

- [X] exchanges
- [X] exchanges/list
- [ ] exchanges/{id}
- [ ] exchanges/{id}/tickers
- [ ] exchanges/{id}/status_updates
- [ ] exchanges/{id}/volume_chart

### Finance

- [ ] finance_platforms
- [ ] finance_products

### Indexes

- [ ] indexes
- [ ] indexes/{market_id}/{id}
- [ ] indexes/list

### Derivatives

- [ ] derivatives
- [ ] derivatives/exchanges
- [ ] derivatives/exchanges/{id}
- [ ] derivatives/exchanges/list

### Status Updates

- [ ] status_updates

### Events

- [ ] events
- [ ] events/countries
- [ ] events/types

### Exchange Trades

- [ ] exchange_rates

### Trending

- [ ] search/trending

### Global

- [ ] global
- [ ] global/decentralized_finance_defi

### Companies

- [ ] companies/public_treasury/{coin_id}
