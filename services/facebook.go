package services

import (
    "context"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/facebook"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
)

type (
    FacebookOAuth struct {
        Cfg    *oauth2.Config
        Common CommonInterface
    }

    FacebookOAuthResponse struct {
        Email string `json:"email"`
        ID    string `json:"id"`
    }
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . FacebookOAuthInterface
type FacebookOAuthInterface interface {
    HandleFacebookLogin(c *gin.Context)
    CallbackFacebook(code string) (email string, err error)
}

func NewFacebookOauth(common CommonInterface) FacebookOAuthInterface {
    cfg := &oauth2.Config{
        ClientID:     os.Getenv("FACEBOOK_APP_ID"),
        ClientSecret: os.Getenv("FACEBOOK_APP_SECRET"),
        Endpoint:     facebook.Endpoint,
        RedirectURL:  "http://localhost:8000/callback/facebook",
        Scopes:       []string{"public_profile", "email"},
    }

    return &FacebookOAuth{Cfg: cfg, Common: common}
}

func (f *FacebookOAuth) HandleFacebookLogin(c *gin.Context) {
    f.Common.HandleLogin(c, f.Cfg, "state")
}

func (f *FacebookOAuth) CallbackFacebook(code string) (email string, err error) {
    token, err := f.Cfg.Exchange(context.TODO(), code)
    if err != nil {
        return email, err
    }

    resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
        url.QueryEscape(token.AccessToken) + "&fields=email")
    if err != nil {
        return email, err
    }
    defer resp.Body.Close()

    res, err := io.ReadAll(resp.Body)
    if err != nil {
        return email, err
    }

    fbRes := FacebookOAuthResponse{}
    if err := json.Unmarshal(res, &fbRes); err != nil {
        log.Println(err)
        return email, err
    }

    return fbRes.Email, nil
}
