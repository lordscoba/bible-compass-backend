package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Email    string
	Username string
	Uid      string
	Type     string
	jwt.StandardClaims
}

// var SECRET_KEY string = config.GetConfig().Server.Secret

func GenerateAllTokens(secretkey string, email string, username string, userType string, uid string) (signedToken string, err error) {
	claims := &SignedDetails{
		Email:    email,
		Username: username,
		Uid:      uid,
		Type:     userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	// refreshClaims := &SignedDetails{
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
	// 	},
	// }

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretkey))
	// refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		fmt.Println(err)
		return
	}

	return token, err
}

func ValidateToken(secretkey string, signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		// msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	return claims, msg
}
