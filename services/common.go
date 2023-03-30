package services

import (
    "golang.org/x/oauth2"
    "log"
    "net/http"
    "net/url"
    "strings"
)

func HandleLogin(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config, oauthState string) {
    URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
    if err != nil {
        log.Println(err)
    }
    parameters := url.Values{}
    parameters.Add("client_id", oauthConf.ClientID)
    parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
    parameters.Add("redirect_uri", oauthConf.RedirectURL)
    parameters.Add("response_type", "code")
    parameters.Add("state", oauthState)
    URL.RawQuery = parameters.Encode()
    url := URL.String()

    log.Println(url)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
