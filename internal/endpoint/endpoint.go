package endpoint

import "golang.org/x/oauth2"

var Capsule = oauth2.Endpoint{
	AuthURL:  "https://api.capsulecrm.com/oauth/authorise",
	TokenURL: "https://api.capsulecrm.com/oauth/token",
}
