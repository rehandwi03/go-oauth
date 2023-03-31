package transport

type FacebookResponse struct {
    Response
    Email string `json:"email"`
}
