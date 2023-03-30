package services

import (
    "context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/facebook"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
)

type FacebookOAuth struct {
    cfg *oauth2.Config
}

func NewFacebookOauth() *FacebookOAuth {
    cfg := &oauth2.Config{
        ClientID:     os.Getenv("FACEBOOK_APP_ID"),
        ClientSecret: os.Getenv("FACEBOOK_APP_SECRET"),
        Endpoint:     facebook.Endpoint,
        RedirectURL:  "http://localhost:8000/callback/facebook",
        Scopes:       []string{"public_profile", "email"},
    }

    return &FacebookOAuth{cfg: cfg}
}

func (f *FacebookOAuth) HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
    HandleLogin(w, r, f.cfg, "state")
}

func (f *FacebookOAuth) CallbackFacebook(w http.ResponseWriter, r *http.Request) {
    code := r.FormValue("code")
    if code == "" {
        http.Error(w, "code not found", http.StatusBadRequest)
    }

    log.Println("code")

    token, err := f.cfg.Exchange(context.TODO(), code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    log.Printf("token: %+v", token)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
        url.QueryEscape(token.AccessToken) + "&fields=email")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    defer resp.Body.Close()

    response, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    w.Write([]byte("Hello, I'm protected\n"))
    w.Write(response)
}
