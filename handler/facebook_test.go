package handler_test

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "golang-oauth/handler"
    "golang-oauth/services/servicesfakes"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler_CallbackFacebookHandler(t *testing.T) {
    fbSvcMock := &servicesfakes.FakeFacebookOAuthInterface{}
    fbSvcMock.CallbackFacebookReturnsOnCall(0, "", errors.New("code not found"))
    h := handler.NewFacebookOauthHandler(fbSvcMock)
    r := gin.New()
    r.GET("/callback/facebook", h.CallbackFacebookHandler)

    tt := []struct {
        Name       string
        URL        string
        Method     string
        StatusCode int
    }{
        {Name: "Success", URL: "/callback/facebook?code=123123", Method: http.MethodGet, StatusCode: http.StatusOK},
        {Name: "Error Code Not Found", URL: "/callback/facebook", Method: http.MethodGet, StatusCode: http.StatusBadRequest},
        {Name: "Error Code Not Valid", URL: "/callback/facebook?code=123123", Method: http.MethodGet, StatusCode: http.StatusBadRequest},
    }

    for _, t := range tt {
        req, _ := http.NewRequest(t.Method, t.URL, nil)
        w := httptest.NewRecorder()
        r.ServeHTTP(w, req)
    }
}

func TestHandler_HandleFacebookLoginHandler(t *testing.T) {
    fbSvcMock := &servicesfakes.FakeFacebookOAuthInterface{}
    h := handler.NewFacebookOauthHandler(fbSvcMock)

    r := gin.New()
    r.GET("/login/facebook", h.HandleFacebookLoginHandler)

    req, _ := http.NewRequest(http.MethodGet, "/login/facebook", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    // assert.Equal(t, "/callback/facebook", w.Header().Get("Location"))
}
