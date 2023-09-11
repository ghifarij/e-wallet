package security

import (
	"Kelompok-2/dompet-online/config"
	"Kelompok-2/dompet-online/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJwtToken(user model.Users) (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return "", err
	}

	now := time.Now()
	end := now.Add(time.Duration(60) * time.Minute)

	claims := &AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.UserName,
	}

	tokenJwt := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	tokenString, err := tokenJwt.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create jwt token: %v", err.Error())
	}
	return tokenString, nil
}

func VerifyJwtToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signin method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
