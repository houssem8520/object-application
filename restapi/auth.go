package restapi

// THIS CODE HAS NOT BEEN GENERATED

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/object.application/internal/conf"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	// state carries an internal token during the oauth2 workflow
	// we just need a non empty initial value
	state = "foobar" // Don't make this a global in production.

	// our endpoint to be called back by the redirected client
	callbackURL = "http://127.0.0.1:8080/auth/callback"

	// the description of the OAuth2 flow

)

func GetAuthConfig() oauth2.Config {
	endpoint := oauth2.Endpoint{
		AuthURL:  conf.GetAuthURL(),
		TokenURL: conf.GetTokenURL(),
	}

	config := oauth2.Config{
		ClientID:     conf.GetClientId(),
		ClientSecret: conf.GetClientSecret(),
		Endpoint:     endpoint,
		RedirectURL:  callbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return config
}

func login(r *http.Request) middleware.Responder {
	config := GetAuthConfig()
	// implements the login with a redirection
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
		})
}

func callback(r *http.Request) (string, error) {
	// we expect the redirected client to call us back
	// with 2 query params: state and code.
	// We use directly the Request params here, since we did not
	// bother to document these parameters in the spec.
	config := GetAuthConfig()

	if r.URL.Query().Get("state") != state {
		log.Println("state did not match")
		return "", fmt.Errorf("state did not match")
	}

	myClient := &http.Client{}

	parentContext := context.Background()
	ctx := oidc.ClientContext(parentContext, myClient)

	authCode := r.URL.Query().Get("code")
	log.Printf("Authorization code: %v\n", authCode)

	// Exchange converts an authorization code into a token.
	// Under the hood, the oauth2 client POST a request to do so
	// at tokenURL, then redirects...
	oauth2Token, err := config.Exchange(ctx, authCode)
	if err != nil {
		log.Println("failed to exchange token", err.Error())
		return "", fmt.Errorf("failed to exchange token")
	}

	// the authorization server's returned token
	log.Println("Raw token data:", oauth2Token)
	return oauth2Token.AccessToken, nil
}

func authenticated(token string) (bool, error) {
	// validates the token by sending a request at userInfoURL
	bearToken := "Bearer " + token
	req, err := http.NewRequest("GET", conf.GetUserInfoURL(), nil)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}

	req.Header.Add("Authorization", bearToken)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("fail to get response: %v", err)
	}
	if resp.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}
