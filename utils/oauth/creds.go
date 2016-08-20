package oauthWrapper

import (
	"github.com/garyburd/go-oauth/oauth"
)

var credentials *oauth.Credentials
var client *oauth.Client

func SetCredentials(token, secret string) {
	credentials = &oauth.Credentials{
		Token:  token,
		Secret: secret,
	}
}

func SetClient(token, secret string) {
	client = &oauth.Client{
		Credentials:     oauth.Credentials{Token: token, Secret: secret},
		SignatureMethod: 0,
	}
}
