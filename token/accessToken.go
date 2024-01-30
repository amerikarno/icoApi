package token

import (
	"fmt"

	"github.com/amerikarno/icoApi/common"
	"github.com/amerikarno/icoApi/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AccessClaims struct {
	models.JwtUserModel
	jwt.RegisteredClaims
}

func NewAccessClaims(u *models.JwtUserModel) (ac *AccessClaims) {
	ac = &AccessClaims{JwtUserModel: *u}
	n := now()
	ac.IssuedAt = jwt.NewNumericDate(n)
	ac.ExpiresAt = jwt.NewNumericDate(n.Add(accessExp))
	return
}

// Verify access token.
// Return error if not found or invalid.
func ExtractAccessClaims(c echo.Context) (*AccessClaims, error) {
	tokenString := common.ExtractJWT(c)
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.AccessSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid access token: %v", err)
	}
	ac, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid access token: %v", "invlid claims")
	}

	if ac.IsExpired() {
		return nil, fmt.Errorf("invalid access token: %v", "expired")
	}

	return ac, nil
}

func (ac *AccessClaims) IsExpired() bool {
	now := now()
	exp := ac.ExpiresAt.Time
	return now.After(exp)
}

func (ac *AccessClaims) JwtString() (accessTokenString string, err error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessTokenString, err = accessToken.SignedString([]byte(common.AccessSecret))
	if err != nil {
		return "", fmt.Errorf("error access token SignedString")
	}
	return
}
