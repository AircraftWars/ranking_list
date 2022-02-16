package tools

import "github.com/dgrijalva/jwt-go"

var secret = "f4e2e52034348f86b67cde581c0f9eb5"

func CreateToken(username string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["username"].(string), err
}
