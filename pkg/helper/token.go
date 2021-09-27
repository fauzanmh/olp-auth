package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fauzanmh/olp-auth/schema"
	"github.com/fauzanmh/olp-auth/schema/object"
	"github.com/spf13/viper"
)

// ApplicationName is for JWT Application Name
const ApplicationName = "popaket-dashboard"

// JWTSigningMethod is JWT's signing method
var jwtSigningMethod = jwt.SigningMethodHS256

// GenerateToken will generate both access and refresh token
// for current user.
// Access Token will be expired in 15 Minutes
// Refresh Token will be expired in 6 Months
func GenerateToken(username string) (token *schema.Token, err error) {
	jwtToken, expiresAt, err := GenerateJWT(username)
	if err != nil {
		return
	}

	token = &schema.Token{
		Token:     jwtToken,
		ExpiresAt: expiresAt,
	}

	return
}

// GenerateJWT is
func GenerateJWT(username string) (signedToken string, expiresAt int64, err error) {
	exp := time.Now().UTC().Add(viper.GetDuration("auth.access_token_expiry"))
	claims := object.CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(viper.GetString("auth.access_token_secret")))
	if err != nil {
		return signedToken, exp.Unix(), err
	}

	return signedToken, exp.Unix(), err
}
