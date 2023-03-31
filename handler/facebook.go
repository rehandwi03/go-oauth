package handler

import (
    "github.com/gin-gonic/gin"
    "golang-oauth/services"
    "golang-oauth/transport"
    "net/http"
)

type handler struct {
    svc services.FacebookOAuthInterface
}

func NewFacebookOauthHandler(svc services.FacebookOAuthInterface) *handler {
    return &handler{svc: svc}
}

func (h *handler) CallbackFacebookHandler(c *gin.Context) {
    code := c.Query("code")
    if code == "" {
        c.AbortWithStatusJSON(http.StatusBadRequest, transport.Response{Message: "code not found"})
        return
    }

    res, err := h.svc.CallbackFacebook(code)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, transport.Response{Message: err.Error()})
        return
    }

    c.JSON(http.StatusOK, transport.FacebookResponse{
        Response: transport.Response{Message: "success"},
        Email:    res,
    })
    return
}

func (h *handler) HandleFacebookLoginHandler(c *gin.Context) {
    h.svc.HandleFacebookLogin(c)
}
