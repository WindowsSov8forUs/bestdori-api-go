package dto

type MiracleTicketExchangeInfo struct {
	Name            []*string `json:"name"`
	Ids             []*[]int  `json:"ids"`
	ExchangeStartAt []*string `json:"exchangeStartAt"`
	ExchangeEndAt   []*string `json:"exchangeEndAt"`
}

type MiracleTicketExchangesAll5 map[string]MiracleTicketExchangeInfo
