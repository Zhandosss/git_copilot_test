package capsule

import "golang.org/x/oauth2"

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.capsulecrm.com/oauth/authorise",
	TokenURL: "https://api.capsulecrm.com/oauth/token",
}
