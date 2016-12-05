package oauthWrapper

import "net/http"

const (
	base   = "https://api.tradeking.com/v1"
	stream = "https://stream.tradeking.com/v1"
)

// This example shows how to sign a request when the URL Opaque field is used.
// See the note at http://golang.org/pkg/net/url/#URL for information on the
// use of the URL Opaque field.
func Request(uri, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, base+uri, nil)
	if err != nil {
		return nil, err
	}

	// Sign the request.
	if err = client.SetAuthorizationHeader(req.Header, credentials, method, req.URL, nil); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func Stream(uri, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, stream+uri, nil)
	if err != nil {
		return nil, err
	}

	// Sign the request.
	if err = client.SetAuthorizationHeader(req.Header, credentials, method, req.URL, nil); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
