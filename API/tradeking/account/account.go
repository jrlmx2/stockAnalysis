package account

type Accounts struct {
	Accounts []*Account `xml:"accounts"`
}

type Account struct {
	Summary *AccountSummary `xml:"accountsummary"`
}

type AccountSummary struct {
	Holdings *AccountHoldings `xml:"accountholdings"`
	ID       int64            `xml:"account"`
	Balance  *AccountBalance  `xml:"accountbalance"`
}

type AccountBalance struct {
	Value       float64      `xml:"accountvalue"`
	ID          int64        `xml:"account"`
	BuyingPower *BuyingPower `xml:"buyingpower"`
	Money       *Money       `xml:"money"`
	Securities  *Securities  `xml:"securities"`
	FedCall     float32      `xml:"fedcall"`
	HouseCall   float32      `xml:"housecall"`
}

type BuyingPower struct {
	Withdrawable     float64 `xml:"cashavailableforwithdrawal"`
	Daytrading       float64 `xml:"daytrading"`
	EquityPercentage float32 `xml:"equitypercentage"`
	Options          float64 `xml:"options"`
	Soddaytrading    float64 `xml:"soddaytrading"`
	Sodoptions       float64 `xml:"sodoptions"`
	Sodstock         float64 `xml:"sodstock"`
	Stock            float64 `xml:"stock"`
}

type Money struct {
	AccruedInterest   float64 `xml:"accruedinterest"`
	Cash              float64 `xml:"cash"`
	CashAvailable     float64 `xml:"cashavailable"`
	MarginBalance     float64 `xml:"marginbalance"`
	MMF               float64 `xml:"mmf"`
	Total             float64 `xml:"total"`
	UnclearedDeposits float64 `xml:"uncleareddeposits"`
	UnsettledFunds    float64 `xml:"unsettledfunds"`
	Yield             float64 `xml:"yield"`
}

type Securities struct {
	LongOptions  float64 `xml:"longoptions"`
	LongStocks   float64 `xml:"longstocks"`
	Options      float64 `xml:"options"`
	ShortOptions float64 `xml:"shortoptions"`
	ShortStocks  float64 `xml:"shortstocks"`
	Stocks       float64 `xml:"stocks"`
	Total        float64 `xml:"total"`
}

type AccountHoldings struct {
	Holdings        []*Holding `xml:"holding"`
	DisplayData     *Data      `xml:"displaydata"`
	TotalSecurities float64    `xml:"totalsecurities"`
}

type Data struct {
	Total string `xml:"totalsecurities"`
}

type Holding struct {
	AccountType       int              `xml:"accounttype"`
	Costbasis         float64          `xml:"costbasis"`
	Data              *DisplayData     `xml:"displaydata"`
	Gainloss          float64          `xml:"gainloss"`
	Instrument        SymbolDescriptor `xml:"instrument"`
	MarketValue       float64          `xml:"marketvalue"`
	MarketValueChange float64          `xml:"marketvaluechange"`
	Price             float64          `xml:"price"`
	PurchasePrice     float64          `xml:"price"`
	Quantity          float64          `xml:"qty"`
	Quote             *Quote           `xml:"quote"`
	Underlying        *Underlying      `xml:"underlying"`
}

type SymbolDescriptor struct {
	CUSIP        int64   `xml:"cusip"`
	Description  string  `xml:"desc"`
	Factor       float32 `xml:"factor"`
	SecurityType string  `xml:"sectyp"`
	Symbol       string  `xml:"sym"`
}

type Underlying struct {
	//???
}

type Quote struct {
	Change    float64 `xml:"change"`
	Lastprice float64 `xml:"lastprice"`
}

type DisplayData struct {
	AccountType       string `xml:"accounttype"`
	AssetClass        string `xml:"assetclass"`
	Change            string `xml:"change"`
	Costbasis         string `xml:"costbasis"`
	Description       string `xml:"desc"`
	Last              string `xml:"lastprice"`
	MarketValue       string `xml:"marketvalue"`
	MarkeyValueChange string `xml:"marketvaluechange"`
	Quantity          int    `xml:"qty"`
	Symbol            string `xml:"symbol"`
}
