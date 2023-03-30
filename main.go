package main

import (
    "github.com/joho/godotenv"
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

    fb := services.NewFacebookOauth()

    http.HandleFunc("/login/facebook", fb.HandleFacebookLogin)
    http.HandleFunc("/callback/facebook", fb.CallbackFacebook)

    log.Println("application started on port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
