package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "golang-oauth/handler"
    "golang-oauth/services"
    "log"
    "net/http"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("error load env: %+v", err)
    }
}

func main() {

    r := gin.New()
    fbSvc := services.NewFacebookOauth()
    fbHandler := handler.NewFacebookOauthHandler(fbSvc)

    r.GET("/login/facebook", fbHandler.HandleFacebookLoginHandler)
    r.GET("/callback/facebook", fbHandler.CallbackFacebookHandler)

    // http.HandleFunc("/login/facebook", fbSvc.HandleFacebookLoginHandler)
    // http.HandleFunc("/callback/facebook", fbSvc.CallbackFacebookHandler)

    log.Println("application started on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r.Handler()))
}
