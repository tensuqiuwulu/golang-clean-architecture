package model

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	Audience  string  `json:"aud,omitempty"`
	ExpiresAt float64 `json:"exp,omitempty"`
	Id        string  `json:"jti,omitempty"`
	IssuedAt  float64 `json:"iat,omitempty"`
	Issuer    string  `json:"iss,omitempty"`
	NotBefore float64 `json:"nbf,omitempty"`
	Subject   string  `json:"sub,omitempty"`
	jwt.StandardClaims
}
