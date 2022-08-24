package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Chave secreta para gerar a criptografia do JWT
var jwtKey = []byte("superkey")

// Estrutura do PAYLOAD do token
type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(id string, email string) (tokenString string, err error) {

	//Gerando o payload (conteúdo do token)
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Gerando token com payload e método de assinatura
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Gerando token com payload, método de assinatura e chave secreta.
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {

	// Transformando, validando, e retornando um token.
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	// Se algum erro acontecer será cortada função e repassado o erro para a pilha
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
	}

	// Verificando se o token ainda está valido
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
	}

	return
}
