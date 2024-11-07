package helpers

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func CreateAuthToken(claims map[string]interface{}) AuthToken {
	return AuthToken{
		AccessToken:  CreateJwtToken(claims, 15),
		RefreshToken: CreateJwtToken(claims, 60),
	}
}

func CreateJwtToken(claims map[string]interface{}, ttl int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims(claims)

	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func Decode(tokenString string) (map[string]interface{}, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		user := map[string]interface{}{
			"username": claims["username"],
			"name":     claims["name"],
		}

		return user, nil
	} else {
		return nil, fmt.Errorf("Invalid token")
	}
}

func Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		user := map[string]interface{}{
			"username": claims["username"],
			"name":     claims["name"],
		}

		return user, nil
	} else {
		return nil, fmt.Errorf("Invalid token")
	}
}
