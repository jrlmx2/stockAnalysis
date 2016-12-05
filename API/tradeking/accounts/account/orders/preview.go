package orders

const uri = "/preview"

func Preview(id int64) {
	req, err := oauthWrapper.Request(URIWithID(id)+uri, "GET")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}
