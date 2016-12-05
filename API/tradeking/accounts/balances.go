package accounts

const uri = "/balances"

func PullAccountsBalances() {
	req, err := oauthWrapper.Request(URI+uri, "GET")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}
