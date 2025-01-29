package oauth

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"optiflow/config"
	"time"
)

var (
	OAuthConfig *oauth2.Config
	secretKey   = []byte("your_secret_key")
)

func InitOAuth(cfg *config.Config) {
	OAuthConfig = &oauth2.Config{
		ClientID:     cfg.OAuth.ClientID,
		ClientSecret: cfg.OAuth.ClientSecret,
		RedirectURL:  cfg.OAuth.RedirectURL,
		Scopes:       []string{"read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
}

// GetOAuthURL generates the OAuth 2.0 authorization URL
func GetOAuthURL() string {
	return OAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

// ExchangeCodeForToken exchanges the authorization code for an access token
func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return OAuthConfig.Exchange(context.Background(), code)
}

// GenerateToken generates a new JWT token
func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			return false, errors.New("token has expired")
		}
		return true, nil
	}

	return false, errors.New("invalid token")
}
