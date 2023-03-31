package services_test

import (
    "github.com/stretchr/testify/assert"
    "golang-oauth/services"
    "golang.org/x/oauth2"
    "testing"
)

func TestCommon_HandleLogin(t *testing.T) {

    t.Run("success", func(t *testing.T) {
        commonSvc := services.NewCommon()
        url, err := commonSvc.HandleLogin(&oauth2.Config{}, "")
        assert.NoError(t, err)
        assert.NotEmpty(t, url)
    })

    t.Run("error invalid url", func(t *testing.T) {
        commonSvc := services.NewCommon()
        invalidUrl := "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require"
        url, err := commonSvc.HandleLogin(&oauth2.Config{Endpoint: oauth2.Endpoint{AuthURL: invalidUrl}}, "")
        assert.NoError(t, err)
        assert.NotEmpty(t, url)
    })
}
