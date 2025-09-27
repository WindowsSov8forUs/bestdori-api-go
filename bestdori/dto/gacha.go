package dto

type GachaAll1Info struct {
	GachaName []*string `json:"gachaName"`
}

type GachaAll1 map[string]GachaAll1Info

type GachaAll3Info struct {
	GachaAll1Info         `json:",inline"`
	ResourceName          string    `json:"resourceName"`
	BannerAssetBundleName string    `json:"bannerAssetBundleName"`
	PublishedAt           []*string `json:"publishedAt"`
	Type                  string    `json:"type"`
	NewCards              []int     `json:"newCards"`
}

type GachaAll3 map[string]GachaAll3Info

type GachaAll5Info struct {
	GachaAll3Info `json:",inline"`
	ClosedAt      []*string `json:"closedAt"`
}

type GachaAll5 map[string]GachaAll5Info

type GachaDetail struct {
	RarityIndex int  `json:"rarityIndex"`
	Weight      int  `json:"weight"`
	Pickup      bool `json:"pickup"`
}

type GachaRate struct {
	Rate        float64 `json:"rate"`
	WeightTotal int     `json:"weightTotal"`
}

type GachaPaymentMethod struct {
	GachaId          int    `json:"gachaId"`
	PaymentMethod    string `json:"paymentMethod"`
	Quantity         int    `json:"quantity"`
	PaymentMethodId  int    `json:"paymentMethodId"`
	Count            int    `json:"count"`
	Behavior         string `json:"behavior"`
	Pickup           bool   `json:"pickup"`
	CostItemQuantity int    `json:"costItemQuantity"`
	TicketId         *int   `json:"ticketId,omitempty"`
}

type GachaInformation struct {
	Description   []*string `json:"description"`
	Term          []*string `json:"term"`
	NewMemberInfo []*string `json:"newMemberInfo"`
	Notice        []*string `json:"notice"`
}

type GachaInfo struct {
	GachaAll5Info  `json:",inline"`
	Details        []*map[string]GachaDetail `json:"details"`
	Rates          []*map[string]GachaRate   `json:"rates"`
	PaymentMethods []GachaPaymentMethod      `json:"paymentMethods"`
	Description    []*string                 `json:"description"`
	Annotation     []*string                 `json:"annotation"`
	GachaPeriod    []*string                 `json:"gachaPeriod"`
	Information    GachaInformation          `json:"information"`
}
