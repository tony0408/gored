package exchange

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

type ExchangeName string
type DataSource string
type UpdateMethod string

const (
	API_TIGGER  UpdateMethod = "API_TIGGER"
	TIME_TIGGER UpdateMethod = "TIME_TIGGER"

	EXCHANGE_API     DataSource = "EXCHANGE_API"
	MICROSERVICE_API DataSource = "MICROSERVICE_API"
	JSON_FILE        DataSource = "JSON_FILE"
	PSQL             DataSource = "PSQL"

	BCEX         ExchangeName = "BCEX"
	BGOGO        ExchangeName = "BGOGO"
	BIBOX        ExchangeName = "BIBOX"
	BIGONE       ExchangeName = "BIGONE"
	BIKI         ExchangeName = "BIKI"
	BINANCE      ExchangeName = "BINANCE"
	BITFINEX     ExchangeName = "BITFINEX"
	BITFOREX     ExchangeName = "BITFOREX"
	BITMART      ExchangeName = "BITMART"
	BITMAX       ExchangeName = "BITMAX"
	BITMEX       ExchangeName = "BITMEX"
	BITSTAMP     ExchangeName = "BITSTAMP"
	BITTREX      ExchangeName = "BITTREX"
	BITLISH      ExchangeName = "BITLISH"
	BITRUE       ExchangeName = "BITRUE"
	BITZ         ExchangeName = "BITZ"
	BLANK        ExchangeName = "BLANK"
	BLEUTRADE    ExchangeName = "BLEUTRADE"
	COINMEX      ExchangeName = "COINMEX"
	COINBASE     ExchangeName = "COINBASE"
	COINBENE     ExchangeName = "COINBENE"
	COINEAL      ExchangeName = "COINEAL"
	COINEX       ExchangeName = "COINEX"
	COINSUPER    ExchangeName = "COINSUPER"
	CRYPTOPIA    ExchangeName = "CRYPTOPIA"
	DIGIFINEX    ExchangeName = "DIGIFINEX"
	DRAGONEX     ExchangeName = "DRAGONEX"
	EXMO         ExchangeName = "EXMO"
	EXX          ExchangeName = "EXX"
	FATBTC       ExchangeName = "FATBTC"
	FCOIN        ExchangeName = "FCOIN"
	GATEIO       ExchangeName = "GATEIO"
	GRAVIEX      ExchangeName = "GRAVIEX"
	HITBTC       ExchangeName = "HITBTC"
	HOTBIT       ExchangeName = "HOTBIT"
	HUOBI        ExchangeName = "HUOBI"
	HUOBIOTC     ExchangeName = "HUOBIOTC"
	IBANKDIGITAL ExchangeName = "IBANKDIGITAL"
	IDAX         ExchangeName = "IDAX"
	IDEX         ExchangeName = "IDEX"
	KRAKEN       ExchangeName = "KRAKEN"
	KUCOIN       ExchangeName = "KUCOIN"
	LBANK        ExchangeName = "LBANK"
	LIQUID       ExchangeName = "LIQUID"
	LIVECOIN     ExchangeName = "LIVECOIN"
	MXC          ExchangeName = "MXC"
	OKEX         ExchangeName = "OKEX"
	OTCBTC       ExchangeName = "OTCBTC"
	P2PB2B       ExchangeName = "P2PB2B"
	POLONIEX     ExchangeName = "POLONIEX"
	RIGHTBTC     ExchangeName = "RIGHTBTC"
	STEX         ExchangeName = "STEX"
	TOKOK        ExchangeName = "TOKOK"
	TIDEX        ExchangeName = "TIDEX"
	TOPBTC       ExchangeName = "TOPBTC"
	TRADEOGRE    ExchangeName = "TRADEOGRE"
	TRADESATOSHI ExchangeName = "TRADESATOSHI"
	UEX          ExchangeName = "UEX"
	ZBEX         ExchangeName = "ZBEX"
)

func (e *ExchangeManager) initExchangeNames() {
	supportList = append(supportList, BINANCE)  // ID = 1
	supportList = append(supportList, BITTREX)  // ID = 2
	supportList = append(supportList, COINEX)   // ID = 3
	supportList = append(supportList, STEX)     // ID = 4
	supportList = append(supportList, BITMEX)   // ID = 5
	supportList = append(supportList, KUCOIN)   // ID = 6
	supportList = append(supportList, BITMAX)   // ID = 7
	supportList = append(supportList, HUOBIOTC) // ID = 8
	supportList = append(supportList, BITSTAMP) // ID = 9
	supportList = append(supportList, OTCBTC)   // ID = 10
	supportList = append(supportList, HUOBI)    // ID = 11
	supportList = append(supportList, BIBOX)    // ID = 12
	supportList = append(supportList, OKEX)     // ID = 13
	supportList = append(supportList, BITZ)     // ID = 14
	supportList = append(supportList, HITBTC)   // ID = 15
	supportList = append(supportList, DRAGONEX) // ID = 16
	supportList = append(supportList, BIGONE)   // ID = 17
	supportList = append(supportList, BITFINEX) // ID = 18
	supportList = append(supportList, GATEIO)   // ID = 19
	supportList = append(supportList, IDEX)     // ID = 20
	supportList = append(supportList, LIQUID)   // ID = 21
	supportList = append(supportList, BITFOREX) // ID = 22
	supportList = append(supportList, TOKOK)    // ID = 23
	supportList = append(supportList, MXC)      // ID = 24
}
