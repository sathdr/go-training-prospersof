package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateToken() (string, error) {
	mySigningKey := []byte("drowssap")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "Sathaporn",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// TokenHandler handle token generation request
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	t, err := generateToken()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": t,
	})
}

func main() {

	mySigningKey := []byte("drowssap")

	if tokenString, err := generateToken(); err != nil {
		log.Fatal(err)
	} else {
		// validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		fmt.Println(token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}
	}

}
