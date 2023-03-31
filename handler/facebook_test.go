package handler_test

import (
    "errors"
    "github.com/gin-gonic/gin"
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
