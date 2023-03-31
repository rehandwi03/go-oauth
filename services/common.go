package services

import (
    "golang.org/x/oauth2"
    "log"
    "net/url"
    "strings"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . CommonInterface
type CommonInterface interface {
    HandleLogin(oauthConf *oauth2.Config, oauthState string) (res string, err error)
}

type common struct {
}

func NewCommon() CommonInterface {
    return &common{}
}

func (c *common) HandleLogin(oauthConf *oauth2.Config, oauthState string) (res string, err error) {
    URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
    if err != nil {
        log.Println(err)
        return res, err
    }
    parameters := url.Values{}
    parameters.Add("client_id", oauthConf.ClientID)
    parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
    parameters.Add("redirect_uri", oauthConf.RedirectURL)
    parameters.Add("response_type", "code")
    parameters.Add("state", oauthState)
    URL.RawQuery = parameters.Encode()

    return URL.String(), nil
}
