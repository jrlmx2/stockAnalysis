package sec

var resources = map[string]string{
	"comp":        "Companies",
	"finann":      "CoreFinancials/ANN",
	"finqtr":      "CoreFinacials/QTR",
	"finttm":      "CoreFinancials/TTM",
	"finytd":      "CoreFinancials/YTD",
	"insfil":      "Insider/Filers",
	"insiss":      "Insider/Issues",
	"inssum":      "Insider/Summary",
	"instrans":    "Insider/Transactions",
	"owncurrins":  "Ownerships/CurrentIssueHolders",
	"owncurrhold": "Ownerships/CurrentOwnerHoldings",
	"ownown":      "Ownerships/Owners",
	"owniss":      "Ownerships/Issues",
}

var description = map[string]string{
	"owniss":     "description/Ownership-Issues",
	"ownown":     "description/Ownership-Owners",
	"ownownhold": "description/Ownership-CurrentownerHoldings",
	"owninshold": "description/Ownership-CurrentIssueHolders",
	"instrans":   "description/Insider-Transactions",
	"inssum":     "description/Insider-Summary",
	"insiss":     "description/Insider-Issues",
	"insfil":     "description/Insider-Filers",
	"finytd":     "description/CoreFinancials-YTD",
	"finttm":     "description/CoreFinancials-TTM",
	"finqtr":     "description/CoreFinancials-QTR",
	"finann":     "description/CoreFinancials-ANN",
	"comp":       "description/Companies",
}
