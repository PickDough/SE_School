package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type BitcoinRate struct {
	Time     string   `json:"time"`
	Currency Currency `json:"currency"`
}

type Currency struct {
	Code        string  `json:"сode"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}
